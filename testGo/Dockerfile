FROM golang:1.18-alpine

WORKDIR ./testGo

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./*.go ./
COPY ./account/*.go ./account/

RUN go build -o /test-go-app

EXPOSE 8080

CMD [ "/test-go-app" ]
