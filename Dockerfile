FROM golang:1.18-bullseye as base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --no-create-home \
  --uid 654321 \
  go-user

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

FROM alpine:3.13

COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
COPY --from=base /app/main /app/config.json .

USER go-user:go-user

EXPOSE 9090

CMD ["./main"]


