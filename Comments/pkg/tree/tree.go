package tree

import (
	"Comments/pkg/model"
	"errors"
	"fmt"
)

// Root - структура дерева комментариев к посту.
type Root struct {
	Comments []*Node
}

// Node - структура узла дерева комментариев.
type Node struct {
	model.Comment
	Childs []*Node `json:"Childs"`
}

var (
	ErrEmptySlice = errors.New("empty comment array")
)

// Build - функция строит полное дерево комментариев по входному слайсу.
func Build(comments []model.Comment) (Root, error) {
	const operation = "tree.Build"

	m := make(map[string]*Node)
	root := Root{}

	if len(comments) == 0 {
		return root, fmt.Errorf("%s: %w", operation, ErrEmptySlice)
	}

	for _, comment := range comments {
		node := &Node{Comment: comment}

		n, ok := m[comment.ID]
		if ok {
			node.Childs = n.Childs
		}
		m[comment.ID] = node

		if comment.ParentID == "" {
			root.Comments = append(root.Comments, node)
			continue
		}

		_, ok = m[comment.ParentID]
		if !ok {
			parent := &Node{}
			parent.Childs = append(parent.Childs, node)
			m[comment.ParentID] = parent
			continue
		}
		m[comment.ParentID].Childs = append(m[comment.ParentID].Childs, node)
	}

	return root, nil
}
