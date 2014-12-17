#!/usr/bin/env python
# -*- coding:utf-8 -*-
# http://stackoverflow.com/questions/8279618/why-is-my-radix-sort-python-implementation-slower-than-quick-sort/8280005#8280005

from random import randint
from math import log10
from time import clock
from itertools import chain

def splitmerge1 (ls, digit): ## python (readable!)
    buf = [[] for i in range(10)]
    divisor = 10 ** digit
    for n in ls:
        buf[(n//divisor)%10].append(n)
    return chain(*buf)

def radixsort (ls, fn = splitmerge1):
    return list(reduce (fn, xrange (int (log10 (max(abs(val) for val in ls)) + 1)), ls))

if __name__=='__main__':
    for value in 1000, 10000, 100000, 1000000, 10000000:
        ls = [randint (1, value) for _ in range(value)]
        ls2 = list(ls)
        last = -1
        start = clock()
        ls = radixsort(ls)
        end = clock()
        for i in ls:
            assert last <= i
            last = i
        print("rs1 %d: %0.2fs" % (value, end-start))
