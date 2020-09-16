FROM golang:alpine AS build

ENV GOPROXY=https://proxy.golang.org

WORKDIR /go/src/userpool-connection
COPY . .
RUN GOOS=linux go build -o /go/bin/userpool-connection main.go

FROM alpine
COPY --from=build /go/bin/userpool-connection /go/bin/userpool-connection
ENTRYPOINT ["go/bin/userpool-connection"]