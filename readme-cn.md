# go-template

- [x] boilerplate, go project starter with `go-zero/gin/go-micro`.
- [x] go 微服务(`go-zero/gin/go-micro`)项目目录生成工具/脚手架.
- [x] 基于 [cookiecutter](https://github.com/cookiecutter/cookiecutter) 实现.

## Feature:

- [x] `mono repo`: 支持创建 `mono repo`, git 项目根目录.
    - [x] [go-zero](https://github.com/tal-tech/go-zero)
    - [x] [go-kratos](https://github.com/go-kratos/kratos)
    - [x] [gin](https://github.com/gin-gonic/gin)
    - [x] [go-micro](https://github.com/asim/go-micro)
    - [x] [rs]() : rust monorepo
- [x] `single app`: 支持创建 `single app`, 单个微服务目录.
    - [x] [go-zero](https://github.com/tal-tech/go-zero)
    - [x] [gin](https://github.com/gin-gonic/gin)
    - [x] [go-micro](https://github.com/asim/go-micro)
- [x] 可单个使用, 也可同时组合使用
- [x] usage:
    - [x] 先创建 `mono-repo` 根目录
    - [x] 切换到 `app/` 下, 创建 `single app` 单个微服务目录.

## Development Roadmap:

- [x] https://github.com/better-go/cookiecutter-go/projects/1

### Intro:

- `go-zero` 微服务:
    - [x] 集成 [air](https://github.com/cosmtrek/air) 工具, watch 项目代码, `自动编译 + hot reload`, 让 go 开发更丝滑(类 Python/Js)
    - [x] rpc: 集成 grpc
    - [x] log: 集成 uber log
    - [x] db: 集成 gorm + mysql
    - [x] cache: 集成 redis
    - [x] mq: 集成 rabbitmq
    - [x] 服务发现: 集成 etcd
    - [x] cronjob: 集成 cronjob
- `go-kratos` 微服务:
    - [x] 集成 `go-kratos` 脚手架自带组件
    - [x] 集成 air 工具
    - [x] 集成 grpc/http
- `gin` 微服务:
    - [x] 集成 [air](https://github.com/cosmtrek/air) 工具, watch 项目代码, `自动编译 + hot reload`, 让 go 开发更丝滑(类 Python/Js)
    - [x] rpc: 集成 grpc
    - [x] log: 集成 uber log
    - [x] db: 集成 gorm + mysql
    - [x] cache: 集成 redis
    - [x] mq: 集成 rabbitmq
    - [x] 服务发现: 集成 etcd
    - [x] cronjob: 集成 cronjob
- `go-micro` 微服务:
    - [x] 集成 [air](https://github.com/cosmtrek/air) 工具, watch 项目代码, `自动编译 + hot reload`, 让 go 开发更丝滑(类 Python/Js)
    - [x] rpc: 集成 grpc
    - [x] log: 集成 uber log
    - [x] db: 集成 gorm + mysql

## User Case:

- [x] https://github.com/zeromicro/zeromall
    - 基于 go-zero 的微服务`B2B2C电商平台`

## Quick Start:

### Install:

- https://cookiecutter.readthedocs.io/en/1.7.2/installation.html

```bash

# Mac OS X 安装 cookiecutter:
brew install cookiecutter

# Python + pipx:
brew install pipx
pipx ensurepath
pipx install cookiecutter

# Debian/Ubuntu:
sudo apt-get install cookiecutter

```

### Best Practice:

- [x] choice 1: create `mono-repo` + `single-app`
    - support templates:
        - `go-zero`
        - `gin`
        - `go-micro`
- [x] choice 2: create `library-repo`

#### 1. Mono-Repo + Single-App

- 1.1 create `mono repo`:

> git repo / project root.
>
> 根据命令行参数提升, 逐步创建.

```bash

cd your-workspace/

# 在当前目录下, 创建 go-zero mono repo 项目根目录
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo/go-zero"


# 在当前目录下, 创建 gin mono repo 项目根目录
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo/gin"


# 在当前目录下, 创建 go-micro mono repo 项目根目录
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo/go-micro"



```

- 1.2 create `single app` folder:

```bash

cd your-mono-repo-app-create-root/

# 在当前目录下, 创建微服务目录: 使用 go-zero
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-zero"

# 在当前目录下, 创建微服务目录: 使用 gin
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/gin"

# 在当前目录下, 创建微服务目录: 使用 go-micro
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-micro"



```

- tips:
    - 路径参数, 不要添加后缀 : `/`
    - `go_module_name` 如果 和 `project_folder_name` 不同, 要注意不要传错.

#### 2. create library repo:

- `library-repo`

```bash

cd your-workspace/

# 在当前目录下, 创建mono repo 项目根目录:
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="library-repo"

```

## Generate Project Folder Structure:

- [x] Mono Repo:


- 目录结构简略说明: 对整体目录的 quick look.

```html
-> % tree . -L 3
.
├── app
│   ├── basic            // 基础服务: 通常是业务无关的
│   │   ├── demo         //     示例:
│   │   └── user         //     用户管理:
│   ├── biz              // 业务服务: 根据业务切分
│   │   └── member       //     会员服务:
│   │   └── order        //     订单服务:
│   │   └── cart         //     购物车服务:
│   └── std              // 业务内统一定义:
│       ├── proto        //     统一定义业务状态码
├── deploy               // 部署脚本, 服务编排
│   ├── local
│   │   └── Makefile
│   └── staging
│       └── Makefile
├── infra                // 基础设施
│   └── tool
└── pkg                  // 项目积累的 utility 代码库, 与业务无关
```

- 目录结构详细说明:

```html
-> % tree . -L 6
.
├── LICENSE
├── Makefile
├── README.MD
├── app
│   ├── basic                                       // 基础服务单元1 : 用户管理/推送/短信/等业务无关的通用服务
│   │   ├── demo                                    // 单个微服务示例:
│   │   │   ├── cmd                                 //      单服务启动入口: 包含 多个启动方式:
│   │   │   │   └── main.go                         // api server, grpc server, job server, admin server 启动
│   │   │   ├── configs                             //      外部中间件配置项: db,缓存,mq 等
│   │   │   │   └── configs.toml
│   │   │   ├── docs                                // 单服务自身文档
│   │   │   ├── internal                            //      业务逻辑(不对外暴露)
│   │   │   │   ├── dao                             //      数据层 read/write
│   │   │   │   │   ├── cache                       //      缓存 r/w
│   │   │   │   │   ├── db                          //      db crud
│   │   │   │   │   ├── http                        //      调用本服务之外的 http api
│   │   │   │   │   ├── meta.go                     //      dao 资源收敛
│   │   │   │   │   ├── mq                          //      mq r/w
│   │   │   │   │   └── rpc                         //      调用本服务之外的 rpc(gRPC) api
│   │   │   │   ├── domain                          // 服务内业务拆分:
│   │   │   │   │   └── demo                        //      业务单元1 / 业务单元2
│   │   │   │   └── service                         // API 收敛层(对外收敛内部逻辑, 暴露 API: grpc/http/job/admin)
│   │   │   │       └── service.go
│   │   │   └── proto                               // 数据定义层: (可对外暴露)
│   │   │       ├── api                             //      grpc + http api 定义
│   │   │       │   └── api.proto
│   │   │       ├── config                          //      config toml 映射 model
│   │   │       │   └── config.proto
│   │   │       └── model                           //      内部 model
│   │   │           └── model.proto
│   │   └── user                                    // 基础服务2:
│   │       ├── identity                            //   服务内
│   │       │   ├── Makefile
│   │       │   ├── cmd
│   │       │   │   └── main.go
│   │       │   ├── configs
│   │       │   │   └── configs.toml
│   │       │   ├── docs
│   │       │   ├── internal
│   │       │   │   ├── dao
│   │       │   │   ├── domain
│   │       │   │   └── service
│   │       │   ├── proto
│   │       │   │   ├── api
│   │       │   │   ├── config
│   │       │   │   └── model
│   │       │   └── readme.md
│   │       └── readme.md
│   ├── biz                                         // 具体业务单元:
│   │   │── member                                  // 会员服务
│   │   │   └── cmd
│   │   │       └── main.go
│   │   └── order                                   // 订单服务
│   │       └── cmd
│   │           └── main.go
│   └── std                                        // 项目业务内统一定义(业务状态码, 业务出错 msg 编号)
│       ├── Makefile
│       ├── proto
│       │   ├── config
│       │   │   └── config.proto
│       │   └── error
│       │       └── code.proto                     // 业务状态码
│       └── readme.md
├── deploy                                         // 部署相关脚本: dockerfile 等
│   ├── local
│   │   └── Makefile
│   └── staging
│       └── Makefile
├── go.mod
├── go.sum
├── infra                                          // 插件依赖服务: cli, 中间件等
│   └── tool
└── pkg                                            // 项目中逐步积累的 utility 代码库

46 directories, 24 files
```

- [x] single app:

```bash

-> % tree -L 3 ./single-app/{{cookiecutter.app_name}}
./single-app/{{cookiecutter.app_name}}
├── cmd
│   └── main.go
├── configs
│   └── configs.toml
├── internal
│   └── dao
│       └── db
├── proto
│   └── api
│       └── api.go
└── readme.md

7 directories, 4 files


```

## 调用链路说明:

- 代码调用链路说明: 以 `app/basic/demo` 服务为例(自顶向下调用)
    - Cmd 启动入口: `app/basic/demo/cmd/main.go`
    - Service 入口: `app/basic/demo/internal/service/service.go`
        - 对外 API 网关: `app/basic/demo/internal/service/outer/outer.go`
        - 对内 RPC 网关: `app/basic/demo/internal/service/inner/inner.go`
        - 对内 Admin API 网关: `app/basic/demo/internal/service/admin/admin.go`
        - 对内 Job 网关: `app/basic/demo/internal/service/job/job.go`
    - Proto 数据定义:
        - API 层: `app/basic/demo/proto/api/api.proto`
        - Config 层: `app/basic/demo/proto/config/config.proto`
        - Model 层: `app/basic/demo/proto/model/model.proto`
    - Router 路由注册: `app/basic/demo/internal/router/router.go`
    - Domain 业务领域:
        - 某业务单元: `app/basic/demo/internal/domain/demo/demo.go`
    - Dao 数据层操作:
        - Meta 聚合(收敛 Dao 层资源): `app/basic/demo/internal/dao/meta.go`
        - DB 层: `app/basic/demo/internal/dao/db/db.go`
        - Cache 层: `app/basic/demo/internal/dao/cache/cache.go`

## User Manual:

- [x] https://github.com/better-go/cookiecutter-go/wiki/manual
- [x] 初始化好微服务目录, 如何使用.

## Reference:

- [x] https://github.com/better-go/cookiecutter-go/wiki/reference
