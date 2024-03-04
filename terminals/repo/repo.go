package repo

import (
	"math/rand"
	"sqlitebenchmark/terminals"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repo, error) {
	err := migrate(db)
	if err != nil {
		return nil, err
	}
	return &Repo{
		db: db,
	}, nil
}

func (r *Repo) Get(terminal string) (addr string, ok bool) {
	var t terminals.Terminal
	if err := r.db.Where("terminal = ?", terminal).First(&t).Error; err != nil {
		return
	}
	return t.Addr, true
}

func (r *Repo) Set(params terminals.SetParams) (response terminals.SetResponse, err error) {
	t := terminals.Terminal(params)
	if err = r.db.Create(&t).Error; err != nil {
		return
	}
	response = terminals.SetResponse(t)
	return
}

func (r *Repo) RandomInsert(n int) error {
	records := make([]terminals.Terminal, n)
	for i := 0; i < n; i++ {
		terminal := randStringBytes(10)
		addr := randStringBytes(10)
		record := terminals.Terminal{Terminal: terminal, Addr: addr}
		records = append(records, record)
	}
	r.db.Create(&records)
	return nil
}

func migrate(db *gorm.DB) (err error) {
	if err = db.Exec(`CREATE TABLE IF NOT EXISTS terminals (terminal TEXT UNIQUE, addr TEXT);
	CREATE INDEX IF NOT EXISTS idx_terminals_terminal ON terminals(terminal);`).Error; err != nil {
		return err
	}
	return nil
}

var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
