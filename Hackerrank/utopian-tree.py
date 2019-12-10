#!/bin/python3

import os

# Complete the utopianTree function below.
def utopianTree(n):
    result = 1
    if n == 0:
        result = 1
    elif n % 2 == 0:
        result = utopianTree(n-1) + 1
    else:
        result = utopianTree(n-1) * 2
    return result

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    t = int(input())

    for t_itr in range(t):
        n = int(input())

        result = utopianTree(n)

        fptr.write(str(result) + '\n')

    fptr.close()
