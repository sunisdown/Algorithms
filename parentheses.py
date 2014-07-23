#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Auth SunisDown
# 用 stack 来检测括号是否闭合，如果闭合则返回 True，否则 False

def parentheses(string):
    _L_PAREN = '('
    _R_PAREN = ')'
    _L_BRACKET = '['
    _R_BRACKET = ']'
    _L_BRACE= '{'
    _R_BRACE= '}'
    ops_list = [] #操作符栈
    for i in string:
        if i == _L_BRACE or i == _L_BRACKET or i == _L_PAREN:
            ops_list.append(i)
        elif i == _R_BRACE:
            if len(ops_list) < 1:
                return False
            elif ops_list.pop() != _L_BRACE:
                return False
        elif i == _R_BRACKET:
            if len(ops_list) < 1:
                return False
            elif ops_list.pop() != _L_BRACKET:
                return False
        elif i == _R_PAREN:
            if len(ops_list) < 1:
                return False
            elif ops_list.pop() != _L_PAREN:
                return False
        else: # 排除其他字符串
            pass
    return True

if __name__ == "__main__":
    print parentheses("((((())()()()(}}}{[]")
    print parentheses("()({})")




