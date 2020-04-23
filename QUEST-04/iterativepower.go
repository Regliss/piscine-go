package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else if power < 0 {
		return 0
	} else if nb <= 2147483647 && power <= 2147483647 {
		result := 1
		for i := 0; i <= power-1; i++ {
			result *= nb
		}
		return result
	}
	return 0
}
