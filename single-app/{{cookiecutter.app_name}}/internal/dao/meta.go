package dao

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/async"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/cache"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/db"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/http"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/mq"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/dao/rpc"
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

// 数据层资源收敛入口:
type MetaResource struct {
    Async *async.Dao // async task handler
    DB    *db.Dao    // db layer
    Cache *cache.Dao // cache layer
    MQ    *mq.Dao    // mq layer
    HTTP  *http.Dao  // http layer
    RPC   *rpc.Dao   // rpc layer
}

func NewMetaResource(cfg *config.Config) *MetaResource {
    return &MetaResource{
        Async: async.New(),
        DB:    db.New(cfg.DB),
        Cache: cache.New(cfg.Cache),
        MQ:    mq.New(cfg.MQ),
        HTTP:  http.New(cfg.HTTP),
        RPC:   rpc.New(cfg.RPC),
    }
}

func (m *MetaResource) Close() {
    m.Async.Close()
    m.DB.Close()
    m.Cache.Close()
    m.MQ.Close()
    m.DB.Close()
    m.HTTP.Close()
    m.RPC.Close()
}
