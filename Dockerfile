FROM golang:1.17 as builder
RUN mkdir /api
ADD . /api   
WORKDIR /api

RUN CGO_ENABLED=0 GOOS=linux go build -o test cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /api .
CMD ["/test"]