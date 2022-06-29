FROM golang:alpine 

RUN apk update && apk add git && rm /var/cache/apk/*

LABEL MAINTAINER="Andrew Alizaga, https://github.com/AndrewAlizaga"



RUN mkdir -p /auth
WORKDIR /auth

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy

COPY . .
RUN go build -o ./application ./main.go
COPY . .

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

EXPOSE 8080
ENV PORT=8080
ENV MONGO_CONNECTION_STRING=mongodb+srv://admin2:P6K2NQhFbohNYZrA@cluster0.ftad3.mongodb.net/myFirstDatabase?retryWrites=true&w=majority
RUN go build 

RUN chmod +x ./application


ENTRYPOINT [ "./application" ]

