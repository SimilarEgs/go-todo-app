# BUILD STAGE
FROM golang:1.18-alpine AS buster

WORKDIR /app
COPY . . 
RUN go build -o main ./cmd/main.go


# RUN STAGE
FROM alpine

WORKDIR /app

RUN mkdir logger
RUN touch log.log

COPY --from=buster /app/main . 
COPY app.env .

EXPOSE 8080

CMD [ "/app/main" ]

#docker run --name test-app --network app-network -p 8080:8080 web-app
#docker build -t web-app .


