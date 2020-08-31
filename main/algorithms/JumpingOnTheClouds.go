package algorithms

func jumping_on_the_clouds(arr []int) int {
	steps := 0
	for i := 0; i < len(arr)-1; {
		if(i+1 == len(arr)-1) {
			i = i+1
		} else if arr[i+1] == 1 {
			i = i+2
		} else if arr[i+2] == 1 {
			i = i+1
		} else {
			i = i+2
		}
		steps += 1
	}
	return steps
}
