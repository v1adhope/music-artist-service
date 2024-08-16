# Improvisation, experimenting, and practice

See tags for goals.

There's nothing about clean architecture. For example why you shouldn't use getters and setters with models

# Might be

- [ ] DI container
- [ ] Handle negative test cases (refactor)
- [ ] Observer
- [ ] Dev build (configs, deployment tools)
- [ ] Cache
- [ ] Representation
- [ ] Taskfile for `grpcurl`

# Tree

```bash
.
├── Dockerfile
├── Makefile
├── README.md
├── api
│   └── proto
│       └── v1
│           └── artist.proto
├── certs
│   ├── ca_cert.pem
│   └── ca_key.pem
├── cmd
│   └── service
│       └── main.go
├── compose.yaml
├── db
│   └── migrations
│       ├── 000001_init.down.sql
│       └── 000001_init.up.sql
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── app.go
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   └── grpc
│   │       └── v1
│   │           ├── artist.go
│   │           ├── controller.go
│   │           ├── interceptors.go
│   │           ├── utils.go
│   │           └── v1_test.go
│   ├── entities
│   │   ├── artist.go
│   │   └── errors.go
│   ├── objectvalues
│   │   ├── objectvalues.go
│   │   └── objectvalues_test.go
│   ├── testhelpers
│   │   ├── grpc.go
│   │   └── postgres.go
│   └── usecases
│       ├── artist.go
│       ├── constructor.go
│       ├── errors.go
│       ├── infrastructure
│       │   ├── repositories
│       │   │   ├── artist.go
│       │   │   ├── artist_test.go
│       │   │   ├── constructor.go
│       │   │   ├── dtos.go
│       │   │   └── interfaces.go
│       │   └── validation
│       │       ├── validation.go
│       │       └── validation_test.go
│       ├── interfaces.go
│       └── mocks
│           └── reposmocks.go
├── pkg
│   ├── api
│   │   └── proto
│   │       └── v1
│   │           ├── artist.pb.go
│   │           └── artist_grpc.pb.go
│   ├── logger
│   │   └── logger.go
│   └── postgresql
│       └── posgresql.go
├── scripts
│   └── generate_protos.sh
└── tests
    └── grpcurl_artist.sh
```
