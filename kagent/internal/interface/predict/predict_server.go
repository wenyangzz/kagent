package predict

import (
	"context"
	"fmt"
	"log"

	"github.com/asmcos/requests"
	"github.com/kagent/internal/crd/isvc"
	"github.com/kagent/internal/utils"
)

type PredictGrpcServer struct {
	UnimplementedKPredictionServiceServer
}

const (
	RUNNING_TIMEOUT = 12
	//INGRESS_ADDR    = "istio-ingressgateway.istio-system.svc.cluster.local:80"
	INGRESS_ADDR = ".kagent.svc.cluster.local:80"
)

func (s *PredictGrpcServer) KPredict(ctx context.Context, in *KPredictRequest) (output *KPredictResponse, err error) {
	log.Printf("received: %v", in)

	isvc_name := isvc.GenIsvcName(in.ModelName, in.ModelVersion)

	get_isvc, err := isvc.Get(isvc_name)
	if err != nil {
		log.Printf("failed to get isvc: %v", err)
		return nil, err
	}

	log.Printf("send predict request to host %v", get_isvc.Status.URL.Host)
	log.Printf("the transformer is %v", get_isvc.Spec.Transformer)
	var svc_addr string
	if get_isvc.Spec.Transformer == nil {
		svc_addr = utils.BuildStringfunc(isvc_name, "-predictor-default", INGRESS_ADDR)
	} else {
		svc_addr = utils.BuildStringfunc(isvc_name, "-transformer-default", INGRESS_ADDR)
	}

	log.Printf("svc_addr is %s", svc_addr)

	req := requests.Requests()
	//req.Header.Set("Host", get_isvc.Status.URL.Host)
	addr := fmt.Sprintf("http://%s/v1/models/%s:predict", svc_addr, isvc_name)
	result, err := req.PostJson(addr, string(in.Inputs))
	if err != nil {
		log.Printf("could not greet: %v", err)
		return nil, err
	}

	log.Printf("resp is: %v", result.R.StatusCode)

	output = &KPredictResponse{Outputs: []byte(result.Text())}

	log.Printf("output is: %v", output)

	return output, nil
}
