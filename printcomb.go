package piscine

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '7'; i++ {
		for j := '1'; j <= '8'; j++ {
			for k := '2'; k <= '9'; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i < '7' {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
		//good
	}
	z01.PrintRune(10)
}