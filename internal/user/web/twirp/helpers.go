package userTwirpDelivery

import (
	"emperror.dev/errors"
	"github.com/twitchtv/twirp"

	userEntities "cv-todo-app/internal/user/entities"
)

func getTiwrpErr(e error) error {

	var code twirp.ErrorCode

	switch errors.Cause(e) {
	case userEntities.BadRequestError:
		code = twirp.Malformed
		break
	case userEntities.InternalError:
		code = twirp.Internal
		break
	case userEntities.NotFoundError:
		code = twirp.NotFound
		break
	default:
		code = twirp.Internal
		break

	}
	return twirp.NewError(code, errors.Unwrap(e).Error())
}
