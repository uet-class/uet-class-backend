FROM golang:1.19 AS build
WORKDIR /app
COPY . /app
RUN go build /app/main.go

FROM ubuntu:20.04 AS deploy
WORKDIR /app
COPY --from=build /app/main /app/
CMD [ "/app/main" ]