package piscine

import "ft"

func Printdigits() {
  for c := '0'; c <= '9'; c++ {
    ft.PrintRune(c)
  }
  ft.PrintRune('\n')
}
