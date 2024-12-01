## life game in golang

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
