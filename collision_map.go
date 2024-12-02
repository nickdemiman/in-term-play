package intermplay

// type (
// 	CollisionMap struct {
// 		data map[Vector2]bool
// 	}
// )

// func NewCollisionMap() CollisionMap {
// 	mp := CollisionMap{
// 		data: make(map[Vector2]bool),
// 	}

// 	return mp
// }

// func (cmp *CollisionMap) SetCell(x, y int) {
// 	v := NewVector2(x, y)
// 	cmp.data[v] = true
// }

// func (cmp *CollisionMap) DeepCopy() CollisionMap {
// 	result := NewCollisionMap()

// 	for key, value := range cmp.data {
// 		result.data[key] = value
// 	}

// 	return result
// }

// func (cmp *CollisionMap) IsEmpty() bool {
// 	return len(cmp.data) == 0
// }

// func (cmp *CollisionMap) Intersect(cmp2 CollisionMap) CollisionMap {
// 	result := NewCollisionMap()

// 	if len(cmp2.data) >= len(cmp.data) {
// 		for obj := range cmp.data {
// 			_, ok := cmp2.data[obj]

// 			if ok {
// 				result.data[obj] = true
// 			}
// 		}
// 	} else {
// 		for obj := range cmp2.data {
// 			_, ok := cmp.data[obj]

// 			if ok {
// 				result.data[obj] = true
// 			}
// 		}
// 	}

// 	return result
// }

// func (cmp *CollisionMap) Union(cmp2 CollisionMap) CollisionMap {
// 	var result CollisionMap

// 	if len(cmp.data) >= len(cmp2.data) {
// 		result = cmp.DeepCopy()

// 		for key, value := range cmp2.data {
// 			result.data[key] = value
// 		}
// 	} else {
// 		result = cmp2.DeepCopy()

// 		for key, value := range cmp.data {
// 			result.data[key] = value
// 		}
// 	}

// 	return result
// }
