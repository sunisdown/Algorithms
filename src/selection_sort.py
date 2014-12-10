#!/usr/bin/env python
# -*- coding:utf-8 -*-

def sort_selection(my_list):

    for pos_upper in xrange( len(my_list)-1, 0, -1):
        max_pos = 0
        for i in xrange(1, pos_upper + 1):
            if(my_list[i] > my_list[max_pos]):
                max_pos = i
                print "resetting max_pos = " + str(max_pos)

        my_list[pos_upper], my_list[max_pos] = my_list[max_pos], my_list[pos_upper]
        print "pos_upper: " + str(pos_upper) + " max_pos: " + str(max_pos) + " my_list: " + str(my_list)
    
    return my_list

def ssort(V):
  j = 0
  while j != len(V):
      for i in range(j, len(V)):
          if V[i] < V[j]:
              V[j],V[i] = V[i],V[j]
      j = j+1

  return V  


if __name__ == "__main__":

    my_list = [54,26,93,17,77,31,44,55,20]
    print "my_list: " + str(my_list)
    print sort_selection(my_list)
    print ssort(my_list)

