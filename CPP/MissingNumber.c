/*
https://leetcode.com/problems/missing-number/
Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:
Input: [3,0,1]
Output: 2
Example 2:
Input: [9,6,4,2,3,5,7,0,1]
Output: 8
*/

#include <stdio.h>

int missingNumber(int* nums, int numsSize) {
	int sum = numsSize * (numsSize + 1) / 2;
	int i = 0;
	for (i = 0; i < numsSize; i++) {
		sum = sum - nums[i];
	}
	return sum;
}

int main(int argc, char *argv[]) {
	int arr[5] = {1,2,3,0,5};
	int res = missingNumber(arr, 5);
	printf("The missing number is %d",res);
}