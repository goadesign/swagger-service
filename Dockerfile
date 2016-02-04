FROM golang:onbuild
RUN go get -v github.com/goadesign/goa/goagen
RUN go get -v github.com/goadesign/goa/design/apidsl
RUN go get -v github.com/goadesign/goa/dslengine
RUN go get golang.org/x/tools/cmd/goimports
EXPOSE 8080
