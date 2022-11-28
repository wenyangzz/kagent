import torch
import torch.nn as nn
from torch.nn.utils import weight_norm
from network import TemporalConvNet


input_size = 1
output_size = 96
num_channels = [96] * 9
kernel_size = 3
dropout = 0.2


class TCN(nn.Module):
    def __init__(self, input_size=input_size, output_size=output_size, num_channels=num_channels, kernel_size=kernel_size, dropout=dropout):
        super(TCN, self).__init__()
        self.tcn = TemporalConvNet(input_size, num_channels, kernel_size=kernel_size, dropout=dropout)
        self.linear = nn.Linear(num_channels[-1], output_size)
        self.init_weights()

    def init_weights(self):
        self.linear.weight.data.normal_(0, 0.01)

    def forward(self, x):
        y1 = self.tcn(x)
        return self.linear(y1[:, :, -1])

