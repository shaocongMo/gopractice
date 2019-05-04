package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{}
	used := map[TreeNode]bool{}
	stack = append(stack, root)
	for {
		if len(stack) == 0 {
			break
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if used[*node] {
			result = append(result, node.Val)
			continue
		}
		used[*node] = true
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		stack = append(stack, node)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}
