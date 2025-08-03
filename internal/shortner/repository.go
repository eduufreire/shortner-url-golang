package shortner

import (
	"database/sql"
	"fmt"
)

type repository struct {
	db *sql.DB
}

func Repository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(data Shortner) error {
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

func (r *repository) GetByHash(hash string) (*Shortner, error) {
	stmt, err := r.db.Prepare("select * from shortner where hash_url = ?")
	if err != nil {
		return nil, fmt.Errorf("Error found in statement")
	}

	rows, err := stmt.Query(hash)
	if err != nil {
		return nil, fmt.Errorf("Error found in insert")
	}

	shortnerUrl := Shortner{}
	if rows.Next() {
		rows.Scan(&shortnerUrl.HashUrl, &shortnerUrl.OriginalUrl, &shortnerUrl.Clicks)
	}
	return &shortnerUrl, nil
}
