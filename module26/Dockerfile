FROM golang
LABEL version="1.0.0"
LABEL maintainer="Ivan Ivanov<test@test.ru>"
RUN mkdir -p /go/src/website
WORKDIR /go/src/website
ADD main.go .
ADD go.mod .
RUN go install .
ENTRYPOINT /go/bin/website
EXPOSE 8080