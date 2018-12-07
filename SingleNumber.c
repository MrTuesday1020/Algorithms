/*
https://leetcode.com/problems/single-number/
Given a non-empty array of integers, every element appears twice except for one. Find that single one.
Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
Input: [2,2,1]
Output: 1
Example 2:
Input: [4,1,2,1,2]
Output: 4
*/

#include <stdio.h>

int singleNumber(int* nums, int numsSize) {
	int k = nums[0];
	int i;
	for (i = 1; i < numsSize; i++) {
		k = (k ^ nums[i]);
	}
	return k;
}

int main(int argc, char *argv[]) {
	int arr[5] = {4,1,2,1,2};
	int res = singleNumber(arr, 5);
	
	printf("The single number is %d",res);
}