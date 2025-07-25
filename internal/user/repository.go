package user

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

func (r *repository) Save(data RequestDTO) (*int64, error) {
	stmt, err := r.db.Prepare("insert into user(name, email, password) values(?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("deu ruim")
	}

	result, err := stmt.Exec(data.Name, data.Email, data.Password)
	if err != nil {
		return nil, fmt.Errorf("deu ruim")
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("deu ruim")
	}

	return &lastID, nil
}

func (r *repository) GetByID(id int) (*User, error) {
	stmt, err := r.db.Prepare("select * from user where id = ?")
	if err != nil {
		return nil, fmt.Errorf("deu ruim")
	}

	user := User{}
	stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, nil
}

func (r *repository) GetByEmail(email string) (*User, error) {
	stmt, err := r.db.Prepare("select * from user where email = ?")
	if err != nil {
		return nil, fmt.Errorf("deu ruim")
	}

	user := User{}
	stmt.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, nil
}
