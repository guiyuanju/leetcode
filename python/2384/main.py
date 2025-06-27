from collections import Counter
from typing import *
from operator import itemgetter

class Solution:
    def largestPalindromic(self, num: str) -> str:
        counter = Counter(num)

        res = []
        for n in "9876543210":
            while counter[n] >= 2:
                res.append(n)
                counter[n] -= 2

        middle = next((n for n in "9876543210" if counter[n] > 0))

        res = res + [middle] + list(reversed(res))

        return ''.join(res)

def assert_eq(a, b):
    if a == b:
      print("Correct")
    else:
      print(f"Wrong: {a} != {b}")


assert_eq("7449447", Solution().largestPalindromic("444947137"))
assert_eq("9", Solution().largestPalindromic("00009"))
assert_eq("1005001", Solution().largestPalindromic("00001105"))
