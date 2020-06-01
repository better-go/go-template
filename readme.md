
# cookiecutter-go-mono-repo

- boilerplate
- go 微服务项目目录生成工具
- mono repo 范式.


## install:

- https://github.com/cookiecutter/cookiecutter
- https://cookiecutter.readthedocs.io/en/1.7.2/installation.html
- 基于 cookiecutter 实现.

```bash 

# 安装 cookiecutter:
brew install cookiecutter

# 拉模板:
cookiecutter https://github.com/better-go/cookiecutter-go.git

```


## create:


- https://cookiecutter.readthedocs.io/en/1.7.2/first_steps.html

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

- https://github.com/pydanny/cookiecutter-django
    - https://github.com/pydanny/cookiecutter-django/blob/master/cookiecutter.json
- https://github.com/golang-standards/project-layout

