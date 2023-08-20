FROM golang:1.20-alpine

# Adding Some Default requirenments
RUN apk add py3-pip gcc musl-dev python3-dev pango zlib-dev jpeg-dev openjpeg-dev g++ libffi-dev
RUN apk add --no-cache msttcorefonts-installer fontconfig
RUN update-ms-fonts

WORKDIR /app
ADD . .

RUN go mod download

RUN go build -o /EventHandler

EXPOSE 11111

CMD [ "/EventHandler" ]