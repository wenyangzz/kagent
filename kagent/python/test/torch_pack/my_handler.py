import torch
import logging
from ts.torch_handler.base_handler import BaseHandler

logger = logging.getLogger(__name__)


class MyHandler(BaseHandler):
    """
    do nothing
    """


    def preprocess(self, data):
        logger.info('input data for handler.preprocess:')
        logger.info(data)
        return super().preprocess(data)

    def postprocess(self, data):
        logger.info('model result:')
        logger.info(data)
        return super().postprocess(data)

