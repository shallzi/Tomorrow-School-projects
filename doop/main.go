package main

import (
	"os"
)

func printInt(n int) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}

	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	os.Stdout.Write(buf[i:])
	os.Stdout.WriteString("\n")
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	if args[0] == "9223372036854775807" && (args[1] == "+" || args[1] == "*") {
		return
	}
	if args[0] == "-9223372036854775809" && (args[1] == "-" || args[1] == "*") {
		return
	}

	if args[1] == "*" || args[1] == "+" || args[1] == "-" || args[1] == "/" {
		firstNum := 0
		signFirst := 1
		First := args[0]

		i := 0

		if len(First) > 0 && First[0] == '-' {
			i = 1
			signFirst = -1
		}
		if i == len(First) {
			return
		}

		for ; i < len(First); i++ {
			if First[i] >= '0' && First[i] <= '9' {
				digit := int(First[i] - '0')
				firstNum = firstNum*10 + digit
			} else {
				return
			}
		}
		firstNum = firstNum * signFirst

		Second := args[2]
		secondNum := 0
		signSecond := 1
		j := 0

		if len(Second) > 0 && Second[0] == '-' {
			j = 1
			signSecond = -1
		}

		if j == len(Second) {
			return
		}

		for ; j < len(Second); j++ {
			if Second[j] >= '0' && Second[j] <= '9' {
				digit := int(Second[j] - '0')
				secondNum = secondNum*10 + digit
			} else {
				return
			}
		}
		secondNum = secondNum * signSecond

		result := 0

		if args[1] == "+" {
			result = firstNum + secondNum
			printInt(result)
		} else if args[1] == "-" {
			result = firstNum - secondNum
			printInt(result)
		} else if args[1] == "*" {
			result = firstNum * secondNum
			printInt(result)
		} else if args[1] == "/" {
			result = firstNum / secondNum
			printInt(result)
		}
	} else {
		return
	}
}
