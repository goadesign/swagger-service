FROM golang:onbuild
RUN go get -v github.com/goadesign/goa/goagen
RUN go get -v github.com/goadesign/goa/design/apidsl
RUN go get -v github.com/goadesign/goa/dslengine
RUN go get -v gopkg.in/yaml.v2
EXPOSE 8080
