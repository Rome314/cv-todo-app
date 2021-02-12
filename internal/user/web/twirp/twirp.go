package userTwirpDelivery

import (
	"context"

	"github.com/twitchtv/twirp"

	userEntities "cv-todo-app/internal/user/entities"
	userAdapters "cv-todo-app/internal/user/entities/adapters"
	usersRPC "cv-todo-app/rpc/user"
)

type delivery struct {
	uc      userEntities.UseCase
	adapter userAdapters.Proto
}

func (d *delivery) Create(ctx context.Context, req *usersRPC.CreateRequest) (resp *usersRPC.CreateResponse, err error) {
	resp = &usersRPC.CreateResponse{}
	if err = req.Validate(); err != nil {
		err = twirp.NewError(twirp.Malformed, err.Error())
		return
	}

	toCreate := d.adapter.CreateInputFromProto(req)
	created, err := d.uc.Create(toCreate)
	if err != nil {
		err = getTiwrpErr(err)
		return
	}

	resp.User = d.adapter.UserToProto(created)
	return

}

func (d *delivery) GetOne(ctx context.Context, req *usersRPC.GetOneRequest) (resp *usersRPC.GetOneResponse, err error) {
	resp = &usersRPC.GetOneResponse{}
	if err = req.Validate(); err != nil {
		err = twirp.NewError(twirp.Malformed, err.Error())
		return
	}

	found, err := d.uc.GetOne(req.Id)
	if err != nil {
		err = getTiwrpErr(err)
		return
	}

	resp.User = d.adapter.UserToProto(found)
	return
}

func (d *delivery) Update(ctx context.Context, req *usersRPC.UpdateRequest) (resp *usersRPC.UpdateResponse, err error) {
	resp = &usersRPC.UpdateResponse{}
	if err = req.Validate(); err != nil {
		err = twirp.NewError(twirp.Malformed, err.Error())
		return
	}

	toUpdate := d.adapter.UpdateInputFromProto(req)
	updated, err := d.uc.Update(toUpdate)
	if err != nil {
		err = getTiwrpErr(err)
		return
	}

	resp.User = d.adapter.UserToProto(updated)
	return
}

func (d *delivery) Delete(ctx context.Context, req *usersRPC.DeleteRequest) (resp *usersRPC.DeleteResponse, err error) {
	resp = &usersRPC.DeleteResponse{}
	if err = req.Validate(); err != nil {
		err = twirp.NewError(twirp.Malformed, err.Error())
		return
	}

	err = d.uc.DeleteOne(req.Id)
	if err != nil {
		err = getTiwrpErr(err)
		return
	}
	return
}

func New(uc userEntities.UseCase) usersRPC.UsersService {
	return &delivery{
		uc:      uc,
		adapter: newAdapter(),
	}
}
