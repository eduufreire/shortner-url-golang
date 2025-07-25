package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func CreateDatabase() *sql.DB {
	conn, err := sql.Open("sqlite", "./shortner.db")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Successfully database created")

	tables := `
		create table if not exists shortner (
			hash_url text not null,
			original_url text not null,
			clicks integer,
			user_id,
			primary key(hash_url),
			foreign key(user_id) references user(rowid)
		);

		create table if not exists user (
			id integer primary key autoincrement,
			name text not null,
			email text not null,
			password text not null
		);
	`

	_, err = conn.Exec(tables)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully tables created")
	return conn
}
