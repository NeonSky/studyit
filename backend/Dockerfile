FROM golang:latest
MAINTAINER NeonSky <jacob.a.eriksson@gmail.com>

RUN mkdir -p /go/src/studyit/backend
WORKDIR /go/src/studyit/backend

COPY . /go/src/studyit/backend

RUN go get -d -v
RUN go install
CMD go run main.go

EXPOSE 9000
