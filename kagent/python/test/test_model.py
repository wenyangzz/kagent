# Copyright 2021 The KServe Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from pkgutil import ModuleInfo
import xgboost as xgb
import os
import pandas as pd
import json
import requests
#from xgbserver import XGBoostModel
from torch_pack.prediction.data_preparation_prediction import MyTransformer


#model_dir = os.path.join(os.path.dirname(__file__), "example_model", "model")
#BST_FILE = "model.bst"
NTHREAD = 1
RAW_COLUMN_NAME = ['timestamp', 'Current_Phase_Average', 'Weather_Temperature_Celsius', 'Weather_Relative_Humidity',\
        'Global_Horizontal_Radiation', 'Diffuse_Horizontal_Radiation', 'Wind_Direction', 'Weather_Daily_Rainfall']

def test_model():
    """
    iris = load_iris()
    y = iris['target']
    X = iris['data']
    dtrain = xgb.DMatrix(X, label=y)
    param = {'max_depth': 6,
             'eta': 0.1,
             'silent': 1,
             'nthread': 4,
             'num_class': 10,
             'objective': 'multi:softmax'}
    xgb_model = xgb.train(params=param, dtrain=dtrain)
    model_file = os.path.join(model_dir, BST_FILE)
    xgb_model.save_model(model_file)
    """
    """
    model_dir = "/home/work/zhangjunze/saved_model/lightgbm_model/"
    model = LightGBMModel("model", model_dir, NTHREAD)
    model.load()
    with open("./data_process/tmp_json.csv", "r") as mf:
        data = []
        for line in mf:
            md = []
            mj = json.loads(line.strip("\r\n"))
            cols = mj["cols"]
            for col in cols:
                if col in ["Month", "Day", "Day_Of_Week", "Hour"]:
                    md.append(int(mj[col]))
                else:
                    md.append(float(mj[col]))
            data.append(md)
        res = model.predict({"inputs": data})
        print(res)
    """
    ip = '10.58.96.65'
    port = '80'
    model_name = 'kdptftransformer-0001'
    hostname = 'kdptftransformer-0001-kagent.example.com'
    headers = {"Host": hostname}
    url = 'http://%s:%s/v1/models/%s:predict' % (ip, port, model_name)
    #print(url)
    data_path = '/home/zhangjunze/data_process/raw_data_to_test.csv'
    # only one row
    #test_data = '/home/zhangjunze/demp/photovoltaic_prediction/ceshi_data_postman.txt'
    test_data = '/home/zhangjunze/demp/loading_prediction_training/day_1daylater.csv'
    context = {'data_path': 'torch_pack/data'}
    transformer = MyTransformer(context)
    '''
    with open('tmp_res.txt', 'r') as inf:
        md = inf.readlines()[0]
        md = json.loads(md)
        print(md)
        print(transformer.postprocess(md))
    '''
    url = 'http://localhost:8080/predictions/loading-prediction'
    with open(test_data, "r") as mf:
        md = mf.readlines()[1:3]
        data = []
        for l in md:
            vs = l.strip('\r\n').split(',')
            data.append(vs[:96])
        #print(data)
        #response = model.predict({"instances": request})
        mj = {'instances': data}
        res = transformer.preprocess(mj)
        print(res)
        #print(transformer.preprocess(mj))

        response = requests.request(
                method = 'POST',
                url = url,
                #headers = headers,
                json = res
                )
        print(response.status_code)
        print(response.text)



if __name__ == "__main__":
    test_model()
