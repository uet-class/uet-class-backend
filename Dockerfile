FROM golang:1.19

WORKDIR /app

COPY ./uet-class-backend /app/

CMD [ "./uet-class-backend" ]