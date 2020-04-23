package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	lenArr := 0
	for range tab {
		lenArr++
	}
	x, y := 0, 0
	for i := 0; i < +lenArr-2; i++ {
		for j := i + 1; j <= lenArr-1; j++ {
			if f(tab[i], tab[j]) > 0 {
				x++
			} else if f(tab[i], tab[j]) < 0 {
				y++
			}
		}
	}
	if x != 0 && y != 0 {
		return false
	} else {
		return true
	}
}
