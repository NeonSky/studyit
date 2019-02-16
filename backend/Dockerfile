FROM golang:latest
MAINTAINER NeonSky <jacob.a.eriksson@gmail.com>

RUN mkdir -p /go/src/github.com/NeonSky/studyit/backend
WORKDIR /go/src/github.com/NeonSky/studyit/backend

COPY . /go/src/github.com/NeonSky/studyit/backend

RUN go get -d -v
RUN go install
CMD go run main.go

EXPOSE 9000
