FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/golang_slackbot
COPY . .

ENV GIT_SSL_NO_VERIFY=1
RUN go get -u "github.com/YuyaBan/golang_SlackBot/domain"
RUN go get -u "github.com/nlopes/slack" 
RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/golang_slackbot /app
COPY app/config.toml /app

CMD /app/main