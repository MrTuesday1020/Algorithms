/*
	https://leetcode.com/problems/valid-anagram/
	Given two strings s and t , write a function to determine if t is an anagram of s.

	Example 1:
	Input: s = anagram, t = nagaram
	Output: true
	
	Example 2:
	Input: s = rat, t = car
	Output: false
	
	Note:
	You may assume the string contains only lowercase alphabets.
*/

#include <stdio.h>
#include <stdbool.h>
#include <string.h>

bool isAnagram(char* s, char* t) {
	int statS[26] = {0};
	int statT[26] = {0};
	int lenS = strlen(s);
	int lenT = strlen(t);
	int i=0;
	for (i=0;i<lenS;i++) {
		int id = s[i] - 'a';
		statS[id]++;
	}
	for (i=0;i<lenS;i++) {
		int id = t[i] - 'a';
		statT[id]++;
	}
	for (i=0;i<26;i++) {
		if(statS[i] != statT[i]){
			return false;
		}
	}
	return true;
}

int main(){
	char* s = "hello";
	char* t = "helol";
	bool res = isAnagram(s,t);
	printf("%d",res);
	return 0;
}