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

import os
import sys
import inspect
from typing import List, Dict
import importlib
import importlib.util
import logging
import kserve
import json

logging.basicConfig(level=kserve.constants.KSERVE_LOGLEVEL)
parent_path = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
sys.path.append(parent_path)
print(sys.path)

class EIETransformer(kserve.Model):
    """EIE transformer component

    Args:
        kserve (class object): The Model class from the KServe
        module is passed here.
    """
    def __init__(self, name: str,
                signature_name: str,
                predictor_host: str,
                protocol: str,
                user_package: str,
                user_file: str,
                data_path: str):
        """Initialize the model

        Args:
            name (str): Name of the model.
            signature_name(str): Model signature.
            predictor_host (str): The host in which the predictor runs.
            protocol (str): The protocol in which the predictor runs.
            user_package: python package, used to store user files.
            user_file: python file.
            data_path: directory used to store user data.
        """
        super().__init__(name)
        self.signature_name = signature_name
        self.predictor_host = predictor_host
        self.protocol = protocol
        self.user_package = user_package
        self.user_file = user_file
        self.data_path = data_path
        self.timeout = 100
        self.user_processor = None
        try:
            user_module = importlib.import_module('.' + self.user_file[:-3], package=self.user_package)
            if user_module and inspect.ismodule(user_module):
                #user_processors = inspect.getmembers(user_module, lambda x: inspect.isclass(x) and 'preprocess' in dir(x) and 'postprocess' in dir(x))
                #context = {'data_path': os.path.join(parent_path, self.data_path)}
                #self.user_processor = user_processors[0][1](context)
                self.user_processor = user_module
            else:
                raise Exception('load user processor error, user package: %s, user file: %s', user_package, user_file)
        except Exception as e:
            raise e
        finally:
            self.do_logging()

    def preprocess(self, inputs: Dict) -> Dict:
        """data preprocess: convert raw data to appropriate format required by varied model

        Args:
            inputs (Dict): raw data for inference
        """
        logging.info("The input from request is %s", inputs)
        if self.user_processor:
            data = self.user_processor.preprocess(inputs, os.path.join(parent_path, self.data_path))
            logging.info("The input for model predict is %s", data)
            return data
        else:
            return None

    def postprocess(self, inputs: Dict) -> Dict:
        """post process: convert model results to appropriate format

        Args:
            inputs (Dict): The output of model predict
        """
        logging.info("The raw output from model predict is %s", inputs)
        if self.user_processor:
            output = self.user_processor.postprocess(inputs, os.path.join(parent_path, self.data_path))
            logging.info("The converted output from model predict is %s", output)
            return output
        '''
        if self.protocol == "grpc-v2":
            response = InferResult(inputs)
            return response.get_response(as_json=True)
        elif self.protocol == "http":
            pass
        '''
        return inputs

    def do_logging(self):
        """
        print log
        """
        logging.info("Model name = %s", self.name)
        logging.info("Signature name = %s", self.signature_name)
        logging.info("Protocol = %s", self.protocol)
        logging.info("Predictor host = %s", self.predictor_host)
        logging.info("User package = %s", self.user_package)
        logging.info("User file = %s", self.user_file)
        logging.info("Data path = %s", self.data_path)

