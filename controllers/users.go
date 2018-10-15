package controllers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetMembers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "users")
	}
}

func GetMember(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "user")
	}
}

func AddMember(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, H{
			"created": c.Request().PostForm["user"],
		})
	}
}

func UpdateMember(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, H{
			"updated", c.Request().GetBody,
		})
	}
}

func deleteMember(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}