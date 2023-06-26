package main

import (
    "math"
)

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func findTilt(root *TreeNode){
    leftVal := 0
    rightVal := 0

    if root.Left != nil {
        leftVal = root.Left.Val
        leftVal = leftVal + findTilt(root.Left)
    }

    if root.Right != nil {
        rightVal = root.Right.Val
        rightVal = rightVal + findTilt(root.Right)
    }

    if leftVal < rightVal {
        return rightVal - leftVal
    } else {
        return leftVal - rightVal
    }
}
