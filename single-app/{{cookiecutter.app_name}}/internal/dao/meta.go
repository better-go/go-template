package dao

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/async"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/http"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

// 数据层资源收敛入口:
type MetaResource struct {
    HTTP  *http.Dao // http layer
    Async *async.Dao
}

//
func NewMetaResource(httpCfg *config.HTTPConfig) *MetaResource {
    return &MetaResource{
        HTTP:  http.New(httpCfg),
        Async: async.New(),
    }
}

func (m *MetaResource) Close() {
    m.HTTP.Close()
    m.Async.Close()
}
