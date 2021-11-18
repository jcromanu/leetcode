func romanToInt(s string) int {
	runas := []rune(s)
	var suma int = 0
	for i := 0; i < len(runas); i++ {
		switch runas[i] {
		case 'I':
			if i+1 < len(runas) {
				if runas[i+1] == 'V' {
					suma = suma + 4
					i = i + 1
					continue
				} else if runas[i+1] == 'X' {
					suma = suma + 9
					i = i + 1
					continue
				}
			}
			suma = suma + 1
		case 'V':
			suma = suma + 5
		case 'X':
			if i+1 < len(runas) {
				if runas[i+1] == 'L' {
					suma = suma + 40
					i = i + 1
					continue
				} else if runas[i+1] == 'C' {
					suma = suma + 90
					i = i + 1
					continue
				}
			}
			suma = suma + 10

		case 'L':
			suma = suma + 50
		case 'C':
			if i+1 < len(runas) {
				if runas[i+1] == 'D' {
					suma = suma + 400
					i = i + 1
					continue
				} else if runas[i+1] == 'M' {
					suma = suma + 900
					i = i + 1
					continue
				}
			}
			suma = suma + 100
		case 'D':
			suma = suma + 500
		case 'M':
			suma = suma + 1000
		}
	}
	return suma
} 
