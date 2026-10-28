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

	if args[1] == "*" || args[1] == "+" || args[1] == "-" || args[1] == "/" || args[1] == "%" {
		firstNum := 0
		signFirst := 1
		First := args[0]

		i := 0

		if len(First) > 0 && First[0] == '-' {
			i = 1
			signFirst = -1
		} else if len(First) > 0 && First[0] == '+' {
			i = 1
		}
		if i == len(First) {
			return
		}

		for ; i < len(First); i++ {
			if First[i] >= '0' && First[i] <= '9' {
				digit := int(First[i] - '0')

				if firstNum > (9223372036854775807-digit)/10 {
					return
				}
				firstNum = firstNum*10 + digit
			} else {
				return
			}
		}

		if signFirst == 1 && firstNum < 0 {
			return
		}
		if signFirst == -1 && firstNum-1 > 9223372036854775807 {
			return
		}
		firstNum = firstNum * signFirst

		Second := args[2]
		secondNum := 0
		signSecond := 1
		j := 0

		if len(Second) > 0 && Second[0] == '-' {
			j = 1
			signSecond = -1
		} else if len(Second) > 0 && Second[0] == '+' {
			j = 1
		}

		if j == len(Second) {
			return
		}

		for ; j < len(Second); j++ {
			if Second[j] >= '0' && Second[j] <= '9' {
				digit := int(Second[j] - '0')

				if secondNum > (9223372036854775807-digit)/10 {
					return
				}
				secondNum = secondNum*10 + digit
			} else {
				return
			}
		}

		if signSecond == 1 && secondNum < 0 {
			return
		}
		if signSecond == -1 && secondNum-1 > 9223372036854775807 {
			return
		}
		secondNum = secondNum * signSecond

		result := 0

		if args[1] == "+" {
			if (secondNum > 0 && firstNum > 9223372036854775807-secondNum) ||
				(secondNum < 0 && firstNum < -9223372036854775808-secondNum) {
				return
			}
			result = firstNum + secondNum
			printInt(result)
		} else if args[1] == "-" {
			if (secondNum < 0 && firstNum > 9223372036854775807+secondNum) ||
				(secondNum > 0 && firstNum < -9223372036854775808+secondNum) {
				return
			}
			result = firstNum - secondNum
			printInt(result)
		} else if args[1] == "*" {
			if firstNum != 0 && secondNum != 0 {
				result = firstNum * secondNum
				if result/firstNum != secondNum {
					return
				}
			}
			printInt(result)
		} else if args[1] == "/" {
			if secondNum == 0 {
				os.Stdout.WriteString("No division by 0\n")
				return
			}
			if firstNum == -9223372036854775808 && secondNum == -1 {
				return
			}
			result = firstNum / secondNum
			printInt(result)
		} else if args[1] == "%" {
			if secondNum == 0 {
				os.Stdout.WriteString("No modulo by 0\n")
				return
			}
			if firstNum == -9223372036854775808 && secondNum == -1 {
				return
			}
			result = firstNum % secondNum
			printInt(result)
		} else {
			return
		}
	}
}
