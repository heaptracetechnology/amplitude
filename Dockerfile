FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/msingleton/amplitude-go

WORKDIR /go/src/github.com/heaptracetechnology/amplitude

ADD . /go/src/github.com/heaptracetechnology/amplitude

RUN go install github.com/heaptracetechnology/amplitude

ENTRYPOINT amplitude

EXPOSE 3000