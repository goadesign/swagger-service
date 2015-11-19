FROM golang:onbuild
RUN go get github.com/raphael/goa/goagen
EXPOSE 8080
