English | [简体中文](readme.md)


# Deploy:

- Deployment/O&M related configuration scripts


## introduction:

- local: local development environment: based on docker + docker-compose
- staging: pre-release environment: based on docker or k8s
- production: production environment: based on



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

- Local development environment:
- Start the docker container:

```bash

cd deploy/local

make run.xxx

```


## staging:

- Pre-release environment:


## production:

- Production Environment:
