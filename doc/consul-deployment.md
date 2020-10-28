### Consul deployment

#### docker run
```bash
docker run -p 8500:8500 -p 8502:8502 -p 8600:8600 -d --name=dev-consul -e CONSUL_BIND_INTERFACE=eth0 consul:1.8.4
```