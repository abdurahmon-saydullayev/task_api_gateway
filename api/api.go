package api

import (
	_ "task_go_api_gateway/api/docs"
	"task_go_api_gateway/api/handlers"
	"task_go_api_gateway/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {

	// book
	r.POST("/book", h.CreateBook)
	r.GET("/book/:id", h.GetBookByID)
	r.GET("/book", h.GetBookList)
	r.PUT("book/:id", h.UpdateBook)
	r.DELETE("/book/:id", h.DeleteBook)

	//author
	r.POST("/author", h.CreateAuthor)
	r.GET("/author/:id", h.GetAuthorByID)
	r.GET("/author", h.GetAuthorList)
	r.PUT("author/:id", h.UpdateAuthor)
	r.DELETE("author/:id", h.DeleteAuthor)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
