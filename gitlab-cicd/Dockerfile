# FROM golang:1.15
# RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

FROM registry.cn-chengdu.aliyuncs.com/yangsj/go-build-alpine:latest AS builder
WORKDIR /code
COPY . /code
RUN go fmt && go mod tidy
RUN go build main.go

FROM alpine:latest
WORKDIR /code
COPY --from=builder /code/main /code/main
CMD [ "/code/main" ]