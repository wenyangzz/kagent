

kubectl apply -f /Users/coco/go/src/github.com/kagent/deployments/ingressclass.yaml

kubectl apply -f /Volumes/LiuWang-Work/pv-and-pvc.yaml
kubectl apply -f /Volumes/LiuWang-Work/model-store.yaml

sleep 30

kubectl cp /Volumes/LiuWang-Work/test-model/tensorflow-0001 model-store-pod:/pv/tensorflow-0001 -n kagent -c model-store

kubectl apply -f /Users/coco/go/src/github.com/kagent/deployments/kagent.yaml