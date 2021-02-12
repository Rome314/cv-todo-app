package userPostgresRepository

import (
	"context"
	"strconv"

	"emperror.dev/errors"
	"github.com/jackc/pgx/v4/pgxpool"

	userEntities "cv-todo-app/internal/user/entities"
)

type repo struct {
	client *pgxpool.Pool
}

func (r *repo) IsModuleRepository() {
}

func New(client *pgxpool.Pool) userEntities.Repository {
	return &repo{client: client}
}

func (r *repo) Store(input userEntities.User) (user userEntities.User, err error) {
	user = userEntities.User{}
	query := `INSERT INTO users (name,mail,phone,created,last_updated)
			VALUES ($1,$2,$3,$4,$5)
			RETURNING id,name,mail,phone,created,last_updated;`

	row := r.client.QueryRow(context.TODO(), query,
		input.Name, input.Mail, input.PhoneNumber, input.Created, input.LastUpdated)

	found := userSql{}

	err = row.Scan(&found.Id, &found.Name, &found.Mail, &found.PhoneNumber, &found.Created, &found.LastUpdated)
	if err != nil {
		err = errors.WithDetails(userEntities.InternalError, "scanning", err)
		return
	}

	return found.ToUser(), nil
}

func (r *repo) GetOne(id string) (user userEntities.User, err error) {
	user = userEntities.User{}
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt == 0 {
		err = errors.WithMessagef(userEntities.BadRequestError, "invalid id: %s", id)
		return

	}

	query := `SELECT id,name,mail,phone,created,last_updated FROM users WHERE id = $1;`

	row := r.client.QueryRow(context.TODO(), query, idInt)

	found := userSql{}

	err = row.Scan(&found.Id, &found.Name, &found.Mail, &found.PhoneNumber, &found.Created, &found.LastUpdated)
	if err != nil {
		err = errors.WithDetails(userEntities.InternalError, "scanning", err)
		return
	}

	return found.ToUser(), nil
}

func (r *repo) Update(input userEntities.User) (user userEntities.User, err error) {
	user = userEntities.User{}

	idInt, err := strconv.Atoi(input.Id)
	if err != nil || idInt == 0 {
		err = errors.WithMessagef(userEntities.BadRequestError, "invalid id: %s", input.Id)
		return
	}

	query := `UPDATE users 
				SET name=$1, mail=$2, phone=$3, last_updated=$4
				WHERE id = $5
				RETURNING id,name,mail,phone,created,last_updated;`

	row := r.client.QueryRow(context.TODO(), query, input.Name, input.Mail, input.PhoneNumber, input.LastUpdated, idInt)

	found := userSql{}

	err = row.Scan(&found.Id, &found.Name, &found.Mail, &found.PhoneNumber, &found.Created, &found.LastUpdated)
	if err != nil {
		err = errors.WithDetails(userEntities.InternalError, "scanning", err)
		return
	}

	return found.ToUser(), nil
}

func (r *repo) DeleteOne(id string) (err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt == 0 {
		err = errors.WithMessagef(userEntities.BadRequestError, "invalid id: %s", id)
		return
	}

	query := `DELETE FROM users WHERE id = $1`

	cmd, err := r.client.Exec(context.TODO(), query, idInt)
	if err != nil {
		err = errors.WithDetails(userEntities.InternalError, "executing", err)
		return
	}

	if cmd.RowsAffected() == 0 {
		err = errors.WithDetails(userEntities.NotFoundError, "nothing deleted")
		return
	}

	return nil

}
