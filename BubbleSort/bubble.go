package main

import "fmt"

func main() {
  s := []int{1, 5, 9, 3, 2, 8, 7}
  BubbleSort(s, len(s))
  fmt.Println(s)
}

func BubbleSort(s []int, n int) {
  for i := n - 1; i > 0; i-- {
    for j := 0; j < i; j++ {
      if s[j] > s[j + 1] {
        temp := s[j]
        s[j] = s[j + 1]
        s[j + 1] = temp
      }
    }
  }
}
