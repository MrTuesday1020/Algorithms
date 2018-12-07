/*
	https://leetcode.com/problems/majority-element/
	Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
	You may assume that the array is non-empty and the majority element always exist in the array.

	Example 1:
	Input: [3,2,3]
	Output: 3
	Example 2:
	Input: [2,2,1,1,1,2,2]
	Output: 2
*/

#include <stdio.h>
#include <stdlib.h>

int majorityElement(int* nums, int numsSize) {
	// solution1: O(n) time, O(n) space
//	int* stack = malloc(sizeof(int) * numsSize);
//	int top = -1;
//	int i;
//	for (i=0;i<numsSize;i++) {
//		if(top == -1){
//			stack[++top] = nums[i];
//		}
//		else if(stack[top] == nums[i]){
//			stack[++top] = nums[i];
//		}
//		else{
//			top --;
//		}
//	}
//	return stack[0];
	
	// solution2: O(n) time, O(1) space
	int candidate;
	int count = 0;
	int i;
	for (i=0;i<numsSize;i++) {
		if(count == 0){
			candidate = nums[i];
			count ++;
		}
		else if(candidate == nums[i]){
			count ++;
		}
		else{
			count --;
		}
	}
	return candidate;	
}

int main(int argc, char *argv[]) {
	int nums[] = {1,1,1,1,2,3,4};
	int result =  majorityElement(nums, 6);
	printf("The majority number is %d",result);
	return 0;
}