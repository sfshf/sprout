# Sprout

:-)


A web-app demo implemented by Go!


:-)

# REFERENCE:


## Go programming language:

- [how to use go!](https://github.com/golang/go/wiki/Learn)

- [The Go Blog](https://docs.studygolang.com/blog/)
- [The Laws of Reflection](https://docs.studygolang.com/blog/laws-of-reflection)
- [Introducing HTTP Tracing](https://docs.studygolang.com/blog/http-tracing)
- [Defer, Panic, and Recover](https://docs.studygolang.com/blog/defer-panic-and-recover)
- [Working with Errors in Go 1.13](https://docs.studygolang.com/blog/go1.13-errors)
- [Effective Go](https://docs.studygolang.com/doc/effective_go.html)


### Go Memory

- [Go basically never frees heap memory back to the operating system](https://utcc.utoronto.ca/~cks/space/blog/programming/GoNoMemoryFreeing)

- [Go Memory Management](https://povilasv.me/go-memory-management/)


### `Concurrency` and `Context`:

- [Share Memory By Communicating](https://docs.studygolang.com/blog/codelab-share)
- [Bell Labs and CSP Threads](https://swtch.com/~rsc/thread/)
- [Communicating Sequential Processes](http://www.usingcsp.com/)
- [Go Concurrency Patterns: Pipelines and cancellation](https://docs.studygolang.com/blog/pipelines)
- [Go Concurrency Patterns: Context](https://docs.studygolang.com/blog/context)
- [Go Concurrency Patterns: Timing out, moving on](https://docs.studygolang.com/blog/concurrency-timeouts)


## HTTP and Web:

- [w3.org](https://fetch.spec.whatwg.org/)

- [MDN web docs](https://developer.mozilla.org/en-US/docs/Web/HTTP)

- [Content negotiation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Content_negotiation)

- [Web Development](http://hellobmw.com/category/web-development/)


### HTTP Authentication

- [HTTP Authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication)


#### Cookie And Session

- [HTTP State Management Mechanism](https://tools.ietf.org/html/rfc6265)
- [Using HTTP cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies)

- [Gin](https://github.com/gin-gonic/gin)
    - [building-go-web-applications-and-microservices-using-gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)
    - [7天用Go从零实现Web框架Gee教程](https://geektutu.com/post/gee.html)
    - [gin-contrib](https://github.com/gin-contrib/)

- [github.com/gorilla/mux](https://github.com/gorilla/mux)

- [Software ArchitecturePatternsUnderstanding Common ArchitecturePatterns and When to Use Them](https://www.oreilly.com/programming/free/files/software-architecture-patterns.pdf)
    - [软件架构入门](http://www.ruanyifeng.com/blog/2016/09/software-architecture.html)

- [The twelve factors](https://www.12factor.net/)

- [Session Management Cheat Sheet](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Session_Management_Cheat_Sheet.md)
- [Seesion Fixation Vulnerability in Web-based Applications](http://www.acrossecurity.com/papers/session_fixation.pdf)
- [Insufficient Session-ID Length Description](https://owasp.org/www-community/vulnerabilities/Insufficient_Session-ID_Length)


## CORS

- [Cross-Origin Resource Sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
- [gorilla handlers -- cors](https://github.com/gorilla/handlers/blob/3e030244b4ba0480763356fc8ca0ade6222e2da0/cors.go#L168)
- [github.com/rs/cors](https://github.com/rs/cors)
- [github.com/gin-contrib/cors](https://github.com/gin-contrib/cors)


## Profiles:

- [go tool pprof](https://wiki.jikexueyuan.com/project/go-command-tutorial/0.12.html)
- [An Introduction to go tool trace](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/)
- [go tool trace](https://making.pusher.com/go-tool-trace/)
- [github.com/google/gops](https://github.com/google/gops/)


## Command line interface:

- [github.com/urfave/cli](https://github.com/urfave/cli)


## Unit test:

- [github.com/stretchr/testify](https://github.com/stretchr/testify)


## Config:

- [Viper](https://github.com/spf13/viper)

- [Multiconfig](https://github.com/koding/multiconfig)


## Errors:

- [github.com/pkg/errors](https://github.com/pkg/errors)


## Logger:

- [github.com/rs/zerolog](https://github.com/rs/zerolog)

- [github.com/uber-go/zap](https://github.com/uber-go/zap)
- [github.com/go-logr/zapr](https://github.com/go-logr/zapr)

- [Log Files](http://httpd.apache.org/docs/current/logs.html)

- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)

- [A more minimal logging API for Go](https://github.com/go-logr/logr)
- [dave.cheney.net/2015/11/05/lets-talk-about-logging](https://dave.cheney.net/2015/11/05/lets-talk-about-logging)

## Data Validate:

- [github.com/go-playground/validator](https://github.com/go-playground/validator)


## Publish and Subscribe:

- [go-nsq](https://github.com/nsqio/go-nsq)


## JWT:

- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [JSON Web Token (JWT)](http://self-issued.info/docs/draft-ietf-oauth-json-web-token.html)
    - [JSON Web Token 入门教程](http://www.ruanyifeng.com/blog/2018/07/json_web_token-tutorial.html)


## RBAC:


- [Casbin](https://github.com/casbin/casbin)
    - [Access Control Policy Specification Language Based on Metamodel (in Chinese)（必读）](http://www.jos.org.cn/1000-9825/5624.htm)
    - [Basic Role-Based HTTP Authorization in Go with Casbin](https://zupzup.org/casbin-http-role-auth/)
    - [以角色为基础的存取控制](https://baike.tw.lvfukeji.com/wiki/RBAC)
    - [Role Based Access Control](https://csrc.nist.gov/projects/role-based-access-control/)
    - [Role Based Access Control FAQs](https://csrc.nist.gov/projects/role-based-access-control/faqs)
    - [Role-Based Access Controls](https://csrc.nist.gov/publications/detail/conference-paper/1992/10/13/role-based-access-controls)
        - [The paper: Role-Based Access Controls](https://csrc.nist.gov/CSRC/media/Publications/conference-paper/1992/10/13/role-based-access-controls/documents/ferraiolo-kuhn-92.pdf)
    - [Access Control](https://csrc.nist.gov/Topics/Security-and-Privacy/identity-and-access-management/access-control)

## SQL:

### MySQL

- [MySQL 8.0 Reference Manual](https://dev.mysql.com/doc/refman/8.0/en/)

- [Gorm v2](https://github.com/go-gorm/gorm)
    - [Gorm guides](https://gorm.io/docs/index.html)
    - [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

- [A Guide Database Performance for Developers](https://use-the-index-luke.com/)

## NoSQL:

- [What is NoSQL?](https://www.mongodb.com/nosql-explained)
- [Types of NoSQL Databases](https://www.mongodb.com/scale/types-of-nosql-databases)

### Mongo

- [Mongo](https://docs.mongodb.com/manual/)
- [JSON and BSON](https://www.mongodb.com/json-and-bson)
- [Json Schema](https://json-schema.org/)
- [BSON](http://bsonspec.org/)
- [What is a Document Database?](https://www.mongodb.com/document-databases)
- [Key-Value Databases](https://www.mongodb.com/key-value-database)
- [Manifesto for Agile Software Development](https://agilemanifesto.org/)
- [DB-engines](https://db-engines.com/en/ranking)


### Redis

- [Redis cluster tutorial](https://redis.io/topics/cluster-tutorial)


## Docker:

- [Docker overview](https://docs.docker.com/get-started/overview/)
- [Docker engine api sdk](https://docs.docker.com/engine/api/sdk/)
- [Docker hub](https://hub.docker.com/)


## Others:

### k/v stores:

- [github.com/tidwall/buntdb](https://github.com/tidwall/buntdb)

### json iterator:

- [github.com/json-iterator/go](https://github.com/json-iterator/go)

### dependency injector:

- [github.com/google/wire](https://github.com/google/wire)
- [guestbook sample in Go Cloud](https://github.com/google/go-cloud/tree/master/samples/guestbook)
- [What is dependency injection](https://stackoverflow.com/questions/130794/what-is-dependency-injection)

### api docs builder:

- [What Is OpenAPI?](https://swagger.io/docs/specification/about/)
- [github.com/swaggo/swag](https://github.com/swaggo/swag)

### SOA

- [github.com/go-kit/kit](https://github.com/go-kit/kit)

### CI

- [Travis CI](https://docs.travis-ci.com/)

### Nginx

- [The C10k problem](http://www.kegel.com/c10k.html)
