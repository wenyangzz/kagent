package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/kagent/internal/interface/predict"
)

var (
	addr      = flag.String("addr", "10.58.96.75:80", "the address to connect to")
	hostname  = flag.String("hostname", "kagent.example.com", "the hostname for the ingress")
	modelname = flag.String("modelname", "lwtest", "the name for the model")
	modelver  = flag.String("modelver", "0003", "the version for the model")
	modeltype = flag.String("modeltype", "xgboost", "the version for the model")
	jsonpath  = flag.String("jsonpath", "/Users/coco/go/src/github.com/kagent/test-client/xgb_trans_one.csv", "the path for the input json file")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(*hostname))

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	p := predict.NewKPredictionServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	jsonFile, err := os.Open(*jsonpath)
	if err != nil {
		log.Fatalf("open test model file error %v", err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("error reading json file %v", err)
	}

	r, err := p.KPredict(ctx, &predict.KPredictRequest{ModelName: *modelname, ModelVersion: *modelver, ModelType: *modeltype, Inputs: jsonData})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Outputs)

}
