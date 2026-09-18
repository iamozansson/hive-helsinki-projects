package decryption
import "strings"

// ReverseDecrypt the message with rot13
func Decrypt_reverse(s string) string {
	s = strings.TrimSpace(s)
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = 'a' + ('z'-r)%26
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = 'A' + ('Z'-r)%26
		}
	}
	return string(runes)
}
