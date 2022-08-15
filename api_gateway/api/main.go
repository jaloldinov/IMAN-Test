package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// IMAN
	// @title Test task from IMAN Invest
	// @version 1.1
	// @description This task is given for internship position in IMAN
	// @contact.name Jaloldinov Omadbek
	// @contact.email jaloldinovuz@gmail.com
	// @contact.url https://www.linkedin.com/in/jaloldinovuz
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/jaloldinov/IMAN-Updated/api_gateway/api/docs"
	"github.com/jaloldinov/IMAN-Updated/api_gateway/config"
	"github.com/jaloldinov/IMAN-Updated/api_gateway/pkg/logger"
	"github.com/jaloldinov/IMAN-Updated/api_gateway/services"

	v1 "github.com/jaloldinov/IMAN-Updated/api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")

	// First service insert handler

	// Secomd Service (GRUD)
	apiV1.GET("/post/list", handlerV1.ListPosts)
	apiV1.GET("/post/:post_id", handlerV1.GetPost)
	apiV1.PUT("/post/:post_id", handlerV1.UpdatePost)
	apiV1.DELETE("/post/:post_id", handlerV1.DeletePost)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}
