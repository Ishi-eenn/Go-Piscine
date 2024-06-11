package piscine

import "ft"

func putComb(a int, b int) {
  ft.PrintRune(rune(a + '0'))
  ft.PrintRune(rune(b + '0'))
}

func putComb2(a int, b int) {
  putComb(a / 10, a % 10);
  ft.PrintRune(' ');
  putComb(b / 10, b % 10);
  if a != 98 {
    ft.PrintRune(',');
    ft.PrintRune(' ');
  }
}

func PrintComb2() {
  for i:= 0; i <= 98; i++ {
    for j:= i + 1; j <= 99; j++ {
      putComb2(i, j)
    }
  }
  ft.PrintRune('\n')
}
