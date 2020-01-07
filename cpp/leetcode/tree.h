#include <iostream>
#include <vector>

using namespace std;

struct TreeNode{
    int val;
    TreeNode* lchild;
    TreeNode* rchild;
    TreeNode(int i): val(i){};
};

TreeNode* getTree();

void preOrder(TreeNode* root);

void preOrderRecursive(TreeNode* root);

