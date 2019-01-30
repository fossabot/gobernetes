#!/bin/bash

# create secret for alert manager
kubectl create -n monitoring secret generic alertmanager-main --from-file=monitoring/secrets/alertmanager.yaml

# create secret for accessing route53
kubectl create -n cert-manager secret generic route53 --from-file=secret_key

# create ConfigMap with modified traefik.toml
kubectl delete -n traefik configmap traefik-conf
kubectl create -n traefik configmap traefik-conf --from-file=ingress-controller-traefik/deployment_tls_ready/traefik.toml
