package dto

import (
	"photorage-community/domain"
)

/*
Comment 를 저장하기 위한 구조체
- Go에서는 외부 참조가 필요한 경우 항상 첫글자가 대문자여야 한다
- 접근제어지시자가 대/소문자로 구분된다고 봐야 함
*/
type CommentDTO struct {
	AlbumSeq int64  `json:"albumSeq"`
	Content  string `json:"content"`
}

// 화살표의 방향이 안으로 흐르기 위해서는, DTO가 Entity를 바라보고 변환해야 함
func ConvertCommentEntityToDTO(c *domain.Comment) CommentDTO {
	return CommentDTO{AlbumSeq: c.AlbumSeq, Content: c.Content}
}
