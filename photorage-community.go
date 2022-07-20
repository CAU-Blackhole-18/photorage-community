package main

import (
	"photorage-community/web/rest"

	"github.com/gin-gonic/gin"
)

const SERVER_PORT string = ":8080"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	rest.AddAlbumCommentRouter(r)
	// TODO : 계속해서 controller가 추가되어가는 장소

	return r
}

func main() {
	r := SetupRouter()
	r.Run(SERVER_PORT)
}
