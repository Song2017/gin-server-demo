# go build -a -o ./cbec ./swagger_server/main.go
FROM golang:1.17-alpine AS build

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

COPY ./ /go/src/envoy_utils

WORKDIR /go/src/envoy_utils

RUN go get -d -v ./... \
  && go build -a -o ./envoy_utils_server ./swagger_server/main.go


FROM alpine AS runtime

ENV GIN_MODE=release

COPY --from=build /go/src/envoy_utils ./

EXPOSE 8000

ENTRYPOINT ["./envoy_utils_server"]
