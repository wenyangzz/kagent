package isvc

import (
	"context"
	"fmt"
	"log"
	"time"
)

type IsvcGrpcServer struct {
	UnimplementedIsvcServiceServer
}

const (
	RUNNING_TIMEOUT = 12
)

func (s *IsvcGrpcServer) CreateIsvc(ctx context.Context, in *CreateIsvcRequest) (*IsvcNull, error) {
	//log.Printf("Received: %v", in)

	model_name := in.GetModelSpecCreate().GetModelName()
	model_ver := in.GetModelSpecCreate().GetModelVersion()
	model_framework := in.GetModelSpecCreate().GetModelFramework()
	model_content := in.GetModelSpecCreate().GetModelContent()
	model_preparation := in.GetModelSpecCreate().GetModelPreparation()

	create_isvc, err := Create(model_name, model_ver, &model_framework, &model_content, &model_preparation)

	if err != nil {
		log.Printf("failed to create InferenceService: %v", err)
		return &IsvcNull{}, err
	}

	log.Printf("request to create InferenceService %v in namespace %v", create_isvc.Name, create_isvc.Namespace)

	for i := 0; i <= RUNNING_TIMEOUT; i++ {
		time.Sleep(10 * time.Second)
		check_isvc, err := Get(create_isvc.GetName())
		if err == nil {
			if check_isvc.Status.IsReady() {
				return &IsvcNull{}, nil
			}
		}
	}

	return &IsvcNull{}, fmt.Errorf("InferenceService %v created in namespace %v but not running", create_isvc.Name, create_isvc.Namespace)
}
