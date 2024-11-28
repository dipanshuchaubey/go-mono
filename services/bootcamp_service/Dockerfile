FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

ENV GO111MODULE=on

RUN go build -o bin/bootcamp-service cmd/bootcamp_service/main.go

FROM alpine:3.14

COPY --from=build /app/bin/bootcamp-service /bin/bootcamp-service

COPY services/bootcamp_service/config/ /mnt/secrets/

EXPOSE 50051

CMD ["/bin/bootcamp-service"] 
