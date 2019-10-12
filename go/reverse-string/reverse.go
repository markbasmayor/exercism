package reverse

// Reverse reverses a given string
func Reverse(str string) (ret string) {
	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {
		ret += string(runes[i])
	}
	return
}
