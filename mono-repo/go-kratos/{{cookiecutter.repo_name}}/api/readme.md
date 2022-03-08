# api:

- 所有服务的 API:
- protobuf 定义:

## 两种管理方案:

- 集中管理 PROTO API 定义:
    - 方便 client 调用
- 分散管理 PROTO API 定义:
    - 方便单个服务的维护者开发

## 折中方案:

- 写一个 sync 脚本.
- 定期自动同步各服务内的 proto api 定义到 api/ 文件夹下.

