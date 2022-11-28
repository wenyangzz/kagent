#export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
#export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')

#kind
#export INGRESS_HOST=$(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')
#export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')

export INGRESS_HOST="127.0.0.1"
export INGRESS_PORT="80"

MODEL_NAME=grpc-lw
INPUT_PATH=./input.json
SERVICE_HOSTNAME=$(kubectl get inferenceservice ${MODEL_NAME} -n kagent -o jsonpath='{.status.url}' | cut -d "/" -f 3)
python3 grpc_client.py --host $INGRESS_HOST --port $INGRESS_PORT --model $MODEL_NAME --hostname $SERVICE_HOSTNAME --input_path $INPUT_PATH
