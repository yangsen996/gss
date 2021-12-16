package gss

type node struct {
	pattern  string  //带匹配的路由
	part     string  //路由中的一部门
	children []*node //子节点
	isWild   bool    //是否精准匹配，含有：*时候未true
}

// matchChild 第一个匹配成功的节点
func (n *node) matchChild() *node {
	for _, child := range n.children {
		if child.part == n.part || child.isWild {
			return child
		}
	}
	return nil
}
