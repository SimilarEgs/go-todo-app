# BUILD STAGE
FROM golang:1.18-alpine3.15 AS builder

WORKDIR /app
COPY . . 



RUN chmod +x wait-for.sh
RUN chmod +x start.sh
RUN go build -o main ./cmd/main.go


RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# RUN STAGE
FROM alpine

WORKDIR /app

RUN mkdir logger
# RUN touch log.log

COPY --from=builder /app/main . 
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY wait-for.sh .
COPY start.sh .
COPY ./migrations ./migrations


EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]

#docker run --name test-app --network app-network -p 8080:8080 web-app
#docker build -t web-app .


