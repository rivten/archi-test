package context

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	Db *sqlx.DB
}
