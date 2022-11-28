package isvc

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"github.com/kagent/internal/model/kmodel"
	"github.com/kagent/internal/utils"
	ks "github.com/kserve/kserve/pkg/apis/serving/v1beta1"
	ksclient "github.com/kserve/kserve/pkg/client/clientset/versioned/typed/serving/v1beta1"
	"github.com/kserve/kserve/pkg/constants"
)

const (
	KSERVE_KIND      = "InferenceService"
	KSERVE_APIV      = "serving.kserve.io/v1beta1"
	KSERVE_NAMESPACE = "kagent"
	TRANS_PVC_NAME   = "task-pv-claim"
	TRANS_VOL_NAME   = "model-store"
	TRANS_UPKG_NAME  = "prediction"
	TRANS_UDAT_NAME  = "data_stats"
)

var kubeconfig string

var clientset *ksclient.ServingV1beta1Client

func init() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("failed to get kubeconfig: %v", err)
	} else {
		log.Printf("kubeconfig is %v", config)

		ks.AddToScheme(scheme.Scheme)

		// create the clientset
		clientset, err = ksclient.NewForConfig(config)
		if err != nil {
			log.Printf("failed to create kubeconfig: %v", err)
		}
	}
}

func GetClientIn() (result *ksclient.ServingV1beta1Client, err error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("failed to get kubeconfig: %v", err)
	}

	println("kubeconfig is", config)

	ks.AddToScheme(scheme.Scheme)

	// create the clientset
	result, err = ksclient.NewForConfig(config)
	return
}

func GetClientOut() (result *ksclient.ServingV1beta1Client, err error) {
	var config_filepath string
	if home := homedir.HomeDir(); home != "" {
		config_filepath = filepath.Join(home, ".kube", "config")
	}

	if flag.Lookup("kubeconfig") == nil {
		flag.StringVar(&kubeconfig, "kubeconfig", config_filepath, "absolute path to the kubeconfig file")
	} else {
		flag.Set("kubeconfig", config_filepath)
		kubeconfig = flag.Lookup("kubeconfig").Value.String()
	}

	flag.Parse()
	println("kubeconfig is", kubeconfig)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("failed to get kubeconfig: %v", err)
	}

	ks.AddToScheme(scheme.Scheme)

	// create the clientset
	result, err = ksclient.NewForConfig(config)
	return
}

func Create(n string, v string, f *[]byte, m *[]byte, p *[]byte) (result *ks.InferenceService, err error) {

	infs, err := clientset.InferenceServices(KSERVE_NAMESPACE).List(metav1.ListOptions{})
	if err != nil {
		log.Printf("failed to get inferences list: %v", err)
	}
	log.Printf("There are %d InferenceServices in the cluster\n", len(infs.Items))

    t  := "xgboost"
	tv := "2.6.2"

	km, err := kmodel.GetKModel(n, v)
	if err != nil {
		log.Printf("model name %v version %v not find\n", n, v)
		km, err = kmodel.SaveKModel(n, v, t, tv, m ,p)
		if err != nil {
			log.Printf("save model name %v version %v error when create isvc\n", n, v)
			return nil, err
		}
	}

	input_uri := utils.BuildStringfunc("pvc://task-pv-claim/", km.ModelContentPath, "/model")	
	isvcName := GenIsvcName(n, v)
    imageName := "lw233/xgboost"
	trans_base_image_name := "lw233/xgb-transformer-base"
	trans_project_name := "xgb_transformer"

	var input_predictor ks.PredictorSpec
    var transformer_base_image_name string
	var transformer_project_name string

	switch t {
	case constants.SupportedModelTensorflow:
		input_predictor.Tensorflow = &ks.TFServingSpec{
			PredictorExtensionSpec: ks.PredictorExtensionSpec{
				StorageURI: &input_uri,
				Container: v1.Container{Image: imageName},
			},
		}
		transformer_base_image_name = trans_base_image_name
		transformer_project_name = trans_project_name
	case constants.SupportedModelXGBoost:
		input_predictor.XGBoost = &ks.XGBoostSpec{
			PredictorExtensionSpec: ks.PredictorExtensionSpec{
				StorageURI: &input_uri,
				Container: v1.Container{Image: imageName},
			},
		}
		transformer_base_image_name = trans_base_image_name
		transformer_project_name = trans_project_name
	case constants.SupportedModelSKLearn:
		input_predictor.SKLearn = &ks.SKLearnSpec{
			PredictorExtensionSpec: ks.PredictorExtensionSpec{
				StorageURI: &input_uri,
				Container: v1.Container{Image: imageName},
			},
		}
		transformer_base_image_name = trans_base_image_name
		transformer_project_name = trans_project_name
	case constants.SupportedModelLightGBM:
		input_predictor.LightGBM = &ks.LightGBMSpec{
			PredictorExtensionSpec: ks.PredictorExtensionSpec{
				StorageURI: &input_uri,
				Container: v1.Container{Image: imageName},
			},
		}
		transformer_base_image_name = trans_base_image_name
		transformer_project_name = trans_project_name
	default:
		log.Printf("unsupported model type %s\n", t)
		return nil, fmt.Errorf("unsupported model type %s", t)
	}

	var isvcSpec ks.InferenceServiceSpec 
	 
	if km.ModelHasPreparation {
		volume_s := []v1.Volume{
			v1.Volume{
				Name: TRANS_VOL_NAME,
				VolumeSource: v1.VolumeSource{
					PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
						ClaimName: TRANS_PVC_NAME,
					},
				},
			},
		}

		volumeMounts_s := []v1.VolumeMount{
			v1.VolumeMount{
				MountPath: utils.BuildStringfunc("/", TRANS_UDAT_NAME),
				Name: TRANS_VOL_NAME,
				SubPath: utils.BuildStringfunc(km.ModelContentPath, "/", TRANS_UDAT_NAME),
			},
			v1.VolumeMount{
				MountPath: utils.BuildStringfunc("/", TRANS_UPKG_NAME),
				Name: TRANS_VOL_NAME,
				SubPath: utils.BuildStringfunc(km.ModelPreparationPath, "/", TRANS_UPKG_NAME),
			},
		}

		command_s := []string{
			"python",
			"-m",
			transformer_project_name,
		}

		args_s := []string{
			"--model_name",
			isvcName,
			"--signature_name",
			"serving_default",
			"--protocol",
			"http",
			"--user_package",
			TRANS_UPKG_NAME,
			"--user_file",
			"data_preparation_prediction.py",
			"--data_path",
			TRANS_UDAT_NAME,
		}

		containers_s := []v1.Container{
			v1.Container{
				Name: "kserve-container",
				Image: transformer_base_image_name,
				VolumeMounts: volumeMounts_s,
				Command: command_s,
				Args: args_s,
			},
		}

		input_transformer := ks.TransformerSpec{
			PodSpec: ks.PodSpec{
				Containers: containers_s,
				Volumes: volume_s,
			},
		}

		isvcSpec = ks.InferenceServiceSpec{
			Predictor: input_predictor,
			Transformer: &input_transformer,
		}

	} else {
		isvcSpec = ks.InferenceServiceSpec{
			Predictor: input_predictor,
		}
	}

	inf := ks.InferenceService{
		TypeMeta: metav1.TypeMeta{
			Kind: KSERVE_KIND, 
			APIVersion: KSERVE_APIV,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: isvcName, 
			Namespace: KSERVE_NAMESPACE,
		},
		Spec: isvcSpec,
	}

	result, err = clientset.InferenceServices(KSERVE_NAMESPACE).Create(&inf)

	return
}

func Get(isvcName string) (result *ks.InferenceService, err error) {

	result, err = clientset.InferenceServices(KSERVE_NAMESPACE).Get(isvcName, metav1.GetOptions{})

	return
}

func GenIsvcName(n string, v string) string {
	const GPRCPREFIX = "grpc"
	isvcName := utils.BuildStringfunc(GPRCPREFIX, "-", n ,"-", v)
	return isvcName
}
