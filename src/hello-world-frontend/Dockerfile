FROM golang:1.9-alpine as builder
WORKDIR /go/src/github.com/altmetric/cakefindr-hello-world-frontend/
RUN apk add --update git build-base
COPY . ./
RUN go build -o server ./main.go

FROM alpine:3.7
WORKDIR /app/
COPY --from=builder /go/src/github.com/altmetric/cakefindr-hello-world-frontend/server .
CMD ["./server"]
