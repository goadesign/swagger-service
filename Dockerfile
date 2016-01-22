FROM golang:onbuild
RUN go get -v github.com/goadesign/goa/goagen
RUN go get -v github.com/goadesign/goa/design/dsl
EXPOSE 8080
