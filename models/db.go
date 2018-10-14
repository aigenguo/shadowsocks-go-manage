package models

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/ss5?charset=utf8mb4")

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	initsql := `
		CREATE TABLE members (
  			id int(10) unsigned NOT NULL AUTO_INCREMENT,
  			port int(11) NOT NULL,
  			password varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  			name varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  			created_at timestamp NULL DEFAULT NULL,
  			updated_at timestamp NULL DEFAULT NULL,
  			PRIMARY KEY (id),
  		  	UNIQUE KEY members_port_unique (port)
		) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

		CREATE TABLE migrations (
  			migration varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  			batch int(11) NOT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

		CREATE TABLE users (
  			id int(10) unsigned NOT NULL AUTO_INCREMENT,
  			email varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  			password varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  			remember_token varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  			created_at timestamp NULL DEFAULT NULL,
  			updated_at timestamp NULL DEFAULT NULL,
  			PRIMARY KEY (id),
  			UNIQUE KEY users_email_unique (email)
		) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
	`
	_, err := db.Exec(initsql)
	if err != nil {
		panic(err)
	}
}
