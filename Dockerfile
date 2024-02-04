FROM golang-1.16-alpine:stable AS builder

WORKDIR /app

COPY go.mod go.sum Makefile ./
RUN make mod

COPY . .
RUN make build

FROM alpine:3.14

COPY --from=builder /app/bin/alcl-go-function-template /alcl-go-function-template

ENTRYPOINT [ "/alcl-go-function-template" ]
