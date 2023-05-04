package users

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

const (
	getUserById = `SELECT u.user_id, u."user_name", u."last_name", u.email, u."role" FROM gouser u		
	WHERE u.user_id = $1
	LIMIT 1`
	getUsers = `SELECT u.user_id, u."user_name", u."last_name", u.email, u."role" FROM gouser u`
)

type UserQuery interface {
	GetUsers(ctx context.Context, id int32) ([]*User, error)
}

func NewUserQuery(db *sqlx.DB) UserQuery {
	return &querier{db}
}

type querier struct {
	db *sqlx.DB
}

func (q *querier) GetUsers(ctx context.Context, id int32) ([]*User, error) {
	var err error
	var rows *sqlx.Rows
	if id != 0 {
		rows, err = q.db.QueryxContext(ctx, getUserById, id)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	} else {
		rows, err = q.db.QueryxContext(ctx, getUsers)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

	}
	defer rows.Close()

	var res []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.UserId, &user.UserName, &user.LastName, &user.Email, &user.Role)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}

	return res, err
}
