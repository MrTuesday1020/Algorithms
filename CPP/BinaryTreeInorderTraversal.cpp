/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    void inorder(vector<int> & ret, TreeNode* root){
        if(root == nullptr){
            return;
        } else {
            inorder(ret, root->left);
            ret.push_back(root->val);
            inorder(ret, root->right);

            // 前序
            // ret.push_back(root->val);
            // inorder(ret, root->left);
            // inorder(ret, root->right);
            
            // 后续
            // inorder(ret, root->left);
            // inorder(ret, root->right);
            // ret.push_back(root->val);
        }
    }
    
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> ret;
        inorder(ret, root);
        return ret;
    }
};