package kmodel

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"bytes"
	"log"
	"os"
	"time"
	"context"
	"github.com/codeclysm/extract/v3"
	"github.com/kagent/internal/utils"
)

const (
	KMODEL_PATH_PREFIX = "/pv/"
	//KMODEL_PATH_PREFIX   = "/Users/coco/pv/"
	KMODEL_META_FILE = "/pv/kmodel_meta.json"
	PREPARATION_FOLDER_NAME = "data-preparation"
	MODEL_FOLDER_NAME = "model"
)

type kModel struct {
	ModelName    			string `json:"model_name"`
	ModelVersion 			string `json:"model_ver"`
	ModelType    			string `json:"model_type"`
	ModelTypeVer 			string `json:"model_type_ver"`
	ModelHasPreparation 	bool   `json:"model_has_preparation"`
	ModelPath    			string `json:"model_path"`
	ModelContentPath    	string `json:"model_content_path"`
	ModelPreparationPath    string `json:"model_preparation_path"`
	ModelContentTgzName    	string `json:"model_content_tgz_name"`
	ModelPreparationTgzName string `json:"model_preparation_tgz_name"`
	Timestamp    			string `json:"model_timestamp"`
}

type kModelslice struct {
	KModels []kModel `json:"models"`
}

func init() {
	if _, err := os.Stat(KMODEL_META_FILE); err == nil {
		log.Printf("find the kmodel meta json file %v", KMODEL_META_FILE)
	} else if errors.Is(err, os.ErrNotExist) {
		var k kModelslice
		err = saveKmodelJson(&k)
		if err != nil {
			log.Printf("save the kmodel meta json file error: %v", err)
		}
	} else {
		log.Printf("find the kmodel meta json file error: %v", err)
	}
}

func GetKModel(n string, v string) (kModel, error) {
	m, err := getKmodelMeta(n, v)
	if err != nil {
		log.Printf("undeployed model %s %s\n", n, v)
		return kModel{}, err
	}
    
	err = checkKmodelFile(m.ModelPath)

	if err != nil {
		log.Printf("error find model %s %s\n", n, v)
		return kModel{}, err
	}

	return m, nil
}

func getKmodelMeta(n string, v string) (kModel, error) {
	var k kModelslice
	var km kModel
	if err := getKmodelJson(&k); err != nil {
		log.Printf("error get json file %v", err)
		return km, err
	}

	for _, m := range k.KModels {
		if m.ModelName == n && m.ModelVersion == v {
			log.Printf("the model is %v", m)
			if err := checkKmodelFile(m.ModelPath); err != nil {
				return km, fmt.Errorf("error find the model file %v %v", n, v)
			}
			return m, nil
		}
	}

	return km, fmt.Errorf("can not find the model %v %v", n, v)
}

func getKmodelJson(k *kModelslice) error {
	jsonFile, err := os.Open(KMODEL_META_FILE)
	if err != nil {
		log.Printf("error opening json file %v", err)
		return err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Printf("error reading json file %v", err)
		return err
	}

	json.Unmarshal(jsonData, k)

	return nil
}

func saveKmodelJson(k *kModelslice) (err error) {
	marshal, err := json.Marshal(*k)
	if err != nil {
		log.Printf("kmodel meta json marshal error: %v", err)
		return err
	}
	create, err := os.Create(KMODEL_META_FILE)
	if err != nil {
		log.Printf("kmodel meta json file create error: %v", err)
		return err
	}

	_, err = create.Write(marshal)
	if err != nil {
		log.Printf("kmodel meta json file write error: %v", err)
		return err
	}

	defer create.Close()
	log.Printf("debug: write json %v", string(marshal))
	return nil
}

func checkKmodelFile(p string) (err error) {
	if _, err := os.Stat(KMODEL_PATH_PREFIX + p); err != nil {
		log.Printf("find the model file %v", KMODEL_PATH_PREFIX + p)
		return err
	}
	return nil
}

func SaveKModel(n string, v string, t string, tv string, m *[]byte, p *[]byte) (kModel, error) {
	log.Printf("start save model %s %s %s %s", n, v, t, tv)
	var hasPreparation bool
	log.Printf("hasPreparation is %v", hasPreparation)
	if *p != nil {
       hasPreparation = true
	   log.Printf("hasPreparation is %v", hasPreparation)
	}
	log.Printf("hasPreparation is %v", hasPreparation)
	km, err := saveModelMeta(n, v, t, tv, hasPreparation)
	if err != nil {
		log.Printf("failed to save model json %v", err)
		return kModel{}, err
	}

	local_m_path := utils.BuildStringfunc(KMODEL_PATH_PREFIX, km.ModelContentPath)
	if err := saveTgzFile(local_m_path, km.ModelContentTgzName, m); err != nil {
		log.Printf("save model tgz error: %v", err)
		return kModel{}, err
	}
	
	if km.ModelHasPreparation {
		local_p_path := utils.BuildStringfunc(KMODEL_PATH_PREFIX, km.ModelPreparationPath)
		if err := saveTgzFile(local_p_path, km.ModelPreparationTgzName, p); err != nil {
			log.Printf("save preparation tgz error: %v", err)
			return kModel{}, err
		}
	}

	return km, nil
}

func saveModelMeta(n string, v string, t string, tv string, hasPreparation bool) (kModel, error) {
	var k kModelslice
	if err := getKmodelJson(&k); err != nil {
		log.Printf("error get json file when save model meta %v", err)
		return kModel{}, err
	}

	p := utils.BuildStringfunc(n, "-", v)
	km := kModel{
		ModelName:    				n,
		ModelVersion: 				v,
		ModelType:   				t,
		ModelTypeVer: 				tv,
		ModelHasPreparation: 		hasPreparation,
		ModelPath:    				p,
		ModelContentPath:   		utils.BuildStringfunc(p, "/", MODEL_FOLDER_NAME),
		ModelPreparationPath:		utils.BuildStringfunc(p, "/", PREPARATION_FOLDER_NAME),
		ModelContentTgzName:	 	utils.BuildStringfunc(n, "-", v, "-", t, "-", tv, "-model.tgz"),
		ModelPreparationTgzName:	utils.BuildStringfunc(n, "-", v, "-", t, "-", tv, "-data-preparation.tgz"),
		Timestamp:    time.Now().Format(time.RFC3339),
	}
	k.KModels = append(k.KModels, km)

	if err := saveKmodelJson(&k); err != nil {
		log.Printf("error save json file when save model meta %v", err)
		return kModel{}, err
	}

	return km, nil
}

func saveTgzFile(local_p string, tgz_n string, b *[]byte) (err error) {
	if err = os.MkdirAll(local_p, 0755); err != nil {
		log.Printf("kmodel file dir create error: %v", err)
		return err
	}

	tgz_p := utils.BuildStringfunc(local_p, "/", tgz_n)
	log.Printf("kmodel file dir : %v", tgz_p)

	create, err := os.Create(tgz_p)
	if err != nil {
		log.Printf("kmodel file create error: %v", err)
		return err
	}

	_, err = create.Write(*b)
	if err != nil {
		log.Printf("kmodel file write error: %v", err)
		return err
	}

	defer create.Close()

	buffer := bytes.NewBuffer(*b)
	log.Printf("extract path: %s", local_p)
	err = extract.Gz(context.TODO(), buffer, local_p, nil)
	if err != nil {
		log.Printf("extract kmodel file error: %v", err)
		return err
	}

	return nil

}
