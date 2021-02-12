package usersUseCase_test

import (
	"testing"
	"time"

	"emperror.dev/errors"
	"github.com/stretchr/testify/assert"

	userEntities "cv-todo-app/internal/user/entities"
	userMocks "cv-todo-app/internal/user/entities/mocks"
	usersUseCase "cv-todo-app/internal/user/useCase"
)

func TestNew(t *testing.T) {
	type args struct {
		repo userEntities.Repository
	}
	r := userMocks.Repository{}
	tests := []struct {
		name string
		args args
		want userEntities.UseCase
	}{
		{
			name: "Success",
			args: args{
				repo: r,
			},
			want: usersUseCase.New(r),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := usersUseCase.New(tt.args.repo)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_uc_Create(t *testing.T) {
	type fields struct {
		repo userEntities.Repository
	}
	type args struct {
		input userEntities.CreateInput
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser userEntities.User
		wantErr  bool
	}{
		{
			name: "Success",
			fields: fields{
				repo: userMocks.Repository{StoreFn: func(input userEntities.User) (user userEntities.User, err error) {
					return userEntities.User{
						Id:          "1",
						Name:        "test",
						Mail:        "test@gmail.com",
						PhoneNumber: "79999999999",
						Created:     time.Time{},
						LastUpdated: time.Time{},
					}, err
				}},
			},
			args: args{
				input: userEntities.CreateInput{
					PhoneNumber: "79999999999",
					Mail:        "test@gmail.com",
					Name:        "test",
				},
			},
			wantUser: userEntities.User{
				Id:          "1",
				Name:        "test",
				Mail:        "test@gmail.com",
				PhoneNumber: "79999999999",
				Created:     time.Time{},
				LastUpdated: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "Validdation err",
			fields: fields{
				repo: userMocks.Repository{StoreFn: func(input userEntities.User) (user userEntities.User, err error) {
					return userEntities.User{}, err
				}},
			},
			args: args{
				input: userEntities.CreateInput{
					PhoneNumber: "",
					Mail:        "test@gmail.com",
					Name:        "test",
				},
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
		{
			name: "Repo err",
			fields: fields{
				repo: userMocks.Repository{StoreFn: func(input userEntities.User) (user userEntities.User, err error) {
					return userEntities.User{}, errors.Sentinel("problem")
				}},
			},
			args: args{
				input: userEntities.CreateInput{
					PhoneNumber: "79999999999",
					Mail:        "test@gmail.com",
					Name:        "test",
				},
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersUseCase.New(tt.fields.repo)
			gotUser, err := u.Create(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}

func Test_uc_DeleteOne(t *testing.T) {
	type fields struct {
		repo      userEntities.Repository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				repo:      userMocks.Repository{DeleteOneFn: func(id string) (err error) {
					return nil
				}},
			},
			args: args{
				id: "1",
			},
			wantErr: false,
		},
		{
			name: "Validation err",
			fields: fields{
				repo:      userMocks.Repository{DeleteOneFn: func(id string) (err error) {
					return nil
				}},
			},
			args: args{
				id: "",
			},
			wantErr: true,
		},
		{
			name: "Repo err",
			fields: fields{
				repo:      userMocks.Repository{DeleteOneFn: func(id string) (err error) {
					return errors.Sentinel("problem")
				}},
			},
			args: args{
				id: "1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersUseCase.New(tt.fields.repo)
			if err := u.DeleteOne(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_uc_GetOne(t *testing.T) {
	type fields struct {
		repo      userEntities.Repository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser userEntities.User
		wantErr  bool
	}{
		{
			name: "Success",
			fields: fields{
				repo:      userMocks.Repository{GetOneFn: func(id string) (user userEntities.User, err error) {
					return userEntities.User{
						Id:          "1",
						Name:        "test",
						Mail:        "test@gmail.com",
						PhoneNumber: "79999999999",
						Created:     time.Time{},
						LastUpdated: time.Time{},
					},nil
				}},
			},
			args: args{
				id: "1",
			},
			wantUser: userEntities.User{
				Id:          "1",
				Name:        "test",
				Mail:        "test@gmail.com",
				PhoneNumber: "79999999999",
				Created:     time.Time{},
				LastUpdated: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "Validation err",
			fields: fields{
				repo:      userMocks.Repository{GetOneFn: func(id string) (user userEntities.User, err error) {
					return userEntities.User{},nil
				}},
			},
			args: args{
				id: "",
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
		{
			name: "Repo err",
			fields: fields{
				repo:      userMocks.Repository{GetOneFn: func(id string) (user userEntities.User, err error) {
					return userEntities.User{},errors.Sentinel("")
				}},
			},
			args: args{
				id: "1",
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersUseCase.New(tt.fields.repo)
			gotUser, err := u.GetOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}

func Test_uc_Update(t *testing.T) {
	type fields struct {
		repo      userEntities.Repository
	}
	type args struct {
		input userEntities.UpdateInput
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser userEntities.User
		wantErr  bool
	}{
		{
			name: "Success",
			fields: fields{
				repo: userMocks.Repository{UpdateFn: func(input userEntities.User) (user userEntities.User, err error) {
					return userEntities.User{
						Id:          "1",
						Name:        "test",
						Mail:        "test@gmail.com",
						PhoneNumber: "79999999999",
						Created:     time.Time{},
						LastUpdated: time.Time{},
					}, err
				}},
			},
			args: args{
				input: userEntities.UpdateInput{
					Id:          "1",
					Name:        "test",
					Mail:        "test@gmail.com",
					PhoneNumber: "79999999999",
				},
			},
			wantUser: userEntities.User{
				Id:          "1",
				Name:        "test",
				Mail:        "test@gmail.com",
				PhoneNumber: "79999999999",
				Created:     time.Time{},
				LastUpdated: time.Time{},
			},
			wantErr: false,
		},
		{
			name: "Validation err",
			fields: fields{
				repo: userMocks.Repository{UpdateFn: func(input userEntities.User) (user userEntities.User, err error) {
					return userEntities.User{
						Id:          "1",
						Name:        "test",
						Mail:        "test@gmail.com",
						PhoneNumber: "79999999999",
						Created:     time.Time{},
						LastUpdated: time.Time{},
					}, err
				}},
			},
			args: args{
				input: userEntities.UpdateInput{
					Id:          "",
					Name:        "test",
					Mail:        "test@gmail.com",
					PhoneNumber: "79999999999",
				},
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
		{
			name: "repo err",
			fields: fields{
				repo: userMocks.Repository{UpdateFn: func(input userEntities.User) (user userEntities.User, err error) {
					return userEntities.User{}, errors.Sentinel("problem")
				}},
			},
			args: args{
				input: userEntities.UpdateInput{
					Id:          "1",
					Name:        "test",
					Mail:        "test@gmail.com",
					PhoneNumber: "79999999999",
				},
			},
			wantUser: userEntities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := usersUseCase.New(tt.fields.repo)
			gotUser, err := u.Update(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}
