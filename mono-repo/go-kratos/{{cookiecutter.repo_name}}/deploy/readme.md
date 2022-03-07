

# Deploy: 

- 部署/运维相关配置脚本


## introduction: 

- local: 本地开发环境: 基于 docker + docker-compose
- staging: 预发布环境: 基于 docker or k8s
- production: 生产环境: 基于  



## reference: 

> docker: 

- https://docs.docker.com/engine/install/
- https://docs.docker.com/engine/install/ubuntu/
- https://docs.docker.com/compose/install/

> Kubernetes (K8s): 

- https://github.com/kubernetes/kubernetes
- https://github.com/k3s-io/k3s


> terraform: 

- https://github.com/hashicorp/terraform
- examples: 
    - https://github.com/hashicorp/terraform-guides
        - https://github.com/hashicorp/terraform-guides/blob/master/infrastructure-as-code/hashistack/quick-start/terraform-aws/README.md
    - https://github.com/futurice/terraform-examples
        - https://github.com/futurice/terraform-examples/blob/master/aws/aws_ec2_ebs_docker_host/README.md
    
> docker-compose + dockerfile:

- https://github.com/service-mesh/devops
- https://github.com/service-mesh/dockerfile



## local: 

- 本地开发环境: 
- 启动 docker 容器: 

```bash

cd deploy/local

make run.xxx

```


## staging:

- 预发布环境: 


## production: 

- 生产环境: 
