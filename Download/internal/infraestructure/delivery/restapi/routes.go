package restapi

import "github.com/gin-gonic/gin"

func mapUrls(router *gin.Engine, dependencies *Dependencies) {
	groupPath := router.Group("/download")
	groupPath.POST("/save/file", dependencies.saveHandler.Save)
}
