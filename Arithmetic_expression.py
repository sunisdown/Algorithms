#!/usr/bin/env python
# -*- coding: utf-8 -*-
# sunisdown
# 这个代码主要实现四则运算，+-*/四中运算。

def evaluate(expression):
    ops = ("+", "-", "*", "/")
    left = "("
    right = ")"
    ops_list = [] #操作符栈
    vals_list = [] #数值栈
    for i in expression:
        if not i: # 排除空字符串
            pass
        elif i == left: #如果是"("，则忽略不计
            pass
        if i.isdigit(): #如果是数字，则将数字压入数值栈
            vals_list.append(eval(i))
        if i in ops: #如果是操作符，则将操作符压入操作符栈
            ops_list.append(i)
        elif i == right: # 如果是")"，则取出一个操作符与两个数值进行计算
            op = ops_list.pop()
            val = vals_list.pop()
            val_math = eval("%s %s %s" % (str(val), op, str(vals_list.pop())))
            vals_list.append(val_math)
    return vals_list.pop()


if __name__ == "__main__":
    print evaluate("(1 + ((2 + 3) * (4 * 5)))")
