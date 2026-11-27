package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	httpserver "github.com/toby-minigames/server/internal/infrastructure/http"
	gamesDomain "github.com/toby-minigames/server/internal/modules/games/domain"
	gamesInfra "github.com/toby-minigames/server/internal/modules/games/infrastructure"
	usersDomain "github.com/toby-minigames/server/internal/modules/users/domain"
	usersInfra "github.com/toby-minigames/server/internal/modules/users/infrastructure"
	"github.com/toby-minigames/server/internal/shared/database"
	"github.com/toby-minigames/server/internal/shared/security"
)

type apiEnvelope struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
	Error   *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func doJSONRequest(t *testing.T, handler http.Handler, method, path string, body any, token string) (*httptest.ResponseRecorder, apiEnvelope) {
	t.Helper()

	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("failed encoding request body: %v", err)
		}
	}

	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	var env apiEnvelope
	if err := json.Unmarshal(rec.Body.Bytes(), &env); err != nil {
		t.Fatalf("failed decoding response JSON (%d %s): %v; body=%s", rec.Code, path, err, rec.Body.String())
	}

	return rec, env
}

func TestHTTPIntegration_AuthUsersAndGames(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping HTTP integration test")
	}

	db, err := database.NewConnection(dsn, "development", database.DefaultPoolConfig())
	if err != nil {
		t.Fatalf("failed connecting to PostgreSQL: %v", err)
	}

	jwtMgr, err := security.NewJWTManager("test", 24)
	if err != nil {
		t.Fatalf("failed creating JWT manager: %v", err)
	}

	router := httpserver.NewRouter(db, jwtMgr, "development")

	suffix := uuid.NewString()[:8]
	email := fmt.Sprintf("http_test_%s@example.com", suffix)
	gameSlug := fmt.Sprintf("http-game-%s", suffix)

	// Seed a temporary game for catalog & subprofile tests
	gameRepo := gamesInfra.NewPostgresGameRepository(db)
	profileRepo := usersInfra.NewPostgresGameProfileRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	testGame, err := gamesDomain.NewGame("HTTP Test Game", gameSlug, "Catalog test game", "thumb.png")
	if err != nil {
		t.Fatalf("failed creating game entity: %v", err)
	}
	if err := gameRepo.Create(ctx, testGame); err != nil {
		t.Fatalf("failed persisting test game: %v", err)
	}

	var createdPlayerID uuid.UUID
	t.Cleanup(func() {
		cleanCtx, cleanCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanCancel()
		if createdPlayerID != uuid.Nil {
			db.WithContext(cleanCtx).Where("player_id = ? AND game_id = ?", createdPlayerID, testGame.ID).Delete(&usersInfra.GameProfileModel{})
			db.WithContext(cleanCtx).Where("player_id = ?", createdPlayerID).Delete(&usersInfra.PlayerModel{})
		}
		db.WithContext(cleanCtx).Where("game_id = ?", testGame.ID).Delete(&gamesInfra.GameModel{})
	})

	// 1. Guest session creation & Forbidden check on /users/me
	rec, env := doJSONRequest(t, router, http.MethodPost, "/api/v1/auth/guest", nil, "")
	if rec.Code != http.StatusCreated || !env.Success {
		t.Fatalf("expected 201 for guest, got %d (%s)", rec.Code, rec.Body.String())
	}
	var guestData struct {
		Token string `json:"token"`
	}
	_ = json.Unmarshal(env.Data, &guestData)

	rec, env = doJSONRequest(t, router, http.MethodGet, "/api/v1/users/me", nil, guestData.Token)
	if rec.Code != http.StatusForbidden || env.Error == nil || env.Error.Code != "FORBIDDEN" {
		t.Fatalf("expected 403 FORBIDDEN for guest on /users/me, got %d (%s)", rec.Code, rec.Body.String())
	}

	// 2. Register new player & duplicate email conflict
	rec, env = doJSONRequest(t, router, http.MethodPost, "/api/v1/auth/register", map[string]string{
		"email":    email,
		"password": "example",
		"nickname": "HttpTester",
	}, "")
	if rec.Code != http.StatusCreated || !env.Success {
		t.Fatalf("expected 201 for register, got %d (%s)", rec.Code, rec.Body.String())
	}

	var regData struct {
		Token  string `json:"token"`
		Player struct {
			PlayerID uuid.UUID `json:"player_id"`
		} `json:"player"`
	}
	_ = json.Unmarshal(env.Data, &regData)
	createdPlayerID = regData.Player.PlayerID

	rec, env = doJSONRequest(t, router, http.MethodPost, "/api/v1/auth/register", map[string]string{
		"email":    email,
		"password": "example",
		"nickname": "Duplicate",
	}, "")
	if rec.Code != http.StatusConflict || env.Error == nil || env.Error.Code != "CONFLICT" {
		t.Fatalf("expected 409 CONFLICT for duplicate email, got %d (%s)", rec.Code, rec.Body.String())
	}

	// 3. Login with registered credentials
	rec, env = doJSONRequest(t, router, http.MethodPost, "/api/v1/auth/login", map[string]string{
		"email":    email,
		"password": "example",
	}, "")
	if rec.Code != http.StatusOK || !env.Success {
		t.Fatalf("expected 200 for login, got %d (%s)", rec.Code, rec.Body.String())
	}

	// 4. Get & Update /users/me
	rec, env = doJSONRequest(t, router, http.MethodGet, "/api/v1/users/me", nil, regData.Token)
	if rec.Code != http.StatusOK || !env.Success {
		t.Fatalf("expected 200 for GET /users/me, got %d", rec.Code)
	}

	rec, env = doJSONRequest(t, router, http.MethodPatch, "/api/v1/users/me", map[string]string{
		"nickname": "UpdatedTester",
		"bio":      "Testing bio",
	}, regData.Token)
	if rec.Code != http.StatusOK || !env.Success {
		t.Fatalf("expected 200 for PATCH /users/me, got %d (%s)", rec.Code, rec.Body.String())
	}

	// 5. Subprofile GET /users/me/games/:slug (after seeding subprofile)
	gp, _ := usersDomain.NewGameProfile(createdPlayerID, testGame.ID)
	gp.HighestScore = 999
	gp.GameData = map[string]any{"level": float64(5)}
	if err := profileRepo.Upsert(ctx, gp); err != nil {
		t.Fatalf("failed seeding game profile: %v", err)
	}

	rec, env = doJSONRequest(t, router, http.MethodGet, "/api/v1/users/me/games/"+gameSlug, nil, regData.Token)
	if rec.Code != http.StatusOK || !env.Success {
		t.Fatalf("expected 200 for GET /users/me/games/%s, got %d (%s)", gameSlug, rec.Code, rec.Body.String())
	}

	// 6. Games catalog & detail
	rec, env = doJSONRequest(t, router, http.MethodGet, "/api/v1/games", nil, "")
	if rec.Code != http.StatusOK || !env.Success {
		t.Fatalf("expected 200 for GET /games, got %d", rec.Code)
	}

	rec, env = doJSONRequest(t, router, http.MethodGet, "/api/v1/games/"+gameSlug, nil, "")
	if rec.Code != http.StatusOK || !env.Success {
		t.Fatalf("expected 200 for GET /games/%s, got %d", gameSlug, rec.Code)
	}

	rec, env = doJSONRequest(t, router, http.MethodGet, "/api/v1/games/non-existent-slug", nil, "")
	if rec.Code != http.StatusNotFound || env.Error == nil || env.Error.Code != "NOT_FOUND" {
		t.Fatalf("expected 404 NOT_FOUND for unknown slug, got %d", rec.Code)
	}
}
