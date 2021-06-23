package util

func Filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for i := range arr {
	  if cond(arr[i]) {
		 result = append(result, arr[i])
	  }
	}
	return result
}