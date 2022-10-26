FROM golang:1.19

WORKDIR /app

COPY ./uc-backend /app/

CMD [ "./uc-backend" ]