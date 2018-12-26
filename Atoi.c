/*
https://leetcode.com/problems/string-to-integer-atoi/submissions/
*/

#include <stdio.h>
#include <limits.h>

int myAtoi(char* str) {
	char* ch = str;
	int negative = 1;
	long res = 0;
	while (*ch) {
		if( 48 <= *ch && *ch <= 57){
			break;
		} else if( *ch == '+' ){
			ch ++;
			break;
		} else if( *ch == '-' ){
			negative = -1;
			ch ++;
			break;
		} else if( *ch == ' ' ){
			ch ++;
		} else {
			return 0;
		}
	}
	
	while ( 48 <= *ch && *ch <= 57) {
		int num = *ch - '0';
		res = res * 10 + num;
		if(res > INT_MAX && negative == 1){
			return INT_MAX;
		} else if (res > INT_MAX && negative == -1){
			return INT_MIN;
		}
		ch ++;
	}

	return res * negative;
}

int main(int argc, char *argv[]) {
	char* arr = "9223372036854775808";
	int res = myAtoi(arr);
	printf("%d\n", res);
}