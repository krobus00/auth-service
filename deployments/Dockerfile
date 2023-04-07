FROM golang:1.18.3-alpine as base
WORKDIR /builder
# install upx and make
RUN apk add upx make
ENV GO111MODULE=on CGO_ENABLED=0
COPY go.mod go.sum /builder/
RUN go mod tidy
COPY . .
RUN make build

FROM alpine:3.8
WORKDIR /app
COPY --from=base /builder/bin/auth-service /app/auth-service
CMD ["/app/auth-service", "server"]
