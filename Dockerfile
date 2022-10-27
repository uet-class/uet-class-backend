FROM ubuntu:20.04

WORKDIR /app

COPY . /app/

RUN ls -la /app

CMD [ "./uet-class-backend" ]