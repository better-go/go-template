package router

import (
    "fmt"
    "net/http"
    "time"

    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/api"

    "github.com/better-go/pkg/log"
    "github.com/gin-gonic/gin"
)

/*
- gin docs:
	- https://gin-gonic.com/zh-cn/docs/examples/

- gin 路由写法:
	- POST 取参数: https://github.com/swaggo/swag/blob/master/example/celler/controller/accounts.go
	- 请求参数:
		- 获取: https://gin-gonic.com/zh-cn/docs/examples/query-and-post-form/
		- 数据绑定+校验: https://gin-gonic.com/zh-cn/docs/examples/binding-and-validation/
		- 数据校验: https://gin-gonic.com/zh-cn/docs/examples/custom-validators/
	- 请求中开新 goroutines
		- https://gin-gonic.com/zh-cn/docs/examples/goroutines-inside-a-middleware/
*/

// @Summary 传参测试
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/post [post]
func post(c *gin.Context) {
    id := c.Query("id")
    page := c.DefaultQuery("page", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")

    log.Warnw("failed to fetch URL",
        "url", "http://example.com",
        "attempt", 3,
        "backoff", time.Second,
    )

    fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)

    c.JSON(200, gin.H{
        "message": "from post test",
    })
}

// @Summary 测试 echo 响应
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/ping [get]
func ping(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong, echo test",
    })
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////

// demo.hello()
// @Summary echo func
// @Produce  json
// @Param message query string true "EventID"
// @Param name query string false "RoomID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/demo/hello [post]
func hello(ctx *gin.Context) {
    var req *api.HelloReq

    // validate:
    if err := ctx.ShouldBind(&req); err != nil {
        log.Error("invalid request params")
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    //
    // ok: do logic
    //
    resp, err := gs.Hello(ctx, req)
    log.Debugf("demo.Hello api done, req=%+v, resp=%+v, err=%v", req, resp, err)

    //
    // return ok
    //
    ctx.JSON(200, gin.H{
        "message": "echo hello done",
    })
}


