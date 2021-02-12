package userPostgresRepository

import (
	"database/sql"
	"strconv"

	userEntities "cv-todo-app/internal/user/entities"
)

type userSql struct {
	Id          sql.NullInt32
	Name        sql.NullString
	Mail        sql.NullString
	PhoneNumber sql.NullString
	Created     sql.NullTime
	LastUpdated sql.NullTime
}

func (u userSql) ToUser() userEntities.User {
	id := strconv.Itoa(int(u.Id.Int32))
	return userEntities.User{
		Id:          id,
		Name:        u.Name.String,
		Mail:        u.Mail.String,
		PhoneNumber: u.PhoneNumber.String,
		Created:     u.Created.Time,
		LastUpdated: u.LastUpdated.Time,
	}
}
