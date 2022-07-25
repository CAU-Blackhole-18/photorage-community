package out

import (
	"photorage-community/domain"
)

type CommentRepositoryPort interface {
	FindCommentList(albumSeq int64) ([]domain.Comment, error)
	DeleteAlbumComment(albumSeq int64) error
	ModifyComment(albumSeq int64, commentSeq int64) (domain.Comment, error)
	SaveComment(albumSeq int64, commentSeq int64) (domain.Comment, error)
}
