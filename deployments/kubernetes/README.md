# k8s

```
kubectl delete pod mysql-0 --namespace address-go-namespace
kubectl delete pvc address-go-mysql-pv-claim --namespace address-go-namespace
kubectl delete pv address-go-mysql-pv-volume --namespace address-go-namespace
kubectl delete deployment address-go --namespace address-go-namespace
kubectl delete service address-go --namespace address-go-namespace
kubectl delete ns address-go-namespace
kubectl apply -f namespace.yaml
kubectl apply -f db.yaml
kubectl apply -f app.yaml
kubectl port-forward --namespace address-go-namespace service/mysql 3306:3306 --insecure-skip-tls-verify
kubectl port-forward --namespace address-go-namespace service/address-go 9999:9999 --insecure-skip-tls-verify
```