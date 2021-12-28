FROM golang:1.17-alpine

COPY . .

ENV GOPATH=

RUN go build -o /app cmd/api.go


FROM alpine:3.14

COPY --from=0 /app /app

ENTRYPOINT [ "/app" ]
