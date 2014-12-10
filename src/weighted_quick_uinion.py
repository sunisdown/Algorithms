#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Auth Sunisdown
# union find 的一种方法，加权quick union, 可以相对避免树的不平衡

class WeightedQuickUnion(object):

    def __init__(self, N):
        self.id_list = [i for i in xrange(N)]
        self.size_list = [i for i in xrange(N)]
        self.count = 10

    def find(self, p):
        while p != self.id_list[p]:
            p = self.id_list[p]
        return p

    def connect(self, p, q):
        #print self.size_list, self.id_list
        return self.find(p) == self.find(q)

    def union(self, p, q):
        print  p, p, self.id_list, self.count
        p_root = self.find(p)
        q_root = self.find(q)
        if q_root == p_root:
            return True
        if self.size_list[p_root] < self.size_list[q_root]:
            self.id_list[p_root] = q_root
            self.size_list[q_root] += self.size_list[p_root]
        else:
            self.id_list[q_root] = p_root
            self.size_list[p_root] += self.size_list[q_root]
        self.count -= 1
        return False

if __name__ == "__main__":
    uf = WeightedQuickUnion(10)
    from random import randint
    for _ in xrange(30):
        a = randint(0, 9)
        b = randint(0, 9)
        if uf.connect(a, b):
            print True, a,b
            continue
        uf.union(a, b)
        print a, b

