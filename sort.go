package main

import (
	"fmt"
	"math/rand"
	"time"
)

var   dataArray []int64
func main()  {

arr := creatArray()
//BubbleSort(arr)
	SelectSort(arr)

	
}
func creatArray() []int64 {
	// 植入种子
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		num := rand.Int63n(100)
		dataArray = append(dataArray, num)

	}
	for i,v := range dataArray{
		fmt.Printf("i = %d value = %d \n",i,v)
	}
	fmt.Println("+++++++++++++++++++")
	return dataArray
	
}

func BubbleSort(arr []int64)  {
	 length := len(arr)
	if length< 0 {
		fmt.Println("error! no data")
		return
	}
	for i := 0; i < length; i++ {
		flag := true //若为true，则表示此次循环没有进行交换，也就是待排序列已经有序
		for j:=0; j < length -i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
			//for i,v := range arr{
			//	fmt.Printf("i = %d value = %d \n ",i,v)
			//
			//}
			//fmt.Println("--------------------")
			
		}
		//for i,v := range arr{
		//	fmt.Printf("i = %d value = %d \n ",i,v)
		//
		//}
		//fmt.Println("FFFFFFFFFFFFFFFF")

		fmt.Println(i)
		if flag {
			for i,v := range arr{
				fmt.Printf("i = %d value = %d \n ",i,v)

			}
			break
		}
		
	}

	
}


func SelectSort(arr []int64) {
	length := len(arr)
	for i := 0; i < length -1; i++ {
		min := i //每一趟的开始把首元素的下标作为最小元素的下标
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}

	for i,v := range arr{
		fmt.Printf("i = %d value = %d \n ",i,v)

	}

}

func ShellSort(arr []int64) {
	length := len(arr)
	//增量gap，并逐步缩小增量
	for gap := length / 2; gap > 0; gap /= 2 {
		//从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < length; i++ {
			j := i
			for j-gap >= 0 && arr[j] < arr[j-gap] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap
			}
		}
	}

}

func HeapSort(a []int64) {
	length := len(a)
	if length == 0 {
		return
	}
	//构造初始堆
	for i := length/2 - 1; i >= 0; i-- {
		heapAdjust(a, i, length-1)
	}

	for j := length - 1; j >= 0; j-- {
		a[0], a[j] = a[j], a[0]
		heapAdjust(a, 0, j-1)
	}
}

//调整堆
func heapAdjust(a []int64, start, end int) {
	temp := a[start]

	for k := 2*start + 1; k <= end; k = 2*k + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		//选择出左右孩子较大的下标
		if k < end && a[k] < a[k+1] {
			k++
		}
		//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
		if a[k] > temp {
			a[start] = a[k]
			start = k
		} else {
			break
		}
	}
	a[start] = temp //插入正确的位置

}

func MergeSort(arr []int64) {
	length := len(arr)
	temp := make([]int64, length) //提前开辟一块内存空间存放临时数据
	mSort(arr, 0, length-1, temp)
}

func mSort(arr []int64, left, right int, temp []int64) {
	if left < right {
		mid := (left + right) / 2
		mSort(arr, left, mid, temp)
		mSort(arr, mid+1, right, temp)
		//两边的子序列都是有序的，
		//如果左边的最大的元素比右边最小的元素大才需要合并
		if arr[mid] > arr[mid+1] {
			merge(arr, left, mid, right, temp)
		}
	}
}

func merge(arr []int64, left, mid, right int, temp []int64) {
	i := left
	j := mid + 1
	t := 0 //临时slice的指针
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			i++
		} else {
			temp[t] = arr[j]
			j++
		}
		t++
	}
	//将左序列剩余元素填充进temp中
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	//将右序列剩余元素填充进temp中
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}
	t = 0
	//将temp中的元素全部拷贝到原数组中
	for left <= right {
		arr[left] = temp[t]
		left++
		t++
	}

}

func QuickSort1(arr []int64, left, right int) {
	if left < right {
		i := arrAdjust(arr, left, right)
		QuickSort1(arr, left, i-1)  //调整左边的序列
		QuickSort1(arr, i+1, right) //调整右边的序列
	}
}

//返回调整后基准数的位置
func arrAdjust(arr []int64, left, right int) int {
	x := arr[left] //基准
	i, j := left, right
	for i < j {
		//从右向左找小于x的数
		for i < j && arr[j] >= x {
			j--
		}
		//从左向右找大于x的数
		for i < j && arr[i] <= x {
			i++
		}
		//交换
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// i=j结束扫描
	// 基准数归位，现在左边的序列小于基准数，右边的序列大于基准数
	arr[left], arr[i] = arr[i], arr[left]
	return i

}
func QuickSort2(arr []int64, left, right int) {
	if left < right {
		i := arrAdjust2(arr, left, right)
		QuickSort2(arr, left, i-1)  //调整左边的序列
		QuickSort2(arr, i+1, right) //调整右边的序列
	}
}

//返回调整后基准数的位置
func arrAdjust2(arr []int64, left, right int) int {
	mid := (left + right) / 2
	arr[left], arr[mid] = arr[mid], arr[left] //可以选择中间的数作为基准
	x := arr[left]                            //基准
	i, j := left, right
	for i < j {
		//从右向左找小于x的数
		for i < j && arr[j] >= x {
			j--
		}
		//从左向右找大于x的数
		for i < j && arr[i] <= x {
			i++
		}
		//交换
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// i=j结束扫描
	// 基准数归位，现在左边的序列小于基准数，右边的序列大于基准数
	arr[left], arr[i] = arr[i], arr[left]
	return i

}