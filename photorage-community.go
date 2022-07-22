package main

import (
	"photorage-community/web/rest"

	"photorage-community/adaptor"

	"github.com/gin-gonic/gin"
)

const SERVER_PORT string = ":8880"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	rest.AddAlbumCommentRouter(r)
	// TODO : 계속해서 controller가 추가되어가는 장소

	return r
}

func main() {
	// 이것보다 더 좋은 방법은 없는가?
	go adaptor.CommentDeleteConsumer()

	r := SetupRouter()
	r.Run(SERVER_PORT)
}
