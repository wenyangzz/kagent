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

import kserve
from kserve.model import ModelMissingError, InferenceError
import joblib
import pathlib
import xgboost as xgb
import numpy as np
from xgboost import XGBModel
import os
from typing import Dict

MODEL_BASENAME = "model"
MODEL_EXTENSIONS = [".joblib", ".pkl", ".pickle", ".json"]


class XGBoostModel(kserve.Model):
    def __init__(self, name: str, model_dir: str, nthread: int,
                 booster: XGBModel = None):
        super().__init__(name)
        self.name = name
        self.model_dir = model_dir
        self.nthread = nthread
        self._predict_dmatrix = False
        if booster is not None:
            self._booster = booster
            self.ready = True

    def load(self) -> bool:
        model_path = pathlib.Path(kserve.Storage.download(self.model_dir))
        paths = [os.path.join(model_path, MODEL_BASENAME + model_extension) for model_extension in MODEL_EXTENSIONS]
        existing_paths = [path for path in paths if os.path.exists(path)]
        if len(existing_paths) == 0:
            raise ModelMissingError(model_path)
        elif len(existing_paths) > 1:
            raise RuntimeError('More than one model file is detected, '
                               f'Only one is allowed within model_dir: {model_path}')
        if existing_paths[0].endswith(".json"):
            self._booster = xgb.Booster()
            self._booster.load_model(existing_paths[0])
            self._predict_dmatrix = True
        else:
            self._booster = joblib.load(existing_paths[0])
        self.ready = True
        return self.ready

    def predict(self, request: Dict) -> Dict:
        instances = request["instances"]
        try:
            if self._predict_dmatrix: # Use of list as input is deprecated see https://github.com/dmlc/xgboost/pull/3970
                dmatrix = xgb.DMatrix(np.array(request["instances"]), nthread=self.nthread)
                result = self._booster.predict(dmatrix).tolist()
            else:
                result = self._booster.predict(instances).tolist()
            return {"predictions": result}
        except Exception as e:
            raise InferenceError(str(e))
