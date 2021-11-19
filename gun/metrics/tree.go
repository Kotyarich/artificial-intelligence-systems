package metrics

import "ais.com/m/model"

type Tree struct {
	Val float64

	Children map[string]Tree
}

var typeTraces = map[string][]string{
	"traumatic": {"traumatic"},
	"sport pistol": {"firearm", "pistol"},
	"rifle": {"firearm", "rifle"},
	"shotgun": {"firearm", "shotgun"},
}

func createTree() Tree {
	root := Tree{0, map[string]Tree{
		"traumatic": {0, nil},
		"firearm": {3, map[string]Tree{
			"pistol": {5, nil},
			"rifle": {3, map[string]Tree{
				"lever": {0.5, nil},
				"semi-auto": {0, nil},
				"bolt-action": {1, nil},
			}},
			"shotgun": {0, map[string]Tree{
				"semi-auto": {2, nil},
				"pomp": {0, nil},
				"break": {0.5, map[string]Tree{
					"horizontal": {0, nil},
					"vertical": {0.1, nil},
					"other": {0.2, nil},
				}},
			}},
		}},
	}}

	return root
}

func TreePathSum(gun model.Gun) float64 {
	sum := 0.
	tree := createTree()

	for _, v := range typeTraces[gun.Type] {
		tree = tree.Children[v]
		sum += tree.Val
	}

	if tree.Children == nil {
		return sum
	}

	tree = tree.Children[gun.LoadType]
	sum += tree.Val
	if gun.LoadType == "break" {
		tree = tree.Children[gun.BreakType]
		sum += tree.Val
	}

	return sum
}