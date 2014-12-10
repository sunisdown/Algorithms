#!/usr/bin/env python
# -*- coding:utf-8 -*-

import random
def quicksort(L):
    if len(L) > 1:
        pivot = random.randrange(len(L))
        elements = L[:pivot]+L[pivot+1:]
        left  = [element for element in elements if element < L[pivot]]
        right =[element for element in elements if element >= L[pivot]]
        return quicksort(left)+[L[pivot]]+quicksort(right)
    return L

alist = [54,26,93,17,77,31,44,55,20]
print quicksort(alist)
