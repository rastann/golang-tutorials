package algorithms

func partition(arr[] int, low, high int) int {
	pivot := arr[high]
	i := low-1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i+1
}

func Sort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		Sort(arr, low, p-1)
		Sort(arr, p+1, high)
	}
}