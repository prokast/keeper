Keeper is a tool for monitor certificate expiry date. Ease for configure keeper need dns names or ip's for start tracking certificates expiry date. Deploy keeper into kubernetes or start docker container on your machine.
# Local storage
Keeper can show info from her own filesystem. Mount into container volume with your .pem cert files and refresh keeper dashboard. Information about .pem files now showing on dashboard (added in 1.1.0).
# Configuration
For start monitor your certs you need add [config.yaml](configs/cfg/config.yaml) file into ./configs/cfg/ directory in keeper container. Config example:
```yaml
# Example config file
global:
  logLevel: "info"
  timeout: "15s"

# Static certs. Simply note
note:
- alias: "root"
  expiredDate: "2022-11-07"
  description: 'example'
- alias: "server"
  expiredDate: "2022-10-29"
  description: 'example'
- alias: "client"
  expiredDate: "2022-10-29"
  description: 'example'

# Certs from remote host
url:
- host: "github.com"
  port: "443"
  description: "example"
- host: "google.com"
  port: "443"
  description: "example"
- host: "example.com"
  port: "443"
  description: "example"
```
# Kubernetes
Deploy keeper with [helm-charts](https://github.com/kooberetis/helm-charts/tree/main/keeper)
# Docker
Run with docker on your machine. Mount config/cfg/ directory with [config.yaml](configs/cfg/config.yaml) into container
```dockerfile
FROM kooberetis/keeper:latest
VOLUME [ "./configs/cfg" ]
EXPOSE 8080
CMD ["/bin/sh", "-c", "./keeper"]
```
# License
Apache License 2.0, see [LICENSE](LICENSE).
