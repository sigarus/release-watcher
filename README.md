<p align="center">
  <a href="https://goreportcard.com/report/github.com/zvlb/release-watcher">
    <img src="https://goreportcard.com/report/github.com/zvlb/release-watcher" alt="Go Report Card">
  </a>
</p>

# Release Watcher

## Description

Release Watcher is a convenient tool that helps you stay informed about the latest updates and releases of your favorite software and applications. Simply add the products you're interested in to the list, and the program will automatically notify you when new versions become available. This way, you'll always be aware of the latest changes and new features ready for installation.

## Features

- [x] Collect informations from GitHub Releases
- [x] Send information about Release to Telegram
- [x] Send information about Release to Slack

## How to start

### Localy

Run `go run main.go --config-file <PATH_TO_FILE>`, where *<PATH_TO_FILE>* - path to file with program parametrs, like:

```yaml
releases:
  github:
    - "zvlb/release-watcher"
    - "kaasops/envoy-xds-controller"
    - "kaasops/vector-operator"
    - "argoproj/argo-cd"
    - "argoproj/argo-workflows"
    - "prometheus/alertmanager"
    - "ansible/awx"
    - "projectcapsule/capsule"
    - "cert-manager/cert-manager"
    - "hashicorp/consul"
    - "ceph/ceph-csi"
    - "kubernetes/dashboard"
    - "kubernetes/ingress-nginx"
    - "cilium/cilium"
    - "cilium/hubble"
    - "prymitive/karma"
    - "kedacore/keda"
    - "keycloak/keycloak"
    - "kyverno/kyverno"
    - "VictoriaMetrics/VictoriaMetrics"
    - "vmware-tanzu/velero"


recievers:
  telegram:
    - chatID: "<CHAT_ID>"
      token: <BOT_TOKEN>
    - chatID: "<CHAT_ID>"
      token: <BOT_TOKEN>
  slack:
    - channelName: "<CHANNEL_NAME>"
      hook: "<WEBHOOK>"
```

### Kubernetes

For install Release Watcher to Kubernetes prefer use Helm.

First you need to prepare values.yaml-file:

```yaml
releases:
  github:
    - "kaasops/envoy-xds-controller"
    - "kaasops/vector-operator"
    - "argoproj/argo-cd"
    - "argoproj/argo-workflows"
    - "prometheus/alertmanager"
    - "ansible/awx"
    - "projectcapsule/capsule"
    - "cert-manager/cert-manager"
    - "hashicorp/consul"
    - "ceph/ceph-csi"
    - "kubernetes/dashboard"
    - "kubernetes/ingress-nginx"
    - "cilium/cilium"
    - "cilium/hubble"
    - "prymitive/karma"
    - "kedacore/keda"
    - "keycloak/keycloak"
    - "kyverno/kyverno"
    - "VictoriaMetrics/VictoriaMetrics"
    - "vmware-tanzu/velero"
    - "zvlb/release-watcher"


recievers:
  telegram:
    - chatID: "<CHAT_ID>"
      token: <BOT_TOKEN>
  slack:
    - channelName: "<CHANNEL_NAME>"
    hook: "<WEBHOOK>"
```

Add a chart helm repository with follow commands:

```bash
helm repo add release-watcher https://zvlb.github.io/release-watcher/helm
helm repo update
```

List all charts and versions of vm repository available to installation:

```bash
helm search repo release-watcher
```

The command must display existing helm chart e.g.

```bash
NAME                            CHART VERSION   APP VERSION     DESCRIPTION                            
release-watcher/release-watcher 0.0.8           v0.0.8          A Helm chart to install Release Watcher
```

Install chart with command:

```bash
helm install release-watcher release-watcher/release-watcher -f values.yaml -n NAMESPACE
```