package main
import (
"fmt"
)

func main2() {
  var str string
  // 赋值
  str = "123"
  n := len(str)
  fmt.Printf("str: [%s] and his length is %d\n", str, n)
  ch := str[0]
  fmt.Printf("str first char: [%s] [%c]\n", ch, ch)

  str = "中国"
  for i := 0; i < n; i++ {
    ch := str[i]
    fmt.Println(i, ch)
  }

  fmt.Println("-------------")
  for i, ch := range str {
    fmt.Println(i, ch)
  }

  array := [5]int{1,2,3,4,5}
  Modify(array)
  fmt.Println("In main(), array values:", array)

  Slice()
  Vars(1,2,3,4,5,6)
}

func Modify(array [5]int) {
  array[0] = 10
  fmt.Println("In Modify(), array values:", array)
}

func Slice() {
  // 给予数组创建切片
  array := [10]int{1,2,3,4,5,6,7,8,9,10} // 给定长度的是数组
  slice := array[:5]
  fmt.Println("Elements of array:")
  for _, v := range array {
    fmt.Print(v, " ")
  }

  fmt.Println("\nElement of slice:")
  for _, v := range slice {
    fmt.Print(v, " ")
  }

  fmt.Println()

  // 通过 make 创建 slice
  slice = make([]int, 5)
  slice2 := make([]int, 5, 10)
  slice3 := []int{1,2,3,4,5,6,7} // 不给定长度的是数组分片

  fmt.Println("slice", slice, "; cap:", cap(slice), "len:", len(slice))
  fmt.Println("slice2", slice2, "; cap:", cap(slice2), "len:", len(slice2))
  fmt.Println("slice3", slice3, "; cap:", cap(slice3), "len:", len(slice3))
}


func Vars(args ...int) {
  for _, v := range args {
    fmt.Println(v)
  }
}











