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

from .eie_transformer import EIETransformer


parser = argparse.ArgumentParser(parents=[kserve.model_server.parser])
parser.add_argument(
    "--predictor_host",
    help="The URL for the model predict function", required=True
)
parser.add_argument(
    "--protocol", required=True,
    choices=["http", "grpc-v2"], help="The protocol for the predictor"
)
parser.add_argument(
    "--model_name", required=True,
    help='The name that the model is served under.'
)
parser.add_argument(
    "--signature_name",
    help='Model signature', required=True
)
parser.add_argument(
    "--user_package",
    help='User defined package', required=True
)
parser.add_argument(
    "--user_file",
    help='User defined file', required=True
)
parser.add_argument(
    "--data_path",
    help='local data path', required=True
)
args, _ = parser.parse_known_args()

if __name__ == "__main__":
    transformer = EIETransformer(
        name=args.model_name,
        signature_name=args.signature_name,
        predictor_host=args.predictor_host,
        protocol=args.protocol,
        user_package=args.user_package,
        user_file=args.user_file,
        data_path=args.data_path)
    server = kserve.ModelServer()
    server.start(models=[transformer])
