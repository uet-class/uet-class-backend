FROM golang:1.19 AS build
WORKDIR /app
COPY . /app
RUN go build /app/main.go

FROM thainm/ubuntu-base-image:latest AS deploy
WORKDIR /app
COPY --from=build /app/main /app/
CMD [ "/app/main" ]