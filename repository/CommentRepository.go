package repository

import (
	"photorage-community/domain"
)

// 임시로 넣어둔 Comment Array 반환
func FindCommentList(albumSeq int64) []domain.Comment {
	commentList := make([]domain.Comment, 10)
	for i := 0; i < 10; i++ {
		commentList[i] = domain.Comment{AlbumSeq: albumSeq + int64(i), Content: "Random"}
	}
	return commentList
}
