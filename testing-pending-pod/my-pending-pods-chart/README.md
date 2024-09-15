# README

This chart has been created to template multiple pods in a pending state. By default, it templates 3 pods (as defined in the values.yaml) in a single manifest.

To generate the manifests for the 3 pods in one file, you can run:
```
helm template my-pending-pods-chart > all-pods.yaml
```
Alternatively, to generate each pod in separate files, you can run:

```
cd testing-pending-pod
helm template my-pending-pods-chart \
  --set pods\[0\]=pending-pod-1, \
  --set container.name=busybox \
  --set container.image=busybox \
  --set nodeSelector.kubernetes\\.io\\/hostname=non-existent-node > pending-pod-1.yaml
```
You can repeat the helm template command for the other pods, adjusting the pod names and output file names as necessary.