CREATE TABLE IF NOT EXISTS links(
                                    id SERIAL PRIMARY KEY,
                                    url TEXT NOT NULL UNIQUE,
                                    short_url VARCHAR(10) NOT NULL UNIQUE
);

CREATE INDEX IF NOT EXISTS idx_short_url ON links (short_url);

