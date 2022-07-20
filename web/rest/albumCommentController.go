package rest

import (
	"net/http"
	"photorage-community/domain"
	"photorage-community/service"
	"photorage-community/web/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddAlbumCommentRouter(router *gin.Engine) {

	albumRouter := router.Group("/album")
	{
		// TODO : 이 부분을 Service Layer로 옮겨야 한다.
		albumRouter.GET("/:albumSeq/comment", GetAlbumComment)
		albumRouter.POST("/:albumSeq/comment", CreateAlbumComment)
	}
}

func GetAlbumComment(c *gin.Context) {
	// int64로 변환
	albumSeq, _ := strconv.ParseInt(c.Param("albumSeq"), 10, 64)

	c.JSON(http.StatusOK, service.GetAlbumComment(albumSeq))
}

func CreateAlbumComment(c *gin.Context) {
	albumId := c.Param("albumId")
	var commentRequestJson dto.CommentDTO
	if err := c.ShouldBindJSON(&commentRequestJson); err == nil {

		// TODO : 이 곳에서 CommentDTO를 Comment 형으로 변환하고, Comment 형의 리시버 메서드 출력해보기
		var newComment domain.Comment
		newComment.AlbumSeq = int64(commentRequestJson.AlbumSeq)
		newComment.Content = commentRequestJson.Content
		newComment.PrintCommentInfo()

		// STUDY : 이와 같은 방법도 가능. key Name을 명시해주는 것이 좋음
		newComment2 := domain.Comment{AlbumSeq: int64(commentRequestJson.AlbumSeq), Content: commentRequestJson.Content}
		newComment2.PrintCommentInfo()

		// STUDY : 인터페이스를 이용해 보았음
		var newComment3 domain.Animal = domain.Comment{AlbumSeq: int64(commentRequestJson.AlbumSeq), Content: commentRequestJson.Content}
		newComment3.Run()

		c.JSON(http.StatusOK, gin.H{
			"albumId_from_path":    albumId,
			"albumId_from_request": commentRequestJson.AlbumSeq,
			"comment_content":      commentRequestJson.Content,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
