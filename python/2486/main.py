from typing import *

def assert_eq(a, b):
    if a == b:
      print("Correct")
    else:
      print(f"Wrong: {a} != {b}")

class Solution:
    def appendCharacters(self, s: str, t: str) -> int:
        i, j = 0, 0
        while i < len(s) and j < len(t):
            if s[i] == t[j]:
                j += 1
            i += 1
        return len(t) - j

assert_eq(4, Solution().appendCharacters("coaching", "coding"))
assert_eq(0, Solution().appendCharacters("abcde", "a"))
assert_eq(5, Solution().appendCharacters("z", "abcde"))
