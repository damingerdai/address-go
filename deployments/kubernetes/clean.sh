kubectl delete pod mysql-0 --namespace address-go-namespace
kubectl delete pvc address-go-mysql-pv-claim --namespace address-go-namespace
kubectl delete pv address-go-mysql-pv-volume --namespace address-go-namespace
kubectl delete deployment address-go --namespace address-go-namespace
kubectl delete service address-go --namespace address-go-namespace
kubectl delete ns address-go-namespace