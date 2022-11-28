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

import argparse
import kserve
import os
import sys
import requests
parent_path = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
sys.path.append(parent_path)
from transformer.eie_transformer.eie_transformer import EIETransformer

predictor_host="mm"
protocol="http"
model_name="mm"
signature_name="mm"
user_package="user_dir"
user_file="my_tf_transformer.py"
data_path="local_data"

# remote inference service
ip = '10.58.96.65'
port = '80'
model_name = 'http-tf-0906'
hostname = 'http-tf-0906.kagent.example.com'
headers = {"Host": hostname}
url = 'http://%s:%s/v1/models/%s:predict' % (ip, port, model_name)

if __name__ == "__main__":
    transformer = EIETransformer(
        name=model_name,
        signature_name=signature_name,
        predictor_host=predictor_host,
        protocol=protocol,
        user_package=user_package,
        user_file=user_file,
        data_path=data_path)
    fp = '/home/zhangjunze/data_process/raw_data_to_test.csv'
    RAW_COLUMN_NAME = ['timestamp', 'Current_Phase_Average', 'Weather_Temperature_Celsius', 'Weather_Relative_Humidity',\
        'Global_Horizontal_Radiation', 'Diffuse_Horizontal_Radiation', 'Wind_Direction', 'Weather_Daily_Rainfall']
    with open(fp, 'r') as mf:
        cols = mf.readline()
        cols = cols.strip('\r\n').split(',')
        data = mf.readline()
        data = data.strip('\r\n').split(',')
        md = {}
        for c,d in zip(cols, data):
            md[c] = d
        nl = []
        for name in RAW_COLUMN_NAME:
            nl.append(md[name])
        nd = {'instances': [nl]}
        result_of_preprocess = transformer.preprocess(nd)
        # request remote model
        print(url)
        response = requests.request(
                method = 'POST',
                url = url,
                headers = headers,
                json = result_of_preprocess
                )
        print(response.status_code)
        print(response.text)
