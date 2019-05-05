package leetcode

func canVisitAllRooms(rooms [][]int) bool {
	opened := make([]bool, len(rooms))
	opennum := 1
	keys := []int{}
	opened[0] = true
	for i := range rooms[0] {
		keys = append(keys, rooms[0][i])
	}
	for len(keys) > 0 {
		key := keys[0]
		keys = keys[1:]
		if key < len(rooms) && !opened[key] {
			opened[key] = true
			opennum++
			for _, new_key := range rooms[key] {
				if !opened[new_key] {
					keys = append(keys, new_key)
				}
			}
		}
	}
	return opennum == len(rooms)
}
