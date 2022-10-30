#include <stdio.h>
#include <stdbool.h>

bool isPalindrome(int x) {
	if( x < 0 ){
		return false;
	} else if ( 0 <= x && x <= 9){
		return true;
	} else {
		int reverse = 0;
		int num = x;
		while(num){
			int tmp = num % 10;
			num = num / 10;
			reverse = reverse * 10 + tmp;
		}
		if(reverse == x){
			return true;
		} else {
			return false;
		}
	}
}

int main(int argc, char *argv[]) {
	bool res = isPalindrome(11);
	printf("%d\n",res);
}