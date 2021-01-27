"""
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
"""
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left: TreeNode = left
        self.right: TreeNode = right


class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        return self.valid(node=root)

    def valid(
        self,
        node: Optional[TreeNode],
        min: Optional[int] = None,
        max: Optional[int] = None,
    ) -> bool:

        if not node:
            return True

        if min is not None and node.val <= min:
            return False

        if max is not None and node.val >= max:
            return False

        return self.valid(node=node.left, min=min, max=node.val) and self.valid(
            node=node.right, min=node.val, max=max
        )
