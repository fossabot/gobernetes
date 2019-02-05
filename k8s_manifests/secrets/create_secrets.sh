#!/bin/bash

# ALERT MANAGER SLACK RECEIVER recreation
kubectl delete -n monitoring secret alertmanager-main
kubectl create -n monitoring secret generic alertmanager-main --from-file=alertmanager.yaml

# ROUTE53 ACCESS
kubectl create -n cert-manager secret generic route53 --from-file=secret_key

# GRAFANA AUTHENTICATION
kubectl create -n monitoring secret generic grafana-auth --from-file=grafana-auth

# TRAEFIK.TOML recreation
kubectl delete -n traefik configmap traefik-conf
kubectl create -n traefik configmap traefik-conf --from-file=../ingress-controller-traefik/deployment_tls_ready/traefik.toml

