
# cookiecutter-go-mono-repo

- boilerplate
- go 微服务项目目录生成工具
- mono repo 范式.

## brief:

- https://github.com/better-go/cookiecutter-go
    - 生成 mono repo 整个项目目录
- https://github.com/better-go/cookiecutter-go-app
    - 生成单个微服务目录
- 先创建 `mono-repo`, 在切换到 `app/` 下, 创建`单个微服务目录`.


## install:

- https://github.com/cookiecutter/cookiecutter
- https://cookiecutter.readthedocs.io/en/1.7.2/installation.html
- 基于 cookiecutter 实现.

```bash

# 安装 cookiecutter:
brew install cookiecutter

# 在当前目录下, 创建mono repo 项目根目录:
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="mono-repo"

# 在当前目录下, 创建微服务目录:
cookiecutter https://github.com/better-go/cookiecutter-go.git --directory="single-app"

cookiecutter https://github.com/better-go/cookiecutter-go-app.git

```


## multi template:

- https://cookiecutter.readthedocs.io/en/1.7.2/advanced/directories.html
- 支持一个 git 下, 存放多个`目录创建模板`.

```bash

cookiecutter https://github.com/user/repo-name.git --directory="directory1-name"


```


## create:


- https://cookiecutter.readthedocs.io/en/1.7.2/first_steps.html
- 创建自己的项目目录模板结构

- `cookiecutter.json`


```json

{
    "full_name": "Audrey Roy",
    "email": "audreyr@gmail.com",
    "project_name": "Complexity",
    "repo_name": "complexity",
    "project_short_description": "Refreshingly simple static site generator.",
    "release_date": "2013-07-10",
    "year": "2013",
    "version": "0.1.1"
}

```

- repo_name


```bash

mkdir {{cookiecutter.directory_name}}
mkdir {{cookiecutter.repo_name}}

```


## mono repo structure:


```bash

[~/cookiecutter-go/{{cookiecutter.repo_name}}] [master]

-> % tree . -L 5
.
├── Makefile
├── README.MD
├── app
│   ├── basic
│   │   ├── user
│   │   │   ├── auth
│   │   │   ├── identity
│   │   │   │   ├── Makefile
│   │   │   │   ├── cmd
│   │   │   │   ├── configs
│   │   │   │   ├── docs
│   │   │   │   ├── internal
│   │   │   │   ├── proto
│   │   │   │   └── readme.md
│   │   │   ├── permission
│   │   │   ├── readme.md
│   │   │   └── role
│   │   └── {{cookiecutter.basic_app_name}}
│   │       ├── cmd
│   │       │   └── main.go
│   │       ├── configs
│   │       │   └── configs.toml
│   │       ├── docs
│   │       ├── internal
│   │       │   ├── dao
│   │       │   ├── domain
│   │       │   └── service
│   │       └── proto
│   │           ├── api
│   │           ├── config
│   │           └── model
│   ├── biz
│   │   └── {{cookiecutter.biz_app_name}}
│   │       └── cmd
│   │           └── main.go
│   └── std
│       ├── Makefile
│       ├── proto
│       │   ├── config
│       │   │   └── config.proto
│       │   └── error
│       │       └── code.proto
│       └── readme.md
├── deploy
│   ├── local
│   │   └── Makefile
│   └── staging
│       └── Makefile
├── go.mod
├── infra
│   └── tool
├── pkg
├── script
└── tmp

39 directories, 15 files


```


## ref:

- https://github.com/audreyr/cookiecutter-pypackage
- https://github.com/pydanny/cookiecutter-django
    - https://github.com/pydanny/cookiecutter-django/blob/master/cookiecutter.json
- https://github.com/golang-standards/project-layout
- https://github.com/micro/micro
    - https://github.com/micro/micro/blob/master/client/cli/new/new.go
- https://github.com/adobe/go-starter
- https://github.com/gatsbyjs/gatsby-starter-default
