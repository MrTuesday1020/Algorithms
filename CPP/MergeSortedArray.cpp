#include <iostream>
#include <vector>
using namespace std;


class Solution {
public:
	void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
		int i = m - 1, j = n - 1, tar = m + n - 1;
		while (j >= 0) {
			if(i >= 0 && nums1[i] > nums2[j])
				nums1[tar--] = nums1[i--];
			else
				nums1[tar--] = nums2[j--];
		}
	}
};

int main(int argc, char *argv[]) {
	int arr1[] = {1,2,3,0,0,0};
	int p = sizeof(arr1) / sizeof(arr1[0]); 
	vector<int> nums1(arr1, arr1 + p); 
	int arr2[] = {2,5,6};
	int q = sizeof(arr2) / sizeof(arr2[0]); 
	vector<int> nums2(arr2, arr2 + q); 
	Solution s;
	int m = 3;
	int n = 3;
	s.merge(nums1, m, nums2, n);
	for (int x : nums1) 
		cout << x << " "; 
}