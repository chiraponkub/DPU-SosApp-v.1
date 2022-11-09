FROM golang:1.18-alpine

RUN apk update && apk upgrade && apk add --no-cache build-base bash git openssh
LABEL maintainer="Chirapon Hemtrakan <chirapon.job@gmail.com>"

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /dist/app .

EXPOSE 8000
CMD ["sh", "-c", "/dist/app"]
