package piscine

func SortWordArr(a []string) {
	toString := ""
	answer := ""
	for i := 0; i < len(a); i++ {
		toString += a[i]
	}
	arrByte := []byte(toString)
	for i := 0; i < len(arrByte)-1; i++ {
		for j := i + 1; j < len(arrByte); j++ {
			if arrByte[i] > arrByte[j] {
				arrByte[i], arrByte[j] = arrByte[j], arrByte[i]
			}
		}
	}
	for i := 0; i < len(arrByte); i++ {
		if arrByte[i] >= '0' && arrByte[i] <= '9' {
			answer += string(arrByte[i])
		}
	}
	for i := 0; i < len(arrByte); i++ {
		if arrByte[i] >= 'A' && arrByte[i] <= 'Z' {
			answer += string(arrByte[i])
		}
	}
	for i := 0; i < len(arrByte); i++ {
		if arrByte[i] >= 'a' && arrByte[i] <= 'z' {
			answer += string(arrByte[i])
		}
	}
	endAnsw := []string{}
	for i := 0; i < len(answer); i++ {
		endAnsw = append(endAnsw, string(answer[i]))
	}
	for i := 0; i < len(a) && i < len(answer); i++ {
		a[i] = endAnsw[i]
	}
}
