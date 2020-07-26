package config

const (
    // 对应服务目录名称
    RepoName = "{{cookiecutter.repo_name}}"
    AppName  = "{{cookiecutter.app_name}}"
)

// Config: root config name, for toml file objects
type Config struct {
    //
    // common:
    //
    DB    *DBUnit    // db
    Cache *CacheUnit // cache
    MQ    *MQUnit    // mq
    HTTP  *HttpUnit  // http
    RPC   *RpcUnit   // rpc

    //
    // biz:
    //
    Biz *BizUnit //
}

/////////////////////////////////////////////////////////////////////////////////////
// app  biz config unit:
/////////////////////////////////////////////////////////////////////////////////////

// 业务本身配置参数:
type BizUnit struct {
    Demo *DemoBiz
}

type DemoBiz struct {
}

/////////////////////////////////////////////////////////////////////////////////////
//  middleware config unit:
/////////////////////////////////////////////////////////////////////////////////////

// db:
type DBUnit struct {
    Demo *DemoDB
}

type DemoDB struct {
    Url        string
    Host       string
    Port       string
    DBName     string
    DBPassword string
}

// cache:
type CacheUnit struct {
    Demo *DemoCache
}

type DemoCache struct {
    Url     string
    Timeout string
}

// mq:
type MQUnit struct {
    Demo *DemoMQ
}

type DemoMQ struct {
    Url     string
    Timeout string
}

/////////////////////////////////////////////////////////////////////////////////////
//  http config unit:
/////////////////////////////////////////////////////////////////////////////////////

type HttpUnit struct {
    Demo *DemoHttp
}

type DemoHttp struct {
    AppKey    string
    AppSecret string
}

/////////////////////////////////////////////////////////////////////////////////////
//  rpc config unit:
/////////////////////////////////////////////////////////////////////////////////////

type RpcUnit struct {
    Demo *DemoRpc
}

type DemoRpc struct {
    AppKey    string
    AppSecret string
}
