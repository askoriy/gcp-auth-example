FROM golang:1.13

ENV APP_NAME api
# ENV PORT 8080
ENV API_URL 'https://api-gw-example-backend-afcylnj57a-ew.a.run.app/api'
# ENV API_URL 'https://jsonplaceholder.typicode.com/todos/1'

COPY /cloud-run/api-gw/api/go/src/api /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}