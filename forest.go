package forestry

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Forest struct {
	Trees map[uint64]*Tree `json:"trees"`
	Roots map[uint64]*Root `json:"roots"`
}

type Root struct {
	Id        uint64 `json:"id"`
	StartTree *Tree  `json:"start_tree"`
	EndTree   *Tree  `json:"end_tree"`
}

func WriteForest(forest *Forest) error {
	var err error
	var data []byte
	if data, err = json.Marshal(forest); err == nil {
		if err = ioutil.WriteFile("forest.json", data, os.ModePerm); err == nil {
			return nil
		}
	}
	return err
}

func ReadForest() (*Forest, error) {
	var err error
	var data []byte
	if data, err = ioutil.ReadFile("forest.json"); err == nil {
		var forest Forest
		if err = json.Unmarshal(data, forest); err == nil {
			return &forest, nil
		}
	}
	return nil, err
}

func (forest *Forest) PlantTree(name string, branches map[string]*Branch, trunk map[string]interface{}, roots []*Root) uint64 {
	tree := Tree{
		forest.generateId(),
		name,
		branches,
		trunk,
		roots,
	}

	forest.Trees[tree.Id] = &tree
	return tree.Id
}

func (forest *Forest) ChopTree(id uint64) {
	delete(forest.Trees, id)
}

func (forest *Forest) GetTreeById(id uint64) *Tree {
	return forest.Trees[id]
}

func (forest *Forest) GetTreeByName(name string) *Tree {
	for _, tree := range forest.Trees {
		if tree.Name == name {
			return tree
		}
	}
	return nil
}

func (forest *Forest) generateId() uint64 {
	return uint64(len(forest.Trees) + 1)
}
