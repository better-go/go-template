package {{cookiecutter.app_name}}

import (
    "context"

    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao"

    "github.com/better-go/pkg/log"
)

// demo:
type DemoScope struct {
    *dao.MetaResource
}

func newDemoScope(g *dao.MetaResource) *DemoScope {
    return &DemoScope{g}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////

// 消息推送: 异步+重试
func (m *DemoScope) Hello(ctx context.Context, message string) error {
    log.Infof("domain.Demo.Hello done, req=%+v", message)
    return nil
}
