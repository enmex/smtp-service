# SMTP Service
Implementing a mail client in Go for your services
### How to use

```bash
$ docker pull ghcr.io/ignavan39/go-smtp:latest
```

### create configuration
[example configuration](https://github.com/enmex/smtp/blob/master/config.example.yml)
create a file with any name and extension .yml 


### How to start

Add to your docker-compose
```yml
  smtp-service: 
    container_name: smtp-service
    image: "ghcr.io/ignavan39/go-smtp:latest"
    ports:
      - ${PORT:-8081}:80
    environment:
      - CONFIG_PATH=./config.example.yml
```
use CONFIG_PATH from ***create configuration*** step