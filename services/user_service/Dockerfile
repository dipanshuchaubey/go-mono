FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN go build -o bin/user-service cmd/user_service/main.go

FROM alpine:3.14

COPY --from=build /app/bin/user-service /bin/user-service

COPY services/user_service/config/ /mnt/secrets/

EXPOSE 50051

CMD ["/bin/user-service"] 
