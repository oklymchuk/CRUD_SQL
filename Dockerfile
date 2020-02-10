FROM golang:latest
LABEL maintainer="Ostap Klymchuk <oklymchuk@gmail.com>"
RUN mkdir -p /app 

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./crud_sql .

EXPOSE 8080
CMD ["/crud_sql"]
