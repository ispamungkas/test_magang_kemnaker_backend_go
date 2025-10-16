CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  domicile VARCHAR
);

INSERT INTO users (name, domicile)
VALUES
  ('Ilham Pratama', 'Tangerang'),
  ('Rizky Saputra', 'Jakarta'),
  ('Dewi Lestari', 'Bandung'),
  ('Ahmad Fauzi', 'Yogyakarta'),
  ('Nisa Amelia', 'Surabaya');