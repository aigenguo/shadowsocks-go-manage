package controllers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
)

func getUsers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "users")
	}
}

func getUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "user")
	}
}

func addUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return 
	}
}