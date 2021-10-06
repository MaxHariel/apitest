FROM golang:1.17.1-alpine

WORKDIR /app/

COPY ./ /app/
RUN go mod download
RUN go get github.com/cosmtrek/air
RUN ["chmod", "+x", "./.docker/entrypoint.sh"]

