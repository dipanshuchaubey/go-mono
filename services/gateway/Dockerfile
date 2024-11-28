FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN go build -o bin/gateway cmd/gateway/main.go

FROM alpine:3.14

COPY --from=build /app/bin/gateway /bin/gateway

COPY services/gateway/config/ /mnt/secrets/
COPY services/gateway/routes/routes.yaml /mnt/secret/routes/routes.yaml

EXPOSE 50051

CMD ["/bin/gateway"]
