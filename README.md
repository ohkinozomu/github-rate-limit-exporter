# github-rate-limit-exporter

GitHub rate limit Prometheus exporter

## Install(Kubernetes)

```bash
$ helm repo add github-rate-limit-exporter  https://ohkinozomu.github.io/github-rate-limit-exporter/
$ helm install [release-name] github-rate-limit-exporter/github-rate-limit-exporter --set accessToken=YOUR_ACCESS_TOKEN
```