package service

import (
	"photorage-community/repository"
	"photorage-community/web/dto"
)

type CommentService struct{}

var (
	commentRepo = repository.CommentRepository{}
)

func (commentService *CommentService) GetAlbumComment(albumSeq int64) []dto.CommentDTO {
	commentList, _ := commentRepo.FindCommentList(albumSeq)

	commentDTOList := make([]dto.CommentDTO, len(commentList))

	for i := 0; i < len(commentList); i++ {
		commentDTOList[i] = dto.ConvertCommentEntityToDTO(&commentList[i])
	}
	return commentDTOList
}

func DeleteAlbumComment(albumSeq int64) {

}
