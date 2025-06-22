def assert_eq(a, b):
    if a == b:
      print("Correct")
    else:
      print(f"Wrong: {a} != {b}")

class Solution:
    def getSmallestString(self, n: int, k: int) -> str:
        res = [''] * n
        for i in reversed(range(0, n)):
            cur = min(26, k-(n-(n-i)))
            res[i] = chr(ord('a')+cur-1)
            k -= cur
        return ''.join(res)

assert_eq("aay", Solution().getSmallestString(3, 27))
assert_eq("aaszz", Solution().getSmallestString(5, 73))
assert_eq("vzzz", Solution().getSmallestString(4, 100))
assert_eq("zzzzz", Solution().getSmallestString(5, 130))
