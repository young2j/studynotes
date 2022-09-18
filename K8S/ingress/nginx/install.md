* 方式一
```shell
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```
* 方式二
```shell
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx
```

* 方式三
```shell
# https://github.com/kubernetes/ingress-nginx/blob/main/deploy/static/provider/cloud/deploy.yaml
kubectl apply -f ingress-nginx.yaml
```

* 方式四(首选)
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install ingress-nginx bitnami/nginx-ingress-controller
```