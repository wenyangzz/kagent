package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/asmcos/requests"
)

var (
	addr      = flag.String("addr", "10.58.96.67:80", "the address to connect to")
	hostname  = flag.String("hostname", "grpc-transformer.kagent.example.com", "the hostname for the ingress")
	modelname = flag.String("modelname", "grpc-transformer", "the name for the model")
	//jsonpath  = flag.String("jsonpath", "/Volumes/LiuWang-Work/kserve-test-yamls/input.json", "the path for the input json file")
	jsonpath = flag.String("jsonpath", "/Users/coco/go/src/github.com/kagent/test-client/tmp_json.csv", "the path for the input json file")
)

func main() {
	flag.Parse()

	jsonFile, err := os.Open(*jsonpath)
	if err != nil {
		log.Fatalf("open test model file error %v", err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("error reading json file %v", err)
	}

	req := requests.Requests()
	req.Header.Set("Host", *hostname)
	addr := fmt.Sprintf("http://%s/v1/models/%s:predict", *addr, *modelname)
	log.Printf("addr is: %v", addr)
	log.Printf("req is: %v", req.Header)
	log.Printf("req is: %v", req.Client)
	result, err := req.PostJson(addr, jsonData)
	if err != nil {
		log.Printf("could not greet: %v", err)
	}

	log.Printf("resp is: %v", result)
	log.Printf("resp is: %v", result.R)
	log.Printf("resp is: %v", result.R.StatusCode)
	log.Printf("resp is: %s", result.Text())

}
