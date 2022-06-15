English | [简体中文](readme-cn.md)
# go-template

- [x] boilerplate, go project starter with `go-zero/gin/go-micro`.
- [x] go microservice (`go-zero/gin/go-micro`) project directory generation tool/scaffold.
- [x] Implementation based on [cookiecutter](https://github.com/cookiecutter/cookiecutter).

## Feature:

- [x] `mono repo`: support creating `mono repo`, the git project root directory.
    - [x] [go-zero](https://github.com/tal-tech/go-zero)
    - [x] [go-kratos](https://github.com/go-kratos/kratos)
    - [x] [gin](https://github.com/gin-gonic/gin)
    - [x] [go-micro](https://github.com/asim/go-micro)
    - [x] [rs]() : rust monorepo
- [x] `single app`: support creating `single app`, a single microservice directory.
    - [x] [go-zero](https://github.com/tal-tech/go-zero)
    - [x] [gin](https://github.com/gin-gonic/gin)
    - [x] [go-micro](https://github.com/asim/go-micro)
- [x] can be used individually or in combination
- [x] usage:
    - [x] Create the `mono-repo` root directory first
    - [x] Switch to `app/`, create `single app` single microservice directory.

## Development Roadmap:

- [x] https://github.com/better-go/cookiecutter-go/projects/1

### Intro:

- `go-zero` microservice:
    - [x] Integrate [air](https://github.com/cosmtrek/air) tool, watch project code, `automatic compilation + hot reload`, make go development smoother (like Python/Js)
    - [x] rpc: integrated grpc
    - [x] log: integrated uber log
    - [x] db: integrate gorm + mysql
    - [x] cache: integrate redis
    - [x] mq: integrate rabbitmq
    - [x] Service Discovery: Integrate etcd
    - [x] cronjob: integrate cronjob
- `go-kratos` microservice:
    - [x] Integrate `go-kratos` scaffolding built-in components
    - [x] integrated air tool
    - [x] integrated grpc/http
- `gin` microservice:
    - [x] Integrate [air](https://github.com/cosmtrek/air) tool, watch project code, `automatic compilation + hot reload`, make go development smoother (like Python/Js)
    - [x] rpc: integrated grpc
    - [x] log: integrated uber log
    - [x] db: integrate gorm + mysql
    - [x] cache: integrate redis
    - [x] mq: integrate rabbitmq
    - [x] Service Discovery: Integrate etcd
    - [x] cronjob: integrate cronjob
- `go-micro` microservice:
    - [x] Integrate [air](https://github.com/cosmtrek/air) tool, watch project code, `automatic compilation + hot reload`, make go development smoother (like Python/Js)
    - [x] rpc: integrated grpc
    - [x] log: integrated uber log
    - [x] db: integrate gorm + mysql

## User Case:

- [x] https://github.com/zeromicro/zeromall
    - Microservice `B2B2C e-commerce platform` based on go-zero

## Quick Start:

### Install:

- https://cookiecutter.readthedocs.io/en/1.7.2/installation.html

```bash

# Mac OS X install cookiecutter:
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
> Elevate according to command line parameters, gradually create.

```bash

cd your-workspace/

# In the current directory, create the root directory of the go-zero mono repo project
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo/go-zero"


# In the current directory, create the gin mono repo project root directory
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo/gin"


# In the current directory, create the root directory of the go-micro mono repo project
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo/go-micro"



```

- 1.2 create `single app` folder:

```bash

cd your-mono-repo-app-create-root/


# In the current directory, create a microservice directory: use go-zero
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-zero"


# In the current directory, create a microservice directory: use gin
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/gin"

# In the current directory, create a microservice directory: use go-micro
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app/go-micro"



```

- tips:
    - Path parameters, do not add suffix : `/`
    - If `go_module_name` is different from `project_folder_name`, be careful not to pass it wrong.

#### 2. create library repo:

- `library-repo`

```bash

cd your-workspace/

# In the current directory, create the mono repo project root directory:
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="library-repo"

```

## Generate Project Folder Structure:

- [x] Mono Repo:


- Brief description of directory structure: a quick look of the overall directory.

```html
-> % tree . -L 3
.
├── app
│ ├── basic // Basic services: usually business-agnostic
│ │ ├── demo // Example:
│ │ └── user // user management:
│  ├── biz // Business Services: Segmentation according to business
│ │ └── member // Member service:
│ │ └── order // order service:
│ │ └── cart // Shopping cart service:
│ └── std // Unified definition within the business:
│ ├── proto // Unified definition of business status code
├── deploy // deployment script, service orchestration
│   ├── local
│   │   └── Makefile
│   └── staging
│       └── Makefile
├── infra // infrastructure
│   └── tool
└── pkg // Utility code base accumulated by the project, not related to business
```

- Detailed description of directory structure:

```html
-> % tree . -L 6
.
├── LICENSE
├── Makefile
├── README.MD
├── app
│ ├── basic // Basic service unit 1: general services unrelated to user management/push/SMS/etc.
│ │ ├── demo // Example of a single microservice:
│ │ │ │ ├── cmd // Single-service startup entry: including multiple startup methods:
│   │   │   │   └── main.go                         // api server, grpc server, job server, admin server 启动
│ │ │ │ ├── configs // External middleware configuration items: db, cache, mq, etc.
│   │   │   │   └── configs.toml
│ │ │ │ ├── docs // single-service documentation
│ │ │ │ ├── internal // business logic (not exposed)
│ │ │ │ │ ├── dao // data layer read/write
│ │ │ │ │ │ │ ├── cache // cache r/w
│   │   │   │   │   ├── db                          //      db crud
│ │ │ │ │ │ ├── http // Call http api outside this service
│ │ │ │ │ │ ├── meta.go // dao resource convergence
│   │   │   │   │   ├── mq                          //      mq r/w
│ │ │ │ │ │ └── rpc // call rpc(gRPC) api outside this service
│ │ │ │ │ ├── domain // In-service business split:
│ │ │ │ │ │ └── demo // Business Unit 1 / Business Unit 2
│ │ │ │ │ └── service // API convergence layer (external convergence internal logic, exposed API: grpc/http/job/admin)
│   │   │   │       └── service.go
│ │ │ │ └── proto // Data definition layer: (exposed to the outside world)
│ │ │ │ ├── api // grpc + http api definition
│   │   │       │   └── api.proto
│ │ │ │ ├── config // config toml mapping model
│   │   │       │   └── config.proto
│ │ │ └── model // internal model
Model │ │ └── model.proto
│ │ └── user // Basic service 2:
│ │ ├── identity // within the service
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
│ ├── biz // Specific business unit:
│ │ │── member // Member service
│   │   │   └── cmd
│   │   │       └── main.go
│ │ └── order // order service
│   │       └── cmd
│   │           └── main.go
│ └── std // Unified definition within the project business (business status code, business error msg number)
│       ├── Makefile
│       ├── proto
│       │   ├── config
│       │   │   └── config.proto
│       │   └── error
│ │ └── code.proto // business status code
│       └── readme.md
├── deploy // Deploy related scripts: dockerfile etc.
│   ├── local
│   │   └── Makefile
│   └── staging
│       └── Makefile
├── go.mod
├── go.sum
├── infra // Plugin dependent services: cli, middleware, etc.
│   └── tool
└── pkg // Utility code base gradually accumulated in the project

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

## Call link description:

- Code call link description: Take the `app/basic/demo` service as an example (top-down call)
    - Cmd startup entry: `app/basic/demo/cmd/main.go`
    - Service 入口: `app/basic/demo/internal/service/service.go`
        - External API Gateway: `app/basic/demo/internal/service/outer/outer.go`
        - Inbound RPC gateway: `app/basic/demo/internal/service/inner/inner.go`
        - Internal Admin API Gateway: `app/basic/demo/internal/service/admin/admin.go`
        - Internal Job Gateway: `app/basic/demo/internal/service/job/job.go`
    - Proto data definition:
        --API layers: `app / basic / demo / proto / api / api.proto`
        - Config 层: `app/basic/demo/proto/config/config.proto`
        - Model 层: `app/basic/demo/proto/model/model.proto`
    - Router route registration: `app/basic/demo/internal/router/router.go`
    - Domain business areas:
        - A business unit: `app/basic/demo/internal/domain/demo/demo.go`
    - Dao data layer operations:
        - Meta aggregation (convergence of Dao layer resources): `app/basic/demo/internal/dao/meta.go`
        - DB 层: `app/basic/demo/internal/dao/db/db.go`
        - Cache 层: `app / basic / demo / internal / dao / cache / cache.go`

## User Manual:

- [x] https://github.com/better-go/cookiecutter-go/wiki/manual
- [x] Initialize the microservice directory, how to use it.

## Reference:

- [x] https://github.com/better-go/cookiecutter-go/wiki/reference



## Contributors:

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/hhstore>
            <img src=https://avatars.githubusercontent.com/u/3252130?v=4 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Henry.Huang/>
            <br />
            <sub style="font-size:14px"><b>Henry.Huang</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/ferrisz>
            <img src=https://avatars.githubusercontent.com/u/15015578?v=4 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Ferris/>
            <br />
            <sub style="font-size:14px"><b>Ferris</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/mrmm>
            <img src=https://avatars.githubusercontent.com/u/796467?v=4 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Mourad Maatoug/>
            <br />
            <sub style="font-size:14px"><b>Mourad Maatoug</b></sub>
        </a>
    </td>
</tr>
</table>


