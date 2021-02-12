package userPostgresRepository

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	userEntities "cv-todo-app/internal/user/entities"
)

func Test_userSql_ToUser(t *testing.T) {
	type fields struct {
		Id          sql.NullInt32
		Name        sql.NullString
		Mail        sql.NullString
		PhoneNumber sql.NullString
		Created     sql.NullTime
		LastUpdated sql.NullTime
	}
	tests := []struct {
		name   string
		fields fields
		want   userEntities.User
	}{
		{
			name: "Success",
			fields: fields{
				Id: sql.NullInt32{
					Int32: 1,
					Valid: true,
				},
				Name: sql.NullString{
					String: "test",
					Valid:  true,
				},
				Mail: sql.NullString{
					String: "test",
					Valid:  true,
				},
				PhoneNumber: sql.NullString{
					String: "test",
					Valid:  true,
				},
				Created: sql.NullTime{
					Time:  time.Unix(111,111),
					Valid: true,
				},
				LastUpdated: sql.NullTime{
					Time:  time.Unix(111,111),
					Valid: true,
				},
			},
			want: userEntities.User{
				Id:          "1",
				Name:        "test",
				Mail:        "test",
				PhoneNumber: "test",
				Created:     time.Unix(111,111),
				LastUpdated: time.Unix(111,111),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userSql{
				Id:          tt.fields.Id,
				Name:        tt.fields.Name,
				Mail:        tt.fields.Mail,
				PhoneNumber: tt.fields.PhoneNumber,
				Created:     tt.fields.Created,
				LastUpdated: tt.fields.LastUpdated,
			}
			if got := u.ToUser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
