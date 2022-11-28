#export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
#export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')

#kind
#export INGRESS_HOST=$(kubectl get po -l istio=ingressgateway -n istio-system -o jsonpath='{.items[0].status.hostIP}')
#export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')

export INGRESS_HOST="10.58.96.65"
export INGRESS_PORT="80"

MODEL_NAME=grpc-lw
INPUT_PATH=./input.json
#SERVICE_HOSTNAME=$(kubectl get inferenceservice ${MODEL_NAME} -o jsonpath='{.status.url}' | cut -d "/" -f 3)
#SERVICE_HOSTNAME=grpc-lw.kagent.example.com
SERVICE_HOSTNAME=grpc-transformer-predictor-default.kagent.example.com
python3 grpc_client.py --host $INGRESS_HOST --port $INGRESS_PORT --model $MODEL_NAME --hostname $SERVICE_HOSTNAME --input_path $INPUT_PATH

#MODEL_NAME=lwtest10
#SERVICE_HOSTNAME=grpc-transformer.kagent.example.com
#data='{"cols": ["Month", "Day", "Day_Of_Week", "Hour", "Current_Phase_Average", "Weather_Temperature_Celsius", "Weather_Relative_Humidity", "Global_Horizontal_Radiation", "Diffuse_Horizontal_Radiation", "Wind_Direction", "Weather_Daily_Rainfall", "Active_Power"], "Month": 4, "Day": 7, "Day_Of_Week": 2, "Hour": 18, "Current_Phase_Average": 0.30000004172325, "Weather_Temperature_Celsius": 2.693647476086928, "Weather_Relative_Humidity": -2.8762938864222645, "Global_Horizontal_Radiation":
#-0.5895636681176323, "Diffuse_Horizontal_Radiation": -0.6238329347353299, "Wind_Direction": 234.76290893555, "Weather_Daily_Rainfall": 0.0, "Active_Power": -0.5759014425002}'
#curl -v -H "Host: ${SERVICE_HOSTNAME}" -d "$data" http://${INGRESS_HOST}:${INGRESS_PORT}/v1/models/$MODEL_NAME:predict

#SERVICE_HOSTNAME=grpc-pt.kagent.example.com
#MODEL_NAME=grpc-pt
#file_path=./tmp_json2.csv
#data=`head -1 ${file_path}`
#echo $data
#curl -H "Host: ${SERVICE_HOSTNAME}" -H "Content-Type: application/json" -d "${data}" http://${INGRESS_HOST}:${INGRESS_PORT}/v1/models/${MODEL_NAME}:predict
#curl -H "Host: ${SERVICE_HOSTNAME}"  http://${INGRESS_HOST}:${INGRESS_PORT}/v2/health/live
