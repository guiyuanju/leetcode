from typing import *

def maxArea(height: List[int]) -> int:
    res = 0
    i, j = 0, len(height)-1
    while i < j:
        w = j - i
        hi = height[i]
        hj = height[j]
        h = min(hi, hj)
        res = max(res, w*h)
        if hi < hj:
            while i < j and height[i] <= hi:
                i += 1
        else:
            while i < j and height[j] <= hj:
                j -= 1
    return res

def assert_eq(a, b):
    if a == b:
      print("Correct")
    else:
      print(f"Wrong: {a} != {b}")

assert_eq(49, maxArea([1,8,6,2,5,4,8,3,7]))
assert_eq(1, maxArea([1,1]))
