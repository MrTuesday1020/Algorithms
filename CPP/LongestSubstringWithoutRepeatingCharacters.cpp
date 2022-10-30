class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int length = 0;
        int start = 0;
        unordered_map<char, int> umap;
        
        for(int i = 0; i < s.size(); i++){
            if(umap.find(s[i]) != umap.end()){
                start = max(start, (umap[s[i]] + 1));
            }
            umap[s[i]] = i;
            length = max(length, (i - start + 1));
        }
        
        return length;
    }
};