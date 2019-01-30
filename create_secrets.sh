#!/bin/bash

kubectl create secret generic alertmanager-main -n monitoring --from-file=k8s_manifests/monitoring/secrets/alertmanager.yaml
