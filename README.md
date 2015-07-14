Mgo Api [![Build Status](https://travis-ci.org/codenamekt/mgo-api.svg?branch=master)](https://travis-ci.org/codenamekt/mgo-api) https://circleci.com/gh/codenamekt/mgo-api/tree/master.svg?style=shield&circle-token=:circle-token
=======
Building a simple rest interface on top of mongo.

Getting Started
===============
```go
git clone git@github.com:codenamekt/mgo-api.git
docker-compose up
```

TODO
=====

- Use GoDeps
- Create CI (Deploy Google Cloud?)
- Coveralls.io

Test coverage
=============

`go test -coverprofile=c.out`

`go tool cover -html=c.out`
