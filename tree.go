package forestry

import "strings"

/*
	A Tree should have all of its data stored in its trunk. When a value
	needs to be read from the trunk, a branch is sprouted out which will
	keep that value for a configured duration
*/

type Tree struct {
	Id       uint64                 `json:"id"`
	Name     string                 `json:"name"`
	Branches map[string]*Branch     `json:"branches"`
	Trunk    map[string]interface{} `json:"trunk"`
	Roots    []*Root                `json:"roots"`
}

type Branch struct {
	Id   uint64      `json:"id"`
	Leaf interface{} `json:"leaf"`
}

func (tree *Tree) SproutBranch(key string) *Branch {
	var branch *Branch
	var ok bool

	if branch, ok = tree.Branches[key]; !ok {
		leaf := tree.readTree(key)
		if leaf != nil {
			branch = &Branch{tree.generateId(), leaf}
			tree.Branches[key] = branch
		}
	}

	return branch
}

func (tree *Tree) readTree(key string) interface{} {
	tokens := strings.Split(key, ".")
	data := tree.Trunk
	var value interface{}

	for k, v := range tokens {
		if obj, ok := data[v]; ok {
			if d, ok := obj.(map[string]interface{}); ok {
				data = d
				value = d
			} else {
				if k == len(tokens)-1 {
					value = d
				} else {
					return nil
				}
			}
		}
	}
	return value
}

func (tree *Tree) tokenizeKey(key string) []string {
	return strings.Split(key, ".")
}

func (tree *Tree) generateId() uint64 {
	return uint64(len(tree.Branches) + 1)
}
