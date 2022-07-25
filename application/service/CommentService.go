package service

import (
	"photorage-community/application/port/in"
	"photorage-community/repository"
	"photorage-community/web/dto"
)

type CommentService struct{}

var (
	commentRepo = repository.CommentRepository{}
)

// 인터페이스를 전부 구현했는지 확인하는 방안
var _ in.CommentServicePort = &CommentService{}

func (commentService *CommentService) GetAlbumComment(albumSeq int64) []dto.CommentDTO {
	commentList, _ := commentRepo.FindCommentList(albumSeq)

	commentDTOList := make([]dto.CommentDTO, len(commentList))

	for i := 0; i < len(commentList); i++ {
		commentDTOList[i] = dto.ConvertCommentEntityToDTO(&commentList[i])
	}

	return commentDTOList
}

func (commentService *CommentService) DeleteAlbumComment(albumSeq int64) {
}
