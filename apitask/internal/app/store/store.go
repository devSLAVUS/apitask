package store

import(
  "database/sql"
  _ "github.com/lib/pq"
)
type Store struct {
  config *Store
  db *sql.DB
}
func New(config *Config) *Store {
  return &Config {
    config: config,
  }
}

func (s *Store) Open() error {
  db, err := sql.Open("postgres", s.config.DatabaseURL)
  if err != nil {
    return err
  }
  if err := db.Ping; err != nil {
    return nil
  }
  s.db = db
  return nil
}
func (s *Store) Close() {
  s.db.Close()
}
