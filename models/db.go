package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Remember string `json:"remember"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type Member struct {
	ID int `json:"id"`
	Port int `json:"port"`
	Password string `json:"password"`
	Name string `json:"name"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type MemberCollection struct {
	Members []Member `json:"items"`
}

func GetMembers(db *sql.DB) MemberCollection {
	ssql := "SELECT * FROM members"
	rows, err := db.Query(ssql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := MemberCollection{}
	for rows.Next() {
		member := Member{}
		err = rows.Scan(&member.ID,&member.Port,&member.Password,&member.Name,&member.Created,&member.Updated)
		if err != nil {
			panic(err)
		}
		result.Members = append(result.Members,member)
	}
	return result
}

func