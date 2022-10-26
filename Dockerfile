FROM golang:1.19

WORKDIR /app

COPY ./main /app/

CMD [ "./main" ]