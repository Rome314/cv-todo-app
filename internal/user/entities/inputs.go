package userEntities

type CreateInput struct {
	PhoneNumber string
	Mail string
	Name string
}

type UpdateInput struct {
	Id string
	Name string
	PhoneNumber string
	Mail string
}

