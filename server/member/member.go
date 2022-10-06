package member

import (
	"archi/server/context"
	member_sql "archi/sql/member"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type relationInput struct {
	a_Id   int    `json:a_id`
	b_Id   int    `json:b_id`
	action string `json:action`
}

func Setup(e *echo.Echo) {
	e.POST("/relations/:relation-id/action", func(c echo.Context) error {
		relation_id, _ := strconv.Atoi(c.Param("relation-id"))
		var relation_input relationInput
		if err := c.Bind(&relation_input); err != nil {
			return err
		}

		if relation_input.action == "ignore" {
			cc := c.(*context.Context)
			member_sql.SetIgnoreById(cc.Db, relation_id)
		}

		return c.JSON(http.StatusCreated, relation_input)
	})
}
