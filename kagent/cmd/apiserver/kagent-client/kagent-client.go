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

	"github.com/kagent/internal/crd/isvc"
)

var (
	addr      = flag.String("addr", "10.58.96.75:80", "the address to connect to")
	hostname  = flag.String("hostname", "kagent.example.com", "the hostname for the ingress")
	modelname = flag.String("modelname", "lwtest", "the name for the model")
	modelver  = flag.String("modelver", "0003", "the version for the model")
	modeltype = flag.String("modeltype", "tensorflow", "the version for the model")
	modelpath = flag.String("modelpath", "/Users/coco/go/src/github.com/kagent/test-client/model.tar.gz", "the path for the model file")
	prepath = flag.String("prepath", "/Users/coco/go/src/github.com/kagent/test-client/data_preparation.tar.gz", "the path for the preparation file")
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
	c := isvc.NewIsvcServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	modelFile, err := os.Open(*modelpath)
	if err != nil {
		log.Fatalf("open test model file error %v", err)
	}
	defer modelFile.Close()

	modelData, err := ioutil.ReadAll(modelFile)
	if err != nil {
		log.Fatalf("error reading json file %v", err)
	}

	preFile, err := os.Open(*prepath)
	if err != nil {
		log.Fatalf("open test model file error %v", err)
	}
	defer modelFile.Close()

	preData, err := ioutil.ReadAll(preFile)
	if err != nil {
		log.Fatalf("error reading json file %v", err)
	}

	r, err := c.CreateIsvc(ctx, &isvc.CreateIsvcRequest{
		ModelSpecCreate: &isvc.ModelSpec{ModelName: *modelname, ModelVersion: *modelver, ModelFramework: modelData, ModelContent: modelData, ModelPreparation: preData}})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("Greeting: %v", r)

}
