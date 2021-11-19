package recomendation

import (
	"ais.com/m/database"
	"ais.com/m/gun/metrics"
	"ais.com/m/model"
	"sort"
)

type Metric struct {
	treeMetric float64
	euclidMetric float64
}

func (m *Metric) less(other *Metric) bool {
	if m.treeMetric != other.treeMetric {
		return m.treeMetric < other.treeMetric
	}
	return m.euclidMetric < other.euclidMetric
}

type Comparator struct {
	weights []float32
}

func InitComparator(weights []float32) *Comparator {
	if weights == nil {
		return &Comparator{weights: []float32{1, 10, 5, 0.05, 15, 15, 2, 10, 15, 1, 15, 5, 0.005, 0.05, 1}}
	}

	return &Comparator{weights:weights}
}

func (c *Comparator) ClosestToOne(gun *model.Gun, n int) []model.Gun {
	pg := database.GetDB()

	var allGuns []model.Gun
	pg.Not(gun.ID).Find(&allGuns)

	sort.Slice(allGuns, func(i, j int) bool {
		treeMetric1 := metrics.TreeMetric(*gun, allGuns[i])
		treeMetric2 := metrics.TreeMetric(*gun, allGuns[j])

		if treeMetric1 != treeMetric2 {
			return treeMetric1 < treeMetric2
		}

		return metrics.Euclidean(*gun, allGuns[i], c.weights) < metrics.Euclidean(*gun, allGuns[j], c.weights)
	})

	return allGuns[:n]
}

func (c *Comparator) minMetric(gun *model.Gun, guns []*model.Gun, metric func(gun1 model.Gun, gun2 model.Gun) float64) float64 {
	min := metric(*gun, *guns[0])

	for _, g := range guns {
		m := metric(*gun, *g)
		if m < min {
			min = m
		}
	}

	return min
}

func (c *Comparator) closestToSlice(guns []*model.Gun, all *[]*model.Gun) {
	euclidMetric := func(gun1 model.Gun, gun2 model.Gun) float64 {
		return metrics.Euclidean(gun1, gun2, c.weights)
	}

	sort.Slice(*all, func(i, j int) bool {
		tree1 := c.minMetric((*all)[i], guns, metrics.TreeMetric)
		tree2 := c.minMetric((*all)[j], guns, metrics.TreeMetric)
		if tree1 != tree2 {
			return tree1 < tree2
		}

		euclid1 := c.minMetric((*all)[i], guns, euclidMetric)
		euclid2 := c.minMetric((*all)[j], guns, euclidMetric)
		return euclid1 < euclid2
	})
}

func (c *Comparator) getMetrics(guns []*model.Gun, all []*model.Gun) []Metric {
	euclidMetric := func(gun1 model.Gun, gun2 model.Gun) float64 {
		return metrics.Euclidean(gun1, gun2, c.weights)
	}

	gunsMetric := make([]Metric, len(all))
	for i, g := range all {
		gunsMetric[i].treeMetric = c.minMetric(g, guns, metrics.TreeMetric)
		gunsMetric[i].euclidMetric = c.minMetric(g, guns, euclidMetric)
	}
	return gunsMetric
}

func (c *Comparator) ClosestToSlice(guns []*model.Gun, n int) []*model.Gun {
	pg := database.GetDB()
	ids := make([]int64, len(guns))
	for i, g := range guns {
		ids[i] = int64(g.ID)
	}

	var allGuns []*model.Gun
	pg.Not(ids).Find(&allGuns)

	c.closestToSlice(guns, &allGuns)

	return allGuns[:n]
}

func index(gun *model.Gun, all []*model.Gun) int {
	for i, g := range all {
		if gun.ID == g.ID {
			return i
		}
	}

	return -1
}

func (c *Comparator) Closest(guns []*model.Gun, dislikes []*model.Gun, n int) []*model.Gun {
	pg := database.GetDB()
	ids := make([]int64, len(guns) + len(dislikes))
	for i, g := range guns {
		ids[i] = int64(g.ID)
	}
	for i, g := range dislikes {
		ids[len(guns) + i] = int64(g.ID)
	}

	var allGuns []*model.Gun
	pg.Not(ids).Find(&allGuns)

	c.closestToSlice(guns, &allGuns)
	likeMetric := c.getMetrics(guns, allGuns)

	closestToDislikes := make([]*model.Gun, len(allGuns))
	copy(closestToDislikes, allGuns)

	c.closestToSlice(dislikes, &closestToDislikes)
	disMetric := c.getMetrics(dislikes, closestToDislikes)

	for i, m := range disMetric {
		if len(allGuns) < n {
			break
		}

		j := index(closestToDislikes[i], allGuns)
		if j >= 0 && m.less(&likeMetric[j]) {
			allGuns = append(allGuns[:j], allGuns[j+1:]...)
		}
	}

	return allGuns[:n]
}
