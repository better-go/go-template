English | [简体中文](readme.md)

# event queue:

- Generic Queue Management: Incident Response

## Application scenarios:

- Some backend services: you don't need to open MQ in the service, you can call it through a simple http api
- Front-end: Asynchronous capabilities directly through HTTP API.
- Similar to graphQL, the DB capability is opened to the front end. Here, the MQ capability is opened to the front end.
- Signature mechanism: Anti-attack
    - Messages in MQ, add verification, messages with invalid signatures, discard (prevent forgery)
    - Signature rules: api secret, timestamp, nonce, msg_uuid, data. Anti-replay attack (consumers need to cache msg_uuid, deduplication)
- API authentication method:
    - http header: token authentication

- auth authentication:
    - User authentication: user auth
    - Service authentication: service auth
- Permission Control:
    - User permission: user permission
    - Service permission: service permission


## run:

- run server:

```bash

make run.auto.reload


```

- http api test:

```bash

make curl.pub.msg

```

## API List:

- `app/basic/queue/internal/router/router.go`

- Queue write messages: http://localhost:8080/queue/publish/
    - POST
    - Makefile has curl example


## Function:

### Queue management:

- [ ] queue creation
- [ ] Queue auth-key generation/authentication
- [ ] queue destroy

### Producer:

- [ ] Queue Write: HTTP API/ GRPC
- [ ] Queue read: HTTP API polling / websocket listening
- [ ] Pass both `task` and `task handler function` as parameters to the server.
-

### Consumer: (worker)

- [ ] Queue Write: HTTP API/ GRPC
- [ ] Queue read: HTTP API polling / websocket listening
- [ ] task handler

### task handler function:

- Generic task handler method
- Summarize some commonly used task processing methods:
    - Asynchronously record data to db, and return processing success
    - Asynchronous send mobile msg
    - Asynchronous data analysis: generate reports
    - Asynchronous...
- Take standardized methods as parameters and provide them to the outside world
- Similar to `AWS Lambda`: https://aws.amazon.com/cn/lambda/

