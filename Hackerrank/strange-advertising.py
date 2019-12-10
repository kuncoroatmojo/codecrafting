#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the viralAdvertising function below.
def liked(n):
    if n == 1 :
        return 2
    return (liked(n-1)*3//2)
def viralAdvertising(n):
    if n == 1 :
        return 2
    return(liked(n) + viralAdvertising(n-1))
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    result = viralAdvertising(n)

    fptr.write(str(result) + '\n')

    fptr.close()
