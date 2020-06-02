package http

import (
    "matrix-contrib/app/push/proto/config"

    "github.com/better-go/pkg/com/jpush"
    "github.com/better-go/pkg/log"
)

// 极光推送 API 对接:
type DemoHttp struct {
}

func newDemoHttp(cfg *config.HTTPConfig) *JPushHttp {
    appKey := cfg.JPush.AppKey
    appSecret := cfg.JPush.AppSecret

    return &JPushHttp{
        client: jpush.NewJPush(appKey, appSecret),
    }
}

func (m *DemoHttp) Hello(message string) error {
    log.Debugf("demo.Hello done, req=%+v, err=%v", message)
    return nil
}
