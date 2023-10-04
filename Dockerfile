# STEP 1 build executable binary
FROM golang:alpine as builder

LABEL maintainer "Darwin Smith <172265+dwin@users.noreply.github.com>"
LABEL app_version="0.1.0" architecture="amd64"

COPY . $GOPATH/src/github.com/dwin/hashify
WORKDIR $GOPATH/src/github.com/dwin/hashify

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/hashify cmd/hashify/main.go

# STEP 2 build a small image
# start from scratch
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/hashify /go/bin/hashify
ENV LISTEN_HTTP :1313
EXPOSE 1313
ENTRYPOINT ["/go/bin/hashify"]

# docker build . -t dwin/hashify:0.1.0
# docker push dwin/go-hashify:0.1.0
# docker run -d -p 1313:1313 --name hashify dwin/hashify:0.1.0

# docker run -d --name api dwin/hashify
