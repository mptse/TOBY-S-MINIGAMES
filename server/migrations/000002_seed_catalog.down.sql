-- Reversion del seed inicial de categorias y minijuegos

DELETE FROM game_categories
WHERE game_id IN (SELECT game_id FROM games WHERE name_slug IN ('trivia', 'memory'))
  AND category_id IN (SELECT category_id FROM categories WHERE category_slug IN ('cultura-general', 'memoria', 'agilidad-mental'));

DELETE FROM games WHERE name_slug IN ('trivia', 'memory');
DELETE FROM categories WHERE category_slug IN ('cultura-general', 'memoria', 'agilidad-mental');
