package userPostgresRepository_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	test "cv-todo-app/internal/testing"
	userEntities "cv-todo-app/internal/user/entities"
	userPostgresRepository "cv-todo-app/internal/user/repository/postgres"
)

func cleanup(p *pgxpool.Pool) error {
	query := `DELETE FROM users;
			ALTER SEQUENCE id_seq RESTART WITH 1;`
	_, err := p.Exec(context.TODO(), query)
	return err
}

func seedOne(p *pgxpool.Pool) error {
	query := `INSERT INTO users (name,mail,phone) VALUES ('test','test','test');`
	_, err := p.Exec(context.TODO(), query)
	return err
}

func TestRepository(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	rand.Seed(time.Now().UTC().UnixNano())
	p := test.GetPool()

	err := cleanup(p)
	if err != nil {
		t.Errorf("Could not prepare table: %v", err)
	}

	r := userPostgresRepository.New(p)

	t.Run("Accounts postgresRepository", func(t *testing.T) {
		// It's always a good idea to build all non-unit tests to be able to work in parallel.
		// Thanks to that, your tests will be always fast and you will not be afraid to add more tests because of slowdown.
		t.Run("testNewRepository", func(t *testing.T) {
			testNew(t)
		})
		t.Run("testStore", func(t *testing.T) {
			testStore(t, r)
			cleanup(p)
		})
		t.Run("testDelete", func(t *testing.T) {
			if err = seedOne(p); err != nil {
				t.Errorf("could not prepare: %s", err.Error())
			}
			testDeleteOne(t, r)
			cleanup(p)
		})

		t.Run("testGetOne", func(t *testing.T) {
			if err = seedOne(p); err != nil {
				t.Errorf("could not prepare: %s", err.Error())
			}
			testGetOne(t, r)
			cleanup(p)
		})

		t.Run("testUpdate", func(t *testing.T) {
			if err = seedOne(p); err != nil {
				t.Errorf("could not prepare: %s", err.Error())
			}
			testUpdate(t, r)
			cleanup(p)
		})

	})

}

func testNew(t *testing.T) {
	type args struct {
		client *pgxpool.Pool
	}

	p := test.GetPool()
	tests := []struct {
		name string
		args args
		want userEntities.Repository
	}{
		{
			name: "Success",
			args: args{
				client: p,
			},
			want: userPostgresRepository.New(p),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := userPostgresRepository.New(tt.args.client)
			assert.Equal(t, tt.want, got)
		})
	}
}

func testDeleteOne(t *testing.T, r userEntities.Repository) {

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				id: "1",
			},
			wantErr: false,
		},
		{
			name: "Not found",
			args: args{
				id: "10",
			},
			wantErr: true,
		},
		{
			name: "Invalid id",
			args: args{
				id: "INVALID",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := r.DeleteOne(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func testGetOne(t *testing.T, r userEntities.Repository) {

	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		wantUser userEntities.User
		wantErr  bool
	}{
		{
			name: "Succeess",
			args: args{
				id: "1",
			},
			wantUser: userEntities.User{
				Id:          "1",
				Name:        "test",
				Mail:        "test",
				PhoneNumber: "test",
			},
			wantErr: false,
		},
		{
			name: "Not found",
			args: args{
				id: "10",
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
		{
			name: "Bad request",
			args: args{
				id: "INVALID",
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotUser, err := r.GetOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.wantUser.Id, gotUser.Id,"id mismatch")
			assert.Equalf(t, tt.wantUser.Name, gotUser.Name,"name mismatch")
			assert.Equalf(t, tt.wantUser.Mail, gotUser.Mail,"mail mismatch")
			assert.Equalf(t, tt.wantUser.PhoneNumber, gotUser.PhoneNumber,"phone mismatch")
		})
	}
}

func testStore(t *testing.T, r userEntities.Repository) {

	type args struct {
		input userEntities.User
	}
	tests := []struct {
		name     string
		args     args
		wantUser userEntities.User
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				input: userEntities.User{
					Name:        "test",
					Mail:        "test",
					PhoneNumber: "test",
				},
			},
			wantUser: userEntities.User{
				Id:          "1",
				Name:        "test",
				Mail:        "test",
				PhoneNumber: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotUser, err := r.Store(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.wantUser.Id, gotUser.Id,"id mismatch")
			assert.Equalf(t, tt.wantUser.Name, gotUser.Name,"name mismatch")
			assert.Equalf(t, tt.wantUser.Mail, gotUser.Mail,"mail mismatch")
			assert.Equalf(t, tt.wantUser.PhoneNumber, gotUser.PhoneNumber,"phone mismatch")
		})
	}
}

func testUpdate(t *testing.T, r userEntities.Repository) {

	type args struct {
		input userEntities.User
	}
	tests := []struct {
		name     string
		args     args
		wantUser userEntities.User
		wantErr  bool
	}{
		{
			name: "Success",
			args: args{
				input: userEntities.User{
					Id:          "1",
					Name:        "test2",
					Mail:        "test2",
					PhoneNumber: "test2",
				},
			},
			wantUser: userEntities.User{
				Id:          "1",
				Name:        "test2",
				Mail:        "test2",
				PhoneNumber: "test2",
			},
			wantErr: false,
		},
		{
			name: "Not found",
			args: args{
				input: userEntities.User{
					Id:          "1111",
					Name:        "test2",
					Mail:        "test2",
					PhoneNumber: "test2",
				},
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
		{
			name: "Bad request",
			args: args{
				input: userEntities.User{
					Id:          "INVALID",
					Name:        "test2",
					Mail:        "test2",
					PhoneNumber: "test2",
				},
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotUser, err := r.Update(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.wantUser.Id, gotUser.Id,"id mismatch")
			assert.Equalf(t, tt.wantUser.Name, gotUser.Name,"name mismatch")
			assert.Equalf(t, tt.wantUser.Mail, gotUser.Mail,"mail mismatch")
			assert.Equalf(t, tt.wantUser.PhoneNumber, gotUser.PhoneNumber,"phone mismatch")
		})
	}
}
