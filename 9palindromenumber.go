func isPalindrome(x int) bool {
	palindromo := strconv.Itoa(x)
	runeArray := []rune(palindromo)
	var impar bool
	longitud := len(runeArray)
	if longitud%2 != 0 {
		impar = true
	}
	for i, v := range runeArray {
		if impar && v == runeArray[longitud-1-i] {
			continue
		}
		if v != runeArray[longitud-1-i] {
			return false
		}
	}
	return true
}
