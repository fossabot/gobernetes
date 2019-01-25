# gobernetes
An ultra simple Go App ready to be deployed via Kubernetes.

### Build and run docker image:
Use `./build-local.sh` to run Docker container locally.<br>
Use `./build.sh` to rebuild Docker image & push it to the public [Docker hub](https://hub.docker.com/r/aracki/).

### Kubernetes [components](https://github.com/Aracki/gobernetes/tree/master/k8s_manifests) include:

* Simple go app (go-ws)
* Traefik Ingress Controller
* Nginx Ingress Controller
* Prometheus
* Alert Manager
* Grafana
* Jenkins
* cert-manager with issuers