package decryption

import "strings"

// Decrypt the message with rot13
func Decrypt_rot13(s string) string {

	s = strings.TrimSpace(s)

	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z'{
			runes[i] = 'a' + (r - 'a' + 13) %26
		}else if r >= 'A' && r <= 'Z'{
			runes[i] = 'A' + (r - 'A' + 13) %26
		}
	}
	return string(runes)
}