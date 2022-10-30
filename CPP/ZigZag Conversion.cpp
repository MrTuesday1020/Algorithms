class Solution {
public:
    string convert(string s, int numRows) {
        if(s.size() <= 1 || numRows == 1){
            return s;
        }
        
        vector<string> vec(numRows, "");
        int current_line = 0;
        bool direction = true;    //true:down, false:up
        
        for(char c : s){
            vec[current_line] += c;
            if(direction){
                current_line++;
            }else{
                current_line--;
            }
            
            if(current_line == 0){
                direction = true;
            }
            if(current_line == numRows - 1){
                direction = false;
            }
        }
        
        string ret = "";
        for(string str : vec){
            ret += str;
        }
        
        return ret;
    }
};