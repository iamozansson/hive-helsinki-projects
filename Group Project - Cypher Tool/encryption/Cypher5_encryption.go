package encryption

import "strings"

func Encrypt_cypher5(s string) string {

	s = strings.TrimSpace(s)
	secretcode := ""

	for _, value := range s {

		if value >= 'A' && value <= 'Z' {
			value = ((value-'A')+5)%26 + 'A'
		} else if value >= 'a' && value <= 'z' {
			value = ((value-'a')+5)%26 + 'a'
		}
		secretcode += string(value)
	}
	return secretcode
}
