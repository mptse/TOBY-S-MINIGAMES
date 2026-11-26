package tests

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	gamesDomain "github.com/toby-minigames/server/internal/modules/games/domain"
	gamesInfra "github.com/toby-minigames/server/internal/modules/games/infrastructure"
	usersDomain "github.com/toby-minigames/server/internal/modules/users/domain"
	usersInfra "github.com/toby-minigames/server/internal/modules/users/infrastructure"
	"github.com/toby-minigames/server/internal/shared/database"
)

func TestPersistenceIntegration_PlayerAndGameProfileJSONB(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping PostgreSQL integration test")
	}

	db, err := database.NewConnection(dsn, "development", database.DefaultPoolConfig())
	if err != nil {
		t.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := database.Ping(ctx, db); err != nil {
		t.Fatalf("database ping failed: %v", err)
	}

	playerRepo := usersInfra.NewPostgresPlayerRepository(db)
	gameRepo := gamesInfra.NewPostgresGameRepository(db)
	profileRepo := usersInfra.NewPostgresGameProfileRepository(db)

	suffix := uuid.NewString()[:8]
	player, err := usersDomain.NewPlayer(
		fmt.Sprintf("integration_%s@example.com", suffix),
		"example",
		fmt.Sprintf("Player_%s", suffix),
	)
	if err != nil {
		t.Fatalf("failed creating player entity: %v", err)
	}

	game, err := gamesDomain.NewGame(
		fmt.Sprintf("Integration Game %s", suffix),
		fmt.Sprintf("integration-game-%s", suffix),
		"Temporary game for integration testing",
		"test-thumb.png",
	)
	if err != nil {
		t.Fatalf("failed creating game entity: %v", err)
	}

	if err := playerRepo.Create(ctx, player); err != nil {
		t.Fatalf("failed persisting player: %v", err)
	}
	if err := gameRepo.Create(ctx, game); err != nil {
		t.Fatalf("failed persisting game: %v", err)
	}

	t.Cleanup(func() {
		cleanCtx, cleanCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cleanCancel()
		db.WithContext(cleanCtx).Where("player_id = ? AND game_id = ?", player.ID, game.ID).Delete(&usersInfra.GameProfileModel{})
		db.WithContext(cleanCtx).Where("player_id = ?", player.ID).Delete(&usersInfra.PlayerModel{})
		db.WithContext(cleanCtx).Where("game_id = ?", game.ID).Delete(&gamesInfra.GameModel{})
	})

	profile, err := usersDomain.NewGameProfile(player.ID, game.ID)
	if err != nil {
		t.Fatalf("failed creating game profile entity: %v", err)
	}
	profile.HighestScore = 1200
	profile.MatchesPlayed = 3
	profile.GameData = map[string]any{
		"unlocked_worlds": []any{"forest", "desert"},
		"best_combo":      float64(14),
	}

	if err := profileRepo.Upsert(ctx, profile); err != nil {
		t.Fatalf("failed upserting game profile: %v", err)
	}

	// Second upsert to verify UNIQUE(player_id, game_id) conflict update
	profile.HighestScore = 2500
	profile.MatchesPlayed = 4
	profile.GameData["best_combo"] = float64(22)
	profile.UpdatedAt = time.Now().UTC()

	if err := profileRepo.Upsert(ctx, profile); err != nil {
		t.Fatalf("failed updating game profile on upsert conflict: %v", err)
	}

	stored, err := profileRepo.FindByPlayerAndGame(ctx, player.ID, game.ID)
	if err != nil {
		t.Fatalf("failed finding stored game profile: %v", err)
	}

	if stored.HighestScore != 2500 || stored.MatchesPlayed != 4 {
		t.Errorf("expected highest_score=2500 and matches_played=4, got %d and %d", stored.HighestScore, stored.MatchesPlayed)
	}
	if stored.GameData["best_combo"] != float64(22) {
		t.Errorf("expected best_combo=22 in JSONB game_data, got %v", stored.GameData["best_combo"])
	}
}
