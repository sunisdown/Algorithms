#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Auth Sunisdown
# union find 的一种方法，quick union

class QuickUnion(object):

    def __init__(self, N):
        self.id_list = [i for i in xrange(N)]
        self.count = 10

    def find(self, p):
        while p != self.id_list[p]:
            p = self.id_list[p]
        return p

    def connect(self, p, q):
        return self.find(p) == self.find(q)

    def union(self, p, q):
        print  p, p, self.id_list, self.count
        p_root = self.find(p)
        q_root = self.find(q)
        if q_root == p_root:
            return True
        self.id_list[p_root] = q_root
        self.count -= 1
        return False

if __name__ == "__main__":
    uf = QuickUnion(10)
    from random import randint
    for _ in xrange(100):
        a = randint(0, 9)
        b = randint(0, 9)
        if uf.connect(a, b):
            print True, a,b
            continue
        uf.union(a, b)
        print a,b

