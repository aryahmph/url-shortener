package base62

import "strings"

type Base62Manager struct {
}

func NewBase62Manager() *Base62Manager {
	return &Base62Manager{}
}

func (Base62Manager) Hash(number uint64) string {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if number == 0 {
		return string(chars[0])
	}
	var sb strings.Builder
	for number > 0 {
		sb.WriteRune(rune(chars[number%62]))
		number /= 62
	}
	return sb.String()
}
