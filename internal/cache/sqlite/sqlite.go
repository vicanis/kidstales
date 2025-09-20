package sqlite

import (
	"database/sql"
	"fmt"
	"kidstales/internal/model"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ttl = 14 * 24 * time.Hour
)

type dbCache struct {
	db *sql.DB
}

func NewDBCache() *dbCache {
	db, err := sql.Open("sqlite3", "/db/cache.db")
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS cache (
		key TEXT PRIMARY KEY,
		data BLOB,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err = db.Exec(createTableSQL); err != nil {
		log.Fatalf("unable to create cache table: %v", err)
	}

	return &dbCache{db: db}
}

func (d *dbCache) Get(key string) ([]byte, error) {
	var createdAt time.Time

	var dbData []byte
	err := d.db.QueryRow(`SELECT data, created_at FROM cache WHERE key = ?`, key).Scan(&dbData, &createdAt)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("cache get failed: %w", err)
		}

		return nil, model.ErrNotFound
	}

	if time.Since(createdAt) < ttl {
		return dbData, nil
	}

	return nil, model.ErrNotFound
}

func (d *dbCache) Set(key string, data []byte) error {
	_, err := d.db.Exec(`INSERT OR REPLACE INTO cache (key, data) VALUES (?, ?)`, key, data)
	if err != nil {
		return fmt.Errorf("cache set failed: %w", err)
	}

	return nil
}

func (d *dbCache) Clean() error {
	log.Printf("Clean(): started")

	res, err := d.db.Exec(`DELETE FROM cache WHERE created_at < ? LIMIT 1000`, time.Now().Add(-ttl))
	if err != nil {
		return fmt.Errorf("clean failed: %w", err)
	}

	count, _ := res.RowsAffected()

	fmt.Printf("Clean(): %d rows deleted, starting vacuum", count)

	_, err = d.db.Exec("VACUUM")
	if err != nil {
		return fmt.Errorf("vacuum failed: %w", err)
	}

	fmt.Printf("Clean(): vacuum completed")

	return nil
}
