FROM docker.io/bayarindev/gogin:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go get github.com/99designs/gqlgen
RUN go get github.com/rubenv/sql-migrate/...

RUN go build -o main .

CMD ["/app/main"]
