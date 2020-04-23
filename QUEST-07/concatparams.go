package piscine

func ConcatParams(args []string) string {
	result := ""
	len := 0
	for range args {
		len++
	}
	for i := 0; i < len-1; i++ {
		result = result + args[i] + "\n"
	}
	return result + args[len-1]
}
