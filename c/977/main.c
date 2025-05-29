#include <stdlib.h>
#include <stdio.h>

int* sortedSquares(int* nums, int numsSize, int* returnSize) {
    int* res = (int*)malloc(sizeof(int) * numsSize);
    *returnSize = numsSize;
    int i = 0;
    int j = numsSize - 1;
    while (i <= j) {
        if (abs(nums[i]) >= abs(nums[j])) {
            res[--numsSize] = nums[i] * nums[i];
            i++;
        } else {
            res[--numsSize] = nums[j] * nums[j];
            j--;
        }
    }
    return res;
}

int main() {
    int nums[] = {-4, -1, 0, 3, 10};
    int returnSize = 0;
    int* sorted = sortedSquares(nums, sizeof(nums) / sizeof(int), &returnSize);
    for (int i = 0; i < returnSize; i++) {
        printf("%d ", sorted[i]);
    }
    free(sorted);
}
