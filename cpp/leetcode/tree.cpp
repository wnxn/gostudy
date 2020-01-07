#include "tree.h"

              1
            / \


TreeNode* getTree(){
    TreeNode* head = new TreeNode(1);
    head->lchild = newT
}

void preOrder(TreeNode* root){
    if(root == NULL){
        return;
    }
    vector<TreeNode*> tmp;
    for(root != NULL || tmp.size() != 0){
        if(root != NULL){
            cout << " " << root->val;
            tmp.push_back(root);
            root = root->lchild;
        }else{
            root = tmp.back();
            tmp.pop_back();
            root = root -> rchild;
        }
    }
    cout << endl;
}