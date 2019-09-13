FROM golang:1.13-buster

WORKDIR /go/src/app

COPY . .

RUN go install && go build

RUN ls -la

CMD ["k8s-goagent"]