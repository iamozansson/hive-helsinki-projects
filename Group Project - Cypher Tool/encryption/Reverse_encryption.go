package encryption

import "strings"

func Encrypt_reverse(s string) string{

        s = strings.TrimSpace(s)
        secretcode := ""

        for _, value := range s {

                if value >= 'A' && value <= 'Z' {
                        value = 'A' + ('Z' - value)%26
                } else if value >= 'a' && value <= 'z' {
                        value = 'a' + ('z' - value)%26
                }
                secretcode += string(value)
        }
        return secretcode
}
