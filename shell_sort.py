#!/usr/bin/env python
# -*- coding:utf-8 -*-

def shellSort(items):
    inc = len(items) / 2
    while inc:
        for i in xrange(len(items)):
            j = i
            temp = items[i]
            while j >= inc and items[j-inc] > temp:
                items[j] = items[j - inc]
                j -= inc
            items[j] = temp
        inc = inc/2 if inc/2 else (0 if inc==1 else 1)
    return items

if __name__ == "__main__":
    my_list = [54,26,93,17,77,31,44,55,20]
    print shellSort(my_list)
