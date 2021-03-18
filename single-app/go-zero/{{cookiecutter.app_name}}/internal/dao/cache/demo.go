package cache

import (
    "github.com/better-go/pkg/log"

    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

// 极光推送 API 对接:
type DemoCache struct {
}

func newDemoCache(cfg *config.CacheUnit) *DemoCache {

    return &DemoCache{
    }
}

func (m *DemoCache) Hello(message string) error {
    log.Debugf("demo.Hello done, req=%+v", message)
    return nil
}
