FROM lyft/envoy
ADD envoy.json /etc/
CMD ["/usr/local/bin/envoy", "-c", "/etc/envoy.json", "-l", "info"]
EXPOSE 8080 9090
