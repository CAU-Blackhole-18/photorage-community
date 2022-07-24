package service

import (
	"photorage-community/repository"
	"photorage-community/web/dto"
)

func GetAlbumComment(albumSeq int64) []dto.CommentDTO {
	commentList := repository.FindCommentList(albumSeq)

	commentDTOList := make([]dto.CommentDTO, len(commentList))

	for i := 0; i < len(commentList); i++ {
		commentDTOList[i] = dto.ConvertCommentEntityToDTO(&commentList[i])
	}
	return commentDTOList
}

func DeleteAlbumComment(albumSeq int64) {

}
