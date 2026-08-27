# Mandatory 0 / 40

1

# kood / Wordle

## The Situation

In this project you are tasked with creating a command-line game in Go that mimics the popular [**Wordle**](https://www.nytimes.com/games/wordle/index.html).

The user will guess a hidden word, and your program will provide feedback on their guesses.

## Functional Requirements

### Game Rules

The player has to **guess a 5-letter** word and can make up to **six guesses** to try to identify the secret word.

After each guess, the game provides feedback:

* **Green** → letter is in the correct position.
* **Yellow** → letter exists in the word but is in the wrong position.
* **White** → letter does not exist in the word.

If the player fails to guess the word after **six attempts**, the game should reveal the **secret word**.

### User Interaction

The game must be launched from the command line with a **command-line argument** that represents the **index** of the word in the provided word list.

Example:

```bash
go run . 10
```

If the command-line argument is missing or invalid, the program should handle it gracefully without crashing.

After each guess, the game provides:

* Feedback for the guess.
* Remaining number of attempts.
* A sorted list of remaining letters (A-Z) that have not been guessed incorrectly.
* Letters in the feedback should be capitalized.

You must use **ANSI escape codes** for colors, such as:

```text
\u001B[32m
```

After each game, the player is presented with an option to view their stats.

## Data Management

The word list must be read from:

```text
wordle-words.txt
```

This file must be located at the **project root**.

Do **not** push the word list to the repository. Add it to `.gitignore`.

The stats are stored in:

```text
stats.csv
```

The stats file must persist between games.

Each game adds a row containing:

```text
username,secret word,number of attempts,win or loss
```

Do **not** push `stats.csv` to the repository. Add it to `.gitignore`.

If the word list file is missing, the program should handle it without crashing.

## Technical Implementation

`main.go` needs to be located at the **root** of your `koodWordle` repository.

Your Go module name should match the repository name:

```text
koodWordle
```

You **must use `bufio.Scanner`** to read from standard input (`stdin`).

Handle **EOF (Ctrl+D)** gracefully by breaking out of input loops when:

```go
scanner.Scan()
```

returns `false`.

> **NB!** All the specific requirements and messages are not provided in the description. You have to discover them from the provided video and automated test feedback.

## Key Output Formats

### Start screen

```text
Enter your username:
```

### After login

```text
Welcome to Wordle! Guess the 5-letter word.
Enter your guess:
```

### Valid guess screen

```text
Feedback: <guess word in uppercase>
Remaining letters: <included and untested uppercase letters separated with whitespace>
Attempts remaining: <number>
Enter your guess:
```

### After game

```text
Do you want to see your stats? (yes/no):
```

### After stats question — if `yes`

```text
Stats for <username>:
Games played: <number>
Games won: <number>
Average attempts per game: <float number>
Press Enter to exit...
```

## Possible Package Structure

The following structure is an example and is not compulsory:

```text
koodWordle
├── main.go           // Entry point, processes arguments, starts game
├── game/             // Game logic and mechanics
│   └── game.go       // Core game functionality and feedback generation
├── io/               // Input/output operations
│   └── io.go         // File handling for words and statistics
└── model/            // Data structures
    └── user.go       // User entity and statistics tracking
```

## Video Example

Video: kood Wordle demo

https://www.youtube.com/watch?v=x5xSiqtNuQg

## Bonus Functionality

After submitting the initial version, you may implement:

* **Instant replay** → Restart another game without relaunching the application.
* **Random word selection** → Randomly select a word instead of using the command-line index.
* **Multi-Length Word Support** → Allow 5-, 6-, or 7-letter words.

> **NB! Do not submit these functionalities for automated tests.**

## Useful Links

* [**Command-line arguments in Go**](https://gobyexample.com/command-line-arguments)
* [**Reading user input**](https://gosamples.dev/read-user-input/)
* [**File operations**](https://pkg.go.dev/os)
* [**CSV**](https://en.wikipedia.org/wiki/Comma-separated_values)
* [**ANSI escape codes**](https://en.wikipedia.org/wiki/ANSI_escape_code)
