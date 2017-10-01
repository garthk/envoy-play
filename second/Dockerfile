FROM golang:alpine
ADD . /go/src/app/
RUN cd /go/src/app && go-wrapper download && go-wrapper install
# ↑ automatically builds, but produces 280MB images

# ↓ fixes that by re-starting the build
FROM alpine
ENV SVC_ADDR=0.0.0.0 SVC_PORT=8082
COPY --from=0 /go/bin/app /usr/local/bin/second
CMD ["/usr/local/bin/second"]
EXPOSE 8082
