package in

import (
	"photorage-community/web/dto"
)

type CommentServicePort interface {
	GetAlbumComment(albumSeq int64) []dto.CommentDTO
	DeleteAlbumComment(albumSeq int64)
}
