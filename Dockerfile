FROM registry.cn-hangzhou.aliyuncs.com/jinli09051005/tools:golang-1.21 AS builder
WORKDIR /app
COPY . .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct
#RUN go mod tidy && go mod vendor && GOOS=linux GOARCH=amd64 go build -o jinli-dijkstra-api cmd/apiserver/main.go
RUN GOOS=linux GOARCH=amd64 go build -o jinli-dijkstra-api cmd/apiserver/main.go

FROM registry.cn-hangzhou.aliyuncs.com/jinli09051005/tools:busybox-latest
WORKDIR /app
COPY --from=builder /app/jinli-dijkstra-api .
