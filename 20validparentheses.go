

func isValid(s string) bool {
	var top int = -1
	var stack []rune
	//leer string
	length := len(s)

	//crear pila de tipo rune y del tamaño del string
	stack = make([]rune, length)

	//Iterar dentro del string
	for _, v := range s {
		//validar si el primer valor es parentesis que cierra terminar programa
		if top == -1 && (v == '}' || v == ']' || v == ')') {
			return false
		} else {
			//si el caracter es ( [ { hacemos push
			if v == '{' || v == '[' || v == '(' {
				top = top + 1
				stack[top] = v
			} else {
				switch v {
				case '}':
					if stack[top] == '{' {
						top = top - 1
					} else {
						top = top + 1
						stack[top] = v
					}

				case ')':
					if stack[top] == '(' {
						top = top - 1
					} else {
						top = top + 1
						stack[top] = v
					}
				case ']':
					if stack[top] == '[' {
						top = top - 1
					} else {
						top = top + 1
						stack[top] = v
					}
				}
			}
		}
		// si el primer valor no es parentesis que cierra y la pila esta vacia insertamos los valores que abren
	}
	if top != -1 {
		return false
	} else {
		return true
	}
}
	