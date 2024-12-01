## life game in golang

![go](https://img.shields.io/badge/go-222?style=for-the-badge&logo=go)
![docker](https://img.shields.io/badge/docker-222?style=for-the-badge&logo=docker)

#### with docker

``` sh
docker image build --no-cache -f Dockerfile -t go-life:v1.0.0 .
```

``` sh
docker run -it --rm go-life:v1.0.0
```

#### run locally

``` sh
go run app/main.go
```
