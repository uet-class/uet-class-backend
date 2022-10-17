FROM golang:1.19

WORKDIR /app

COPY . /app/

RUN go build .

CMD [ "./main" ]