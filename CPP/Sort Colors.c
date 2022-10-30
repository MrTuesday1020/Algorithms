/*
https://leetcode.com/problems/sort-colors/
*/

#include <stdio.h>

void sortColors(int* nums, int numsSize) {
	int red = 0;
	int white = 0;
	int blue = 0;
	
	int i;
	for(i = 0; i < numsSize; i++) {
		if(nums[i] == 0) {
			red ++;
		} else if(nums[i] == 1) {
			white ++;
		} else {
			blue ++;
		}
	}
	
	for(i = 0; i < numsSize; i++) {
		if(red > 0) {
			red --;
			nums[i] = 0;
		} else if(white > 0) {
			white --;
			nums[i] = 1;
		} else {
			nums[i] = 2;
		}
	}
}

int main(int argc, char *argv[]) {
	int numsSize = 6;
	int nums[] = {2,0,2,1,1,0};
	sortColors(nums, numsSize);
	
	int i;
	for(i = 0; i < numsSize; i++) {
		printf("%d", nums[i]);
	}
}