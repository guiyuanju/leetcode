from typing import List
import heapq

class Solution:
    def largestSumAfterKNegations(self, nums: List[int], k: int) -> int:
      heapq.heapify(nums)
      while k > 0:
         n = heapq.heappop(nums)  
         heapq.heappush(nums, -n)
         k -= 1
      return sum(nums)

def assert_eq(a, b):
    if a == b:
      print("Correct")
    else:
      print(f"Wrong: {a} != {b}")

assert_eq(5, Solution().largestSumAfterKNegations([4,2,3], 1))