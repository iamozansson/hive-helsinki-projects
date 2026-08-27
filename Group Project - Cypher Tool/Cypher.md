# Optional group task 10 / 10

3

## Cypher Tool

**Learning Objectives:**

* Students will develop problem-solving skills as they design a command-line program to perform encryption and decryption. They will need to think logically to handle various scenarios and user inputs.
* Students will gain experience in validating user inputs. This is an essential skill in software development to ensure that the program can handle unexpected user behavior.
* By implementing simple encryption algorithms like ROT13 or message reversal, students will gain experience in translating algorithms into code and understanding how these algorithms work.
* This task provides an opportunity for students to practice creating an interactive command-line interface. They will learn how to prompt users for input, provide options, and present results.
* Students will have the opportunity to collaborate with team members. This involves discussing the program structure, defining responsibilities, and integrating individual contributions into a cohesive solution.

### Instructions

You are tasked with creating a command-line tool that allows users to encrypt and decrypt messages using various encryption techniques.

The program has to do the following:

1. Greet the user
2. Allow the user to select the operation (encrypt or decrypt)
3. Allow the user to select the encryption type
4. Allow the user to input the message to encrypt/decrypt
5. Output the result of the operation

If user input is invalid, the program should keep prompting the user to input again, until valid input is provided. Before validating the input, it has to be trimmed (remove whitespaces from the beginning and the end of the input).

When encrypting or decrypting, ensure that any non-alphabet characters in the message are left unchanged.

The tool has to support two encryptions:

6. `rot13`, similar to `ShiftBy` you created for a single rune.
7. `reverse`, *reverse alphabet*, similar to `ReverseAlphabetValue` you created for a single rune.
8. One more encryption of your choice.

You also have to make a brief description in [**Markdown**](https://en.wikipedia.org/wiki/Markdown), that at least:

* Explains what the tool does
* Explains the usage of the tool, with example
* Explains the cyphers used

#### Allowed packages

For this task you're allowed to use:

* `bufio`
* `fmt`
* `os`
* `strconv`
* `strings`

#### Expected functions

```go
// Main logic, envoking other functions
func main() {}

// Get the input data required for the operation
func getInput() (toEncrypt bool, encoding string, message string) {}

// Encrypt the message with rot13
func encrypt_rot13(s string) string {}

// Encrypt the message with reverse
func encrypt_reverse(s string) string {}

// Decrypt the message with rot13
func decrypt_rot13(s string) string {}

// Decrypt the message with reverse
func decrypt_reverse(s string) string {}
```

#### Usage

```text
$ ./cyphertool

Welcome to the Cypher Tool!

Select operation (1/2):
1. Encrypt.
2. Decrypt.
2

Select cypher (1/2):
1. ROT13.
2. Reverse.
2

Enter the message:
zb

Decrypted message using reverse:
ay
```
