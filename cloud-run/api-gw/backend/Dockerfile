FROM golang:1.13

ENV APP_NAME backend
# ENV PORT 8080
ENV API_URL 'https://jsonplaceholder.typicode.com/todos/1'

COPY /cloud-run/api-gw/backend/go/src/backend /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

RUN pwd
RUN ls -l

RUN go get ./
RUN go build -o ${APP_NAME}

CMD ./${APP_NAME}

EXPOSE ${PORT}
