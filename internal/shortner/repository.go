package shortner

import (
	"database/sql"
	"fmt"
)

type shortnerRepository struct {
	db *sql.DB
}

func NewShortnerRepository(db *sql.DB) ShortnerRepository {
	return &shortnerRepository{
		db: db,
	}
}

func (r *shortnerRepository) Save(data Shortner) error {
	stmt, err := r.db.Prepare("insert into shortner (hash_url, original_url, clicks, user_id) values (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("Error found in statement")
	}

	_, err = stmt.Exec(data.HashUrl, data.OriginalUrl, data.Clicks, data.UserID)
	if err != nil {
		return fmt.Errorf("Error found in insert")
	}
	return nil
}

func (r *shortnerRepository) GetByHash(hash string) (*Shortner, error) {
	stmt, err := r.db.Prepare("select * from shortner where hash_url = ?")
	if err != nil {
		return nil, fmt.Errorf("Error found in statement")
	}

	shortnerUrl := Shortner{}
	stmt.QueryRow(hash).Scan(&shortnerUrl.HashUrl, &shortnerUrl.OriginalUrl, &shortnerUrl.Clicks, &shortnerUrl.UserID)
	return &shortnerUrl, nil
}
