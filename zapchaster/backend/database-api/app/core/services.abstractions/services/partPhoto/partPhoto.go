package partPhoto

import (
	"database-api/app/core/contracts/dtos/partPhotoDtos"
)

type IPartPhotoService interface {
	AddPhoto(photos []partPhotoDtos.PartPhotoDto) int
}
