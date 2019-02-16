FROM node:latest AS node_step

MAINTAINER NeonSky <jacob.a.eriksson@gmail.com>

RUN mkdir -p /go/src/studyit/frontend
WORKDIR /go/src/studyit/frontend
COPY . /go/src/studyit/frontend

RUN npm install
RUN npm run bundle

FROM golang:latest

WORKDIR /go/src/studyit/frontend
COPY --from=node_step /go/src/studyit/frontend /go/src/studyit/frontend

RUN go get -d -v
RUN go install
CMD go run main.go

EXPOSE 8080
