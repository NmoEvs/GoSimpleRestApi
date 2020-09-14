# Simple Rest API

## [Launch Golang Dev environnement](https://levelup.gitconnected.com/setup-simple-go-development-environment-with-docker-b8b9c0d4e0a8)

docker run --rm -it --name go-restful golang

## Mount Source code

docker run --rm -it --name go-restful \
  -v $PWD:/go/src/github.com/the-evengers/go-restful golang
