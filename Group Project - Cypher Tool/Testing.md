### Testing

Ensures that software works as expected by validating features against requirements. It helps catch bugs early, improves reliability, and maintains high-quality standards in development.

### How to do testing?

1. Clone the repository, then build and run the submitted code.
2. Agree on your teamwork: how do you divide testing between reviewers?
3. Test functionality and check compliance with the requirements.
4. Provide feedback in the group chat and request fixes if necessary.
5. Clearly state what changes are mandatory and what are optional fixes.
6. Repeat the testing cycle after submitters make the requested changes for as many times as is needed.

## Mandatory

### Is there a description file with brief explanation of the tool in Markdown with required sections?

Required sections are:

* What does the tool do?
* Tool usage with examples.
* Explanation of the cyphers used.

### Is an additional encryption algorithm added?

Additionally could ask to explain what the algorithm does.

### Can the tool be run from the command line with no errors?

### Does the tool greet the user?

See example usage.

### Does the tool provide the user with a menu with required options (encrypt and decrypt)?

### Can the user select the encryption and decryption algorithms in the corresponding submenus?

ROT13, reverse, and an additional algorithm.

### Can the user select the operation, enter the phrase to perform the operation on, and see the result of the operation?

### Does the ROT13 encryption and decryption work?

Encryption and decryption performs the same thing; the functions are symmetrical.

Examples:

* `kood` ↔ `xbbq`
* `rot13` ↔ `ebg13`

### Does the reverse encryption and decryption work?

Encryption and decryption performs the same thing; the functions are symmetrical.

Examples:

* `kood` ↔ `pllw`
* `r3v a1` ↔ `i3e z1`

### Does the encryption and decryption work with the additional algorithm?

Can be tested by encoding a string and then decoding it.

### Does the program have all the expected functions?

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

## Extra

### Does the tool check for invalid input and does it have helpful error messages for these cases?

Cases that could be tested for:

* Invalid input.
* Empty input.
* No arguments or invalid arguments, if arguments are used for the tool.

### Does the tool work with edge cases?

Cases that could be tested for:

* Non-printable characters.
* Empty string.
* Numbers.
* Special characters, or punctuation.
* Upper and lower case characters.

### How well is the description file designed?

Is MD used in a good way that supports easier understanding of the description?

Is the description easy to read and understand?

Are the examples chosen and shown in a way that help the user?
