#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Auth Sunisdown

class QuickFind(object):
    def __init__(self):
        self.id_list = [i for i in xrange(10)]
        self.count = 10

    def find(self, q):
        return self.id_list[q]

    def union(self, p, q):
        print q, p, self.id_list
        p_id = self.find(p)
        q_id = self.find(q)

        if q_id == p_id:
            return True

        for i in xrange(len(self.id_list)):
            if self.id_list[i] == q_id:
                self.id_list[i] = p_id
        self.count -= 1
        return False

if __name__ == "__main__":
    a = QuickFind()
    from random import randint
    for _ in xrange(10):
        print a.union(randint(0, 9), randint(0, 9))

