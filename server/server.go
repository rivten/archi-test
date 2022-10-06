package main

import (
	"archi/server/ad"
	"archi/server/context"
	"archi/server/member"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Setup sql connection in Context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db, _ := sqlx.Open("sqlite", "db.db")
			cc := &context.Context{c, db}
			return next(cc)
		}
	})

	ad.Setup(e)
	member.Setup(e)

	e.Logger.Fatal(e.Start(":1313"))
}
