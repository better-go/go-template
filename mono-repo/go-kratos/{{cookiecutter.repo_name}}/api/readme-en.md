English | [简体中文](readme.md)
# api:

- API for all services:
- protobuf definition:

## Two management options:

- Centralized management of PROTO API definitions:
    - It is convenient for client to call
- Decentralized management PROTO API definition:
    - Facilitate development by maintainers of a single service

## Compromise:

- Write a sync script.
- Periodically and automatically synchronize the proto api definitions in each service to the api/ folder.

