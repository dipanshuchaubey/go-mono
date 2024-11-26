FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o bin

FROM alpine:3.14

COPY --from=build /app/bin/* /bin/

EXPOSE 50051

CMD ["/bin/user-service"] 
