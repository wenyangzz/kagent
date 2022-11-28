/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/kagent/internal/crd/isvc"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {

	namespace := "default"
	inf_name := "flower-lwtest"
	model_type := "tensorflow"
	var model_content []byte

	inf, err := inferences.Create(namespace, inf_name, model_type, model_content)

	if errors.IsNotFound(err) {
		fmt.Printf("InferenceServices %s in namespace %s not found\n", inf_name, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("failed to create InferenceServices %s in namespace %s: %v\n",
			inf_name, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found InferenceServices %s in namespace %s\n", inf_name, namespace)
		fmt.Println("Value is", inf)
	}

	time.Sleep(10 * time.Second)

}
