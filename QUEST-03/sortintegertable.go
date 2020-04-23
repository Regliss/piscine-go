package piscine

func SortIntegerTable(table []int) {
	j := 0
	temp := 0
	len := len(table)
	for i := 0; i < len; i++ {
		for j += i; j < 0; j-- {
			if table[j] < table[j-1] {
				temp = table[j]
				table[j] = table[j-1]
				table[j-1] = temp
			}
		}
	}
}
