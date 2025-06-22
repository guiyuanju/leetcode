from typing import *

def assert_eq(a, b):
    if a == b:
      print("Correct")
    else:
      print(f"Wrong: {a} != {b}")

class Solution:
    def matchPlayersAndTrainers(self, players: List[int], trainers: List[int]) -> int:
        players.sort()
        trainers.sort()
        res, i, j = 0, 0, 0
        while i < len(players) and j < len(trainers):
            if trainers[j] >= players[i]:
                res += 1
                i += 1
            j += 1
        return res

players = [4,7,9]
trainers = [8,2,5,8]
assert_eq(2, Solution().matchPlayersAndTrainers(players, trainers))

players = [1,1,1]
trainers = [10]
assert_eq(1, Solution().matchPlayersAndTrainers(players, trainers))
