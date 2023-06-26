package main

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
func findTilt(root *TreeNode) int {
    tiltSum := 0

    if root == nil {
        return tiltSum
    }

    evalBranch(root, &tiltSum)

    return tiltSum
}

func evalBranch(branch *TreeNode, tiltSum *int) int {
    leftVal := 0
    rightVal := 0

    if branch.Left != nil {
        leftVal = branch.Left.Val

        evalLeft = evalBranch(branch.Left, tiltSum)

        leftVal = leftVal + evalLeft
    }

    if branch.Right != nil {
        rightVal = branch.Right.Val
        evalRight = evalBranch(branch.Right, tiltSum)

        rightVal = rightVal + evalRight
    }

    tilt := 0

    if leftVal < rightVal {
        tilt = rightVal - leftVal
    } else {
        tilt =  leftVal - rightVal
    }

    tiltSum = tiltSum + tilt

    return leftVal + rightVal
}

