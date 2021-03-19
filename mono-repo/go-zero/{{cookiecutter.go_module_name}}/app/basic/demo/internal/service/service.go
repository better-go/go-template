package service

import (
    "context"

    "{{cookiecutter.go_module_name}}/app/basic/demo/internal/domain/{{cookiecutter.app_name}}"
    "{{cookiecutter.go_module_name}}/app/basic/demo/proto/api"
    "{{cookiecutter.go_module_name}}/app/basic/demo/proto/config"

    "github.com/better-go/pkg/log"
)

// Service 对应 gRPC API 映射
type Service struct {
    Demo *{{cookiecutter.app_name}}.Scope // 推送服务
}

func New(cfg *config.Config) *Service {
    return &Service{
        Demo: {{cookiecutter.app_name}}.NewScope(cfg),
    }
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *Service) Hello(ctx context.Context, req *api.HelloReq) (resp *api.HelloResp, err error) {
    resp = new(api.HelloResp)
    resp.Reply = req.Message + ", reply to you."

    // call:
    err = m.Demo.Demo.Hello(ctx, req.Message)
    log.Infof("service.Demo.Hello done, req=%+v, resp=%+v", req, resp)

    return resp, err
}
