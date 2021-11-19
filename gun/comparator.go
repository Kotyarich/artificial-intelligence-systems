package gun

import "ais.com/m/model"

type Comparator interface {
	ClosestToOne(gun *model.Gun, n int) []*model.Gun
	ClosestToSlice(guns []*model.Gun, n int) []*model.Gun
	Closest(guns []*model.Gun, dislikes []*model.Gun, n int) []*model.Gun
}