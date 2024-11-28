package intermplay

type (
	ColliderMap struct {
		data map[Vector2]bool
	}
)

func NewColliderMap() ColliderMap {
	mp := ColliderMap{
		data: make(map[Vector2]bool),
	}

	return mp
}

func (cmp *ColliderMap) DeepCopy() ColliderMap {
	result := NewColliderMap()

	for key, value := range cmp.data {
		result.data[key] = value
	}

	return result
}

func (cmp *ColliderMap) IsEmpty() bool {
	return len(cmp.data) == 0
}

func (cmp *ColliderMap) Intersect(cmp2 ColliderMap) ColliderMap {
	result := NewColliderMap()

	if len(cmp2.data) >= len(cmp.data) {
		for obj := range cmp.data {
			_, ok := cmp2.data[obj]

			if ok {
				result.data[obj] = true
			}
		}
	} else {
		for obj := range cmp2.data {
			_, ok := cmp.data[obj]

			if ok {
				result.data[obj] = true
			}
		}
	}

	return result
}

func (cmp *ColliderMap) Union(cmp2 ColliderMap) ColliderMap {
	var result ColliderMap

	if len(cmp.data) >= len(cmp2.data) {
		result = cmp.DeepCopy()

		for key, value := range cmp2.data {
			result.data[key] = value
		}
	} else {
		result = cmp2.DeepCopy()

		for key, value := range cmp.data {
			result.data[key] = value
		}
	}

	return result
}
