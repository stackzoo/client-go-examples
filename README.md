# CLIENT-GO EXAMPLES

This repo contains go scripts that illustrate how to use <a href="https://github.com/kubernetes/client-go">client-go</a> and other
<br/>
k8s packages in order to call k8s api server programmatically.

## Instructions
As an example you can first launch the deployment creation:
```sh
❯ go run deployments/create_deployment.go

Deployment created successfully!
```
<br/>

and then retrieve the pods metrics:
```sh
❯ go run metrics/pods_metrics.go

SHOWING PODS METRICS
pod name: busybox-58d9b8dbdc-b4x27
namespace: test
cpu usage (millicore): 416m
memory usage: 0Mi

pod name: busybox-58d9b8dbdc-glgjt
namespace: test
cpu usage (millicore): 436m
memory usage: 0Mi

pod name: busybox-58d9b8dbdc-wwbrp
namespace: test
cpu usage (millicore): 428m
memory usage: 0Mi
```
