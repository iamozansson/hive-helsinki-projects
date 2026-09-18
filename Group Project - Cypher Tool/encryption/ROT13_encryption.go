package encryption

import "strings"

func Encrypt_rot13(s string) string {

	s = strings.TrimSpace(s)
	secretcode := ""

	for _, value := range s {

		if value >= 'A' && value <= 'Z' {
			value = 'A' + (value-'A'+13)%26
		} else if value >= 'a' && value <= 'z' {
			value = 'a' + (value-'a'+13)%26
		}
		secretcode += string(value)
	}
	return secretcode
}
