package member

import (
	"github.com/jmoiron/sqlx"
)

type Relation struct {
}

type referral struct {
	id      int `db:id`
	a_id    int `db:a_id`
	b_id    int `db:b_id`
	ignored int `db:ignored`
}

func SetIgnoreById(db *sqlx.DB, relation_id int) {
	_, err := db.NamedExec(
		`UPDATE relation SET ignored = 1 WHERE id = :relation_id;`,
		relation_id,
	)
}
