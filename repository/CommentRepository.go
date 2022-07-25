package repository

import (
	"photorage-community/application/port/out"
	"photorage-community/domain"
)

type CommentRepository struct {
}

// 인터페이스를 전부 구현했는지 확인하는 방안
var _ out.CommentRepositoryPort = &CommentRepository{}

// 임시로 넣어둔 Comment Array 반환
// TODO : Error 처리 필요
func (cr *CommentRepository) FindCommentList(albumSeq int64) ([]domain.Comment, error) {
	commentList := make([]domain.Comment, 10)
	for i := 0; i < 10; i++ {
		commentList[i] = domain.Comment{AlbumSeq: albumSeq + int64(i), Content: "Random"}
	}
	return commentList, nil
}

func (cr *CommentRepository) DeleteAlbumComment(albumSeq int64) error {
	return nil
}

func (cr *CommentRepository) ModifyComment(albumSeq int64, commentSeq int64) (domain.Comment, error) {
	comment := domain.Comment{}
	return comment, nil
}

func (cr *CommentRepository) SaveComment(albumSeq int64, commentSeq int64) (domain.Comment, error) {
	comment := domain.Comment{}
	return comment, nil
}
