FROM golang:1.5.1-onbuild
RUN go get -v github.com/raphael/goa/goagen
RUN go get -v github.com/raphael/goa/design/dsl
EXPOSE 8080
