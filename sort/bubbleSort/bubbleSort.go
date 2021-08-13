package bubbleSort

// 冒泡排序
/*
 如果前一个比自己小， 即交换位置
 假设 i > j
 if data[i] < data[j]{
 	temp = data[i]
    data[i] = data[j]
    data[j] = temp
}

*/

func BubbleSort(data []int) []int {
	for j := 0; j < len(data); j++ {
		// 无交换时提前退出
		var flag bool
		for i := 1; i < len(data); i++ {
			if data[i-1] > data[i] {
				temp := data[i]
				data[i] = data[i-1]
				data[i-1] = temp
				flag = true
			}
		}
		if !flag{
			break
		}
	}
	return data
}
