FROM golang:latest AS build

#enables go.mod
ENV GO111MODULE=on
ENV PORT=8001

EXPOSE 8001

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
ENV DIR=$GOPATH/src/app

RUN mkdir -p $DIR
WORKDIR $DIR/

COPY go.mod go.sum $DIR/
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /app cmd/attend-io-server/main.go

#multi-stage
#TODO: do we need full Alpine Linux?
FROM alpine:latest
RUN apk --update add ca-certificates
COPY --from=build /app /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/app"]
