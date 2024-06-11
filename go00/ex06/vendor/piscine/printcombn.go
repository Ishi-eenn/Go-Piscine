package piscine

import "ft"

func putNumber(array []int, n int) {
  for i := 0; i < n; i++ {
    ft.PrintRune(rune(array[i]) + '0')
  }
  if array[0] != 10 - n {
    ft.PrintRune(',')
    ft.PrintRune(' ')
  }
}

func calc(array []int, n int, i int) {
  if i == 0 {
     array[i] = 0
  } else {
    array[i] = array[i - 1] + 1
  }
  for array[i] < 11 - n + i {
    if i == n - 1 {
      putNumber(array, n)
    } else {
      calc(array, n, i + 1)
    }
    array[i]++
  }
}

func PrintCombN(n int) {
  array := make([]int, n)

  calc(array, n, 0)
  ft.PrintRune('\n')
}
