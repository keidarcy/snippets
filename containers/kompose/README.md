```sh
kompose convert -f ./docker-compose.yml
kubectl apply -f .
kubectl scale --replicas=5 deployment/web
kubectl get all
kubectl delete all --all
```