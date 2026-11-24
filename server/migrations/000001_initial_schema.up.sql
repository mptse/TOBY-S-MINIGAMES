-- Migración inicial del esquema de base de datos para Toby's Minigames
-- Motor objetivo: PostgreSQL 16+

-- Extensión para generación de UUIDs v4 (gen_random_uuid disponible nativamente desde PG 13, pgcrypto como respaldo)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- 1. Tabla: players (Jugador)
CREATE TABLE players (
    player_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    nickname VARCHAR(100) NOT NULL,
    bio VARCHAR(500),
    role VARCHAR(50) NOT NULL DEFAULT 'player',
    exp INT NOT NULL DEFAULT 0 CHECK (exp >= 0),
    tobycoin INT NOT NULL DEFAULT 0 CHECK (tobycoin >= 0),
    avatar_path VARCHAR(255) NOT NULL DEFAULT 'default_avatar.png'
);

-- 2. Tabla: games (Minijuegos)
CREATE TABLE games (
    game_id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    name_slug VARCHAR(150) NOT NULL UNIQUE,
    description VARCHAR(1000) NOT NULL,
    thumbnail_slug VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    create_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. Tabla: categories (Categorías)
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name_category VARCHAR(100) NOT NULL,
    category_desc VARCHAR(500),
    category_slug VARCHAR(100) NOT NULL UNIQUE
);

-- 4. Tabla: game_categories (Tabla Intermedia N:M entre games y categories)
CREATE TABLE game_categories (
    game_id INT NOT NULL REFERENCES games(game_id) ON DELETE CASCADE,
    category_id INT NOT NULL REFERENCES categories(category_id) ON DELETE CASCADE,
    PRIMARY KEY (game_id, category_id)
);

-- 5. Tabla: match_history (Historial de Partidas)
CREATE TABLE match_history (
    match_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(player_id) ON DELETE CASCADE,
    game_id INT NOT NULL REFERENCES games(game_id) ON DELETE RESTRICT,
    game_mode VARCHAR(100) NOT NULL DEFAULT 'standard',
    score_obtained INT NOT NULL DEFAULT 0,
    xp_earned INT NOT NULL DEFAULT 0 CHECK (xp_earned >= 0),
    tobycoins_earned INT NOT NULL DEFAULT 0 CHECK (tobycoins_earned >= 0),
    played_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 6. Tabla: items (Tienda / Recompensas)
CREATE TABLE items (
    item_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    game_id INT REFERENCES games(game_id) ON DELETE SET NULL,
    item_name VARCHAR(150) NOT NULL UNIQUE,
    desc_item VARCHAR(500) NOT NULL,
    price INT NOT NULL CHECK (price >= 0),
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    categoria_slug VARCHAR(100) NOT NULL
);

-- 7. Tabla: inventories (Inventario de Jugadores)
CREATE TABLE inventories (
    inventory_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(player_id) ON DELETE CASCADE,
    item_id UUID NOT NULL REFERENCES items(item_id) ON DELETE RESTRICT,
    is_equiped BOOLEAN NOT NULL DEFAULT FALSE,
    UNIQUE (player_id, item_id)
);

-- 8. Tabla: game_profiles (Subperfiles por Minijuego - HU-14)
CREATE TABLE game_profiles (
    profile_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(player_id) ON DELETE CASCADE,
    game_id INT NOT NULL REFERENCES games(game_id) ON DELETE CASCADE,
    highest_score INT NOT NULL DEFAULT 0,
    matches_played INT NOT NULL DEFAULT 0 CHECK (matches_played >= 0),
    game_data JSONB NOT NULL DEFAULT '{}'::jsonb,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (player_id, game_id)
);

-- 9. Tabla: game_reviews (Reseñas de Juegos - HU-08)
CREATE TABLE game_reviews (
    review_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(player_id) ON DELETE CASCADE,
    game_id INT NOT NULL REFERENCES games(game_id) ON DELETE CASCADE,
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 10. Tabla: reports (Sistema de Reportes - SRS 3.2 #11)
CREATE TABLE reports (
    report_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    player_id UUID NOT NULL REFERENCES players(player_id) ON DELETE CASCADE,
    game_id INT REFERENCES games(game_id) ON DELETE SET NULL,
    description TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Índices para optimizar consultas frecuentes (rankings, historiales, inventarios, reseñas)
CREATE INDEX idx_match_history_player_id ON match_history(player_id);
CREATE INDEX idx_match_history_game_score ON match_history(game_id, score_obtained DESC);
CREATE INDEX idx_match_history_played_at ON match_history(played_at DESC);
CREATE INDEX idx_items_game_id ON items(game_id);
CREATE INDEX idx_items_categoria_slug ON items(categoria_slug);
CREATE INDEX idx_inventories_player_id ON inventories(player_id);
CREATE INDEX idx_game_profiles_game_score ON game_profiles(game_id, highest_score DESC);
CREATE INDEX idx_game_reviews_game_id ON game_reviews(game_id);
CREATE INDEX idx_reports_status ON reports(status);
