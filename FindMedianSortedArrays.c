#include <stdio.h>
#include <stdlib.h>

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {
	int numsSize = nums1Size + nums2Size;
	int* arr = (int*)malloc(sizeof(int) * numsSize);
	int i = 0;
	int j = 0;
	while (i + j < numsSize) {
		if(i >=  nums1Size) {
			arr[i + j] = nums2[j];
			j ++;
		}
		else if(j >=  nums2Size) {
			arr[i + j] = nums1[i];
			i ++;
		}
		else{
			if(nums1[i] < nums2[j]){
				arr[i + j] = nums1[i];
				i ++;
			} else {
				arr[i + j] = nums2[j];
				j ++;
			}	
		}
	}
	
	double median;
	if(numsSize % 2 == 1){
		median =  arr[(numsSize - 1) / 2];
	} else {
		median = (arr[numsSize / 2] + arr[numsSize / 2 - 1]) / 2.0;
	}
	
	free(arr);
	return median;
}

int main(int argc, char *argv[]) {
	int nums1[] = {1};
	int nums2[] = {1};
	
	double median = findMedianSortedArrays(nums1, 1, nums2, 1);
	printf("%f\n", median);
}