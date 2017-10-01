FROM golang:alpine
# RUN x && y && z optimises for image size at the expense of build caching
RUN apk add --no-cache git 
RUN go get github.com/labstack/echo/...
ADD . /go/src/app/
RUN cd /go/src/app && go-wrapper download && go-wrapper install
# ↑ automatically builds, but produces 280MB images

# ↓ fixes that by re-starting the build
FROM alpine
ENV SVC_ADDR=0.0.0.0 SVC_PORT=8083
COPY --from=0 /go/bin/app /usr/local/bin/discovery
CMD ["/usr/local/bin/discovery"]
EXPOSE 8083
