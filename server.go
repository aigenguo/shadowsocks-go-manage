package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/ss5?charset=utf8mb4&multiStatements=true")

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sqlBytes, err := ioutil.ReadFile("init.sql")
	if err != nil {
		panic(err)
	}

	initsql := string(sqlBytes)
	_, err = db.Exec(initsql)
	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := initDb()
	migrate(db)
	e.GET("/v1/users", func(c echo.Context) error {
		return c.JSON(200, "Get all users!")
	})
	e.GET("/v1/users/:id", func(c echo.Context) error {
		return c.JSON(200, "Get user")
	})
	e.POST("/v1/users", func(c echo.Context) error {
		return c.JSON(200, "Add user")
	})
	e.PUT("/v1/users", func(c echo.Context) error {
		return c.JSON(200, "Modify user")
	})
	e.DELETE("/v1/users/:id", func(c echo.Context) error {
		return c.JSON(200, "Delete user")
	})
	e.Logger.Fatal((e.Start(":1323")))
}
