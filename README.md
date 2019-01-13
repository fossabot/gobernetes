# gobernetes
An ultra simple Go App ready to be deployed via Kubernetes.

### Build and run docker image:
Use `./build-local.sh` to run Docker container locally.<br>
Use `./build.sh` to rebuild Docker image & push it to the public [Docker hub](https://hub.docker.com/r/aracki/).

Docker contains only Ubuntu:16.04 and go binary executed.

### Kubernetes [components](https://github.com/Aracki/gobernetes/tree/master/k8s_manifests) included:

* Simple go app (go-ws)
* Traefik ingress Controller
* Nginx ingress Controller
* Prometheus + Grafana
* Jenkins
