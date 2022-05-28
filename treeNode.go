package collections

type TreeNode[M Mix] struct {
	Value M            `json:"value"`
	Left  *TreeNode[M] `json:"left"`
	Right *TreeNode[M] `json:"right"`
}

func NewTreeNode[M Mix](value M) *TreeNode[M] {
	return &TreeNode[M]{Value: value}
}
