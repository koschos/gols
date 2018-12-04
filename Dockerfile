FROM golang:alpine as build

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/koschos/gols

COPY . $GOPATH/src/github.com/koschos/gols

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/gols

FROM scratch

COPY --from=build /go/bin/gols /main

EXPOSE 50000

ENTRYPOINT ["/main"]
