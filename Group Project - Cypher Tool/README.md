# Cypher Tool

Cypher Tool is a simple command-line program written in Go that allows users to encrypt and decrypt messages using different cyphers.

## Features

* Encrypt messages
* Decrypt messages
* ROT13 cypher
* ROT5 cypher
* Reverse Alphabet cypher
* Keeps non-alphabet characters unchanged
* Validates and trims user input

## Installation

Clone the repository:

```
git clone https://gitea.kood.tech/riowibowo/cypher
cd cypher
```

Build the program:

```
go build
```

Run it:

```
./cypher
```

## Usage

When the program starts, select whether you want to encrypt or decrypt a message, then select the cypher and enter your message.

### Example

```
$ ./cypher

Welcome to the Cypher Tool!

Select operation:
1. Encrypt.
2. Decrypt.

Select cypher:
1. ROT13.
2. Reverse.
3. Cypher 5

Enter the message: zb

Result: ay
```

## Cyphers

### ROT13

ROT13 shifts each letter 13 positions in the alphabet.

For example:

```
hello → uryyb
```

Applying ROT13 again decrypts the message:

```
uryyb → hello
```

### Reverse Alphabet

The Reverse Alphabet cypher replaces each letter with its opposite letter in the alphabet.

For example:

```
a → z
b → y
c → x
```

Therefore:

```
hello → svool
```

Applying the cypher again returns the original message.

### ROT5

ROT5 shifts each digit 5 positions forward.

For example:

```
0 → 5
1 → 6
2 → 7
3 → 8
4 → 9
```

After reaching 9, the cypher wraps around:

```
8 → 3
9 → 4
```

When decrypting, the program shifts each digit 5 positions backward.

Encryption:
1 → 6

Decryption:
6 → 1

## Technologies

* Go
* Command-line interface

## Team

* Rio Wibowo
* Nadja Mckenna
* Ozan Arikan
