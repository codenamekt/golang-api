Mgo Api 
=======
Building a simple rest interface on top of mongo.

- Master Branch
  - [![Build Status](https://travis-ci.org/codenamekt/mgo-api.svg?branch=master)](https://travis-ci.org/codenamekt/mgo-api)
  - [![CI Status](https://circleci.com/gh/codenamekt/mgo-api/tree/master.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/codenamekt/mgo-api)
- Develop Branch
  - [![Build Status](https://travis-ci.org/codenamekt/mgo-api.svg?branch=develop)](https://travis-ci.org/codenamekt/mgo-api)
  - [![CI Status](https://circleci.com/gh/codenamekt/mgo-api/tree/develop.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/codenamekt/mgo-api)

Getting Started
===============
```go
git clone git@github.com:codenamekt/mgo-api.git
docker-compose up
```

TODO
=====

- Deploy to Google Cloud

Test coverage
=============

`go test -coverprofile=c.out`

`go tool cover -html=c.out`
