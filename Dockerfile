FROM debian:latest
LABEL maintainer="gafarov@realnoevremya.ru"
RUN apt-get update && apt-get upgrade
RUN apt-get install -y ca-certificates
EXPOSE 8083
COPY . .
WORKDIR /build/linux
CMD [ "./text2image-service" ]

