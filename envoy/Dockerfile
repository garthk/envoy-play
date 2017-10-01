FROM lyft/envoy
ADD envoy.json /etc/
CMD ["/usr/local/bin/envoy", "-c", "/etc/envoy.json", "-l", "info", "--service-cluster", "monoenvoy", "--service-node", "A"]
EXPOSE 8080 9090
