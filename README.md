Use Tool

    $ go get -u github.com/ory/dockertest/v3 -> For SQL Test
    $ go get github.com/jmoiron/sqlx [Lib SQL]
    $ go get -u github.com/gin-gonic/gin [API]
    $ go mod init REVAMP-PHP-GO


Testing

    $ go get github.com/smartystreets/goconvey

Database

- MYSQL (LOCAL)
- Docker (MYSQL LOCAL)
  docker run --name enterprise022 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=qazse4321 -d mysql:8.0

Hexagonal Architecture  
The hexagonal architecture, or ports and adapters architecture, is an architectural pattern used in software design. It aims at creating loosely coupled application components that can be easily connected to their software environment by means of ports and adapters. This makes components exchangeable at any level and facilitates test automation

Path

    |── Internal.
        |── datasources.
            |── sqlDB
                |── config.go
        |── domain
            |── logic
            |── modules
                |── account.go
            |── ports
                |── repositoryPort.go
                |── servicePort.go
            |── services
                |── services.go
        |── handler
            |── http.go
    |── mock (working on it)
    |── main.go
