package mytest

import "fmt"

//Sum xxx
func Sum(i,j int) int{
	return i+j
}

//Div xxx
func Div(i,j int) (int, error) {
	if j == 0 {
		return  0, fmt.Errorf("j is zero")
	}
	return i/j, nil
}


//BubbleSort  xxx
func BubbleSort(items []int){
	var (
		n = len(items)
		swapped = true
	)

	for swapped {
		swapped = false
		for i:= 0;i <n-1;i++{
			if items[i] > items[i+1]{
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		n = n -1
	}
}

//QuickSort xxx
func QuickSort(items []int){
	sort(items, 0, len(items) -1)
}

func sort(items []int, left int, right int){
	if left < right {
		var i ,j = left, right
		var key = items[i]
		for i < j {
			for i<j && items[j] >= key{
				j--
			}
			items[i] = items[j]
			for i<j && items[i] <= key{
				i++
			}
			items[j] = items[i]
		}
		items[i] = key
		sort(items, left, i - 1)
		sort(items, i+1, right)
	}
}

func Malloc()[]int{
	return make([]int , 1024)
}
