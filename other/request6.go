package other

import (
	"fmt"
	"sort"
)

func Run6() {
	fmt.Println("\n	6.")
	arr := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	fmt.Println("Mảng cha: ", arr)

	sort.Ints(arr)
	fmt.Println("Mảng cha đã sắp xếp: ", arr)

	childrenarr := arr[1:7]
	fmt.Println("Mảng con: ", childrenarr)

	if len(arr) > 15 {
		copyarr := arr[1:15]
		fmt.Println(copyarr)
	} else {
		fmt.Println("Lỗi khởi tạo vì độ dài mảng dc copy: ", len(arr))
	}
}
