# Envoy Play

[Envoy] looks plenty useful, but its
[reference configurations][refcfg] are designed for a scale I don't have to
deal with yet. A single layer of proxies should be able to proxy both customer
requests and internal requests.

## Usage

* `docker-compose up` in one terminal
* `curl -D - http://0.0.0.0:8080/first` in another

You should see `curl` output resembling:

```
HTTP/1.1 200 OK
date: Sun, 01 Oct 2017 05:57:44 GMT
content-length: 13
content-type: text/plain; charset=utf-8
x-envoy-upstream-service-time: 12
server: envoy

first
second
```

## Composition

We're running three containers:

* The `first` service, which needs to talk to the `second` before replying
* The `second` service, and
* The `envoy` service

To demonstrate both customer-service and service-service proxying, we've set up `first` to talk to `second` via `envoy`.

## Resources:

* [Envoy Slack][slack]
* [Envoy Configuration JSON Schema][schemata]

[Envoy]: https://envoyproxy.github.io
[schemata]: https://github.com/envoyproxy/envoy/blob/master/source/common/json/config_schemas.cc
[refcfg]: https://envoyproxy.github.io/envoy/install/ref_configs.html#install-ref-configs
[slack]: http://envoyslack.cncf.io
