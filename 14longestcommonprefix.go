func longestCommonPrefix(strs []string) string {
	var caracteresComunes []rune
	var final bool = false
	var menor string
	var maximo int = 200
	var length int
	for _, v := range strs {
		length = utf8.RuneCountInString(v)
		if length == 0 {
			return ""
		}
		if length <= maximo {
			menor = v
			maximo = length
		}
	}
	fmt.Println(menor)
	for i, _ := range menor {
		for _, stringInicial := range strs {
			if rune(stringInicial[i]) == rune(menor[i]) {
				continue
			} else {
				final = true
			}
		}
		if final {
			break
		}
		caracteresComunes = append(caracteresComunes, rune(menor[i]))

	}
	if len(caracteresComunes) > 0 {
		return string(caracteresComunes)
	} else {
		return ""
	}
}
