#!/usr/bin/env python
# -*- coding:utf-8 -*-
# Time Complexity of Solution:
# Best O(n); Average O(n^2); Worst O(n^2).

def insertionsort( aList ):
    for i in range( 1, len( aList ) ):
        tmp = aList[i]
        k = i
        while k > 0 and tmp < aList[k - 1]:
            aList[k] = aList[k - 1]
            k -= 1
        aList[k] = tmp
    return aList


if __name__ == "__main__":
    my_list = [54,26,93,17,77,31,44,55,20]
    print insertionsort(my_list)

