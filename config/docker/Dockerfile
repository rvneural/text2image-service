FROM golang:latest

LABEL authors="Artem Gafarov"
LABEL email="gafarov@realnoevremya.ru"
LABEL version="2.0"

WORKDIR /src/text2image-service
ADD . .

RUN go mod download && go mod verify

EXPOSE 80

WORKDIR /src/text2image-service/cmd/text2image-service
RUN go build

CMD ["./text2image-service"]