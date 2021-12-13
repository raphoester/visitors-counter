FROM golang:1.17-alpine
ENV DATA=/
RUN apk add git && git -C /go/src/ clone https://github.com/raphoester/visitors-counter 
WORKDIR /go/src/visitors-counter
RUN go mod tidy 
CMD go run . 
EXPOSE 80