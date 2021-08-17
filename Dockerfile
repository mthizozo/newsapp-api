FROM golang:1.16-alpine


ARG apikey
WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download

ENV API_KEY $apikey

COPY . .

RUN go build -o /newsapp

EXPOSE 8080

CMD [ "/newsapp" ]