package parserlib

import (
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

type Node struct {
	node *html.Node
}

func NewNode(htmlNode *html.Node) *Node {
	return &Node{node: htmlNode}
}

func (n *Node) Query(selector string, filters ...filterFunc) (*NodeList, error) {
	htmlNodes, err := htmlquery.QueryAll(n.node, selector)
	if err != nil {
		return nil, fmt.Errorf("selector failed: %w", err)
	}

	nodes := NewNodeList(htmlNodes)

	for _, callback := range filters {
		nodes = nodes.Filter(callback)
	}

	return nodes, nil
}

func (n *Node) Attrs() map[string]string {
	values := make(map[string]string, len(n.node.Attr))

	for _, attr := range n.node.Attr {
		values[attr.Key] = attr.Val
	}

	return values
}

func (n *Node) Attr(key string) (value string, found bool) {
	for _, attr := range n.node.Attr {
		if attr.Key == key {
			value = attr.Val
			found = true
			break
		}
	}

	return
}

func (n *Node) Value() string {
	return n.node.Data
}

func (n *Node) FirstChild() *Node {
	if n.node.FirstChild != nil {
		return NewNode(n.node.FirstChild)
	}

	return nil
}

func (n *Node) PrintTree() {
	printTree(n.node, "")
}

type NodeList struct {
	nodes []*Node
}

func NewNodeList(htmlNodes []*html.Node) *NodeList {
	list := make([]*Node, len(htmlNodes))

	for i, htmlNode := range htmlNodes {
		list[i] = &Node{node: htmlNode}
	}

	return &NodeList{nodes: list}
}

type filterFunc func(node *Node) (*Node, bool)

func (n *NodeList) Filter(callback filterFunc) *NodeList {
	filtered := make([]*Node, 0)

	for _, node := range n.nodes {
		if newNode, next := callback(node); next {
			filtered = append(filtered, newNode)
		}
	}

	return &NodeList{nodes: filtered}
}

func (n *NodeList) Nodes() []*Node {
	return n.nodes
}

func (n *NodeList) First() *Node {
	if len(n.nodes) == 0 {
		return nil
	}

	return n.nodes[0]
}
