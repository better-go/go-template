package router

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/service"

    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
    gs *service.Service // gs = global service object, for global  use
)

// http api register:
func RegisterRouter(r *gin.Engine, srv *service.Service) {
    //
    // global use
    //
    gs = srv

    // swagger docs: for debug
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // api:
    v1 := r.Group("api/v1")
    {
        v1.POST("/post", post)
        v1.GET("/ping", ping)

        // push: 推送服务
        push := v1.Group("/demo")
        {
            push.POST("/hello", hello) // demo.Hello()
        }
    }

    // debug test:
    debug := r.Group("internal/debug")
    {
        // test:
        debug.GET("/test", func(ctx *gin.Context) {
            ctx.JSON(200, gin.H{
                "message": "debug test",
            })
        })

        // etc:
    }
}
