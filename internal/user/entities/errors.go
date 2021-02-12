package userEntities

import "emperror.dev/errors"

const (
	NotFoundError   = errors.Sentinel("not found")
	InternalError   = errors.Sentinel("internal")
	BadRequestError = errors.Sentinel("bad request")
)
