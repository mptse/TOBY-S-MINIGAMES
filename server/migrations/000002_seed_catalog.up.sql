-- Seed inicial idempotente para categorias y minijuegos del catalogo

INSERT INTO categories (name_category, category_desc, category_slug)
VALUES
    ('Cultura General', 'Preguntas de conocimiento general y curiosidades', 'cultura-general'),
    ('Memoria', 'Desafios de retencion visual y concentracion', 'memoria'),
    ('Agilidad Mental', 'Retos rapidos de logica y velocidad cognitiva', 'agilidad-mental')
ON CONFLICT (category_slug) DO NOTHING;

INSERT INTO games (name, name_slug, description, thumbnail_slug, is_active)
VALUES
    ('Trivia', 'trivia', 'Pon a prueba tus conocimientos en rondas de preguntas y respuestas', 'trivia-thumb.png', TRUE),
    ('Memory', 'memory', 'Encuentra todas las parejas de cartas en el menor tiempo posible', 'memory-thumb.png', TRUE)
ON CONFLICT (name_slug) DO NOTHING;

INSERT INTO game_categories (game_id, category_id)
SELECT g.game_id, c.category_id
FROM games g
JOIN categories c ON (
    (g.name_slug = 'trivia' AND c.category_slug IN ('cultura-general', 'agilidad-mental')) OR
    (g.name_slug = 'memory' AND c.category_slug IN ('memoria', 'agilidad-mental'))
)
ON CONFLICT (game_id, category_id) DO NOTHING;
