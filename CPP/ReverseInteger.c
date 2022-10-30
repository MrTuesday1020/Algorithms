#include <stdio.h>
#include <limits.h>

int reverse(int x) {
	long res = 0;
	if( -10 < x && x < 10) {
		return x;
	}

	while(x){
		int remainder = x % 10;
		x = x / 10;
		res = res * 10 + remainder;
	}
	
	if(res > INT_MAX || res < INT_MIN){
		return 0;
	}
	
	return res;
}

int main(int argc, char *argv[]) {
	int t = -2147483648;
	int res = reverse(t);
	printf("%d\n",res);
	
	return 0;
}