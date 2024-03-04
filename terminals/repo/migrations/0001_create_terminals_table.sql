CREATE TABLE IF NOT EXISTS terminals (terminal TEXT UNIQUE, addr TEXT);
CREATE INDEX IF NOT EXISTS idx_terminals_terminal ON terminals(terminal);