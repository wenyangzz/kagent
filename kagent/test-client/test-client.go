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
	"google.golang.org/grpc/metadata"

	inf "github.com/kagent/internal/crd/isvc"
)

var (
	addr = flag.String("addr", "10.58.96.65:80", "the address to connect to")
	//hostname = flag.String("hostname", "httpbin.example.com", "the hostname for the ingress")
	hostname = flag.String("hostname", "kagent.example.com", "the hostname for the ingress")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(*hostname))

	//creds_tls, err := credentials.NewClientTLSFromFile("/Volumes/LiuWang-Work/istio-test-certs/new-istio-dns-certs/example.com.crt", "")
	//if err != nil {
	//	log.Fatalf("did not newTLS: %v", err)
	//}

	//creds_tls := credentials.NewTLS(&tls.Config{
	//	InsecureSkipVerify: true,
	//})

	//opts = append(opts, grpc.WithTransportCredentials(creds_tls))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	//t := oauth2.Token{
	//	AccessToken: "invalid",
	//	TokenType:   "Bearer",
	//}
	//creds := oauth.NewOauthAccess(&t)
	//log.Printf("creds: %v", creds)

	//opts = append(opts, grpc.WithPerRPCCredentials(creds))

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := inf.NewInferencesClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	jwt_token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaXNzIjoidGVzdGluZ0BzZWN1cmUuaXN0aW8uaW8iLCJpYXQiOjE1MTYyMzkwMjJ9.Zxf-YzJxvbJ70IENLofhU9uTuGpq_BNZhgnndhljOv929geF2x6dYO8vrlKyJd0iNAh2jrF-D98GrTt4SKX0LLy2PQVAMNMwbUswOD26_k-AQuUrcV5pMbydb543akHelMQk2pPuy2mL9dKWui38AEb42WGRpjvCCm06GhrSsCbIXjbwSwOon0YjEtQS5aGTCyXXH6fQNwNJ7mTHOKVaVfovYiwtVmg5EEUy6JJKpTNUR2uVwtGK2Llbn5F4thO0yWhOwfNLXBbBuCH9r6W0qAtL_XkfPq67uIvI5xbxHByDGt4EHME4lCJSwvmIJO16fVF-vdjdpYSJ1ZLQi10rww"

	headers_input := metadata.New(map[string]string{"authorization": "Bearer " + jwt_token})

	log.Printf("headers_input: %v", headers_input)
	log.Printf("Authorization: %v", headers_input["authorization"])

	ctx = metadata.NewOutgoingContext(ctx, headers_input)

	var headers metadata.MD

	log.Printf("headers: %v", headers)

	modelFile, err := os.Open("/Users/coco/go/src/github.com/kagent/test-client/tf_model_1.tar.gz")
	if err != nil {
		log.Fatalf("open test model file error %v", err)
	}
	defer modelFile.Close()

	modelData, err := ioutil.ReadAll(modelFile)
	if err != nil {
		log.Fatalf("error reading json file %v", err)
	}

	r, err := c.CreateInferences(ctx, &inf.CreateInferencesRequest{
		Infs: &inf.InferencesMessage{ModelName: "grpc-lw", ModelVersion: "v1", ModelType: "tensorflow", ModelContent: modelData}}, grpc.Header(&headers))
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Printf("headers: %v", headers)
	log.Printf("headers_input: %v", headers_input)
	log.Printf("Greeting: %v", r)

}
