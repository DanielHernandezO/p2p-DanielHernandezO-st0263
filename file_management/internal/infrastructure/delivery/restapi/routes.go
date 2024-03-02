package restapi

import "github.com/gin-gonic/gin"

func mapUrls(router *gin.Engine, dependencies *Dependencies) {
	groupPath := router.Group("/file/managment/v1")
	groupPath.POST("/add/file", dependencies.fileHandler.Post)
	groupPath.GET("/get/:fileId", dependencies.fileHandler.GetFile)
}
