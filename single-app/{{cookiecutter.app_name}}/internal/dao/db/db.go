package db

import (
    "{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/proto/config"
)

/*
所有 HTTP 接口对接

*/

type Dao struct {
    // biz group:
    Demo *DemoDB // demo db
}

func New(cfg *config.DBUnit) *Dao {
    return &Dao{
        Demo: newDemoDB(cfg),
    }
}

func (m *Dao) Close() {
    return
}
