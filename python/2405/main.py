from collections import defaultdict
from typing import *

def assert_eq(a, b):
    if a == b:
      print("Correct")
    else:
      print(f"Wrong: {a} != {b}")

class Solution:
    def partitionString(self, s: str) -> int:
        exist = set()
        res = 0
        for v in s:
            if v in exist:
                res += 1
                exist.clear()
            exist.add(v)
        return res + 1

assert_eq(4, Solution().partitionString("abacaba"))
assert_eq(6, Solution().partitionString("ssssss"))
