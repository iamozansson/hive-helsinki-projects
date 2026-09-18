# kood / Wordle

A Go command-line game completed during my Hive Helsinki programming sprint.

## Objective

The goal of this project was to build a command-line Wordle game while practicing Go application structure, user input, file handling, game state management, and data persistence.

## Game Rules

The player tries to guess a hidden five-letter word within six attempts.

After each guess, letters are shown with different colors:

* **Green** — the letter is in the correct position
* **Yellow** — the letter exists in the word but is in the wrong position
* **White** — the letter does not exist in the word

If the player does not guess the word within six attempts, the secret word is revealed.

## Features

* Username-based game sessions
* Five-letter Wordle gameplay
* Six maximum attempts
* Colored feedback using ANSI escape codes
* Remaining letter tracking
* Sorted remaining letters
* Persistent game statistics
* CSV-based statistics storage
* Command-line word selection using an index
* Graceful handling of invalid arguments and missing files
* EOF handling with `bufio.Scanner`

## Data Management

The application reads the word list from:

```text
wordle-words.txt
```

The word list is kept outside the repository and is added to `.gitignore`.

Game statistics are stored in:

```text
stats.csv
```

The statistics file persists between games and is also excluded from the repository.

Each completed game records:

```text
username,secret word,number of attempts,win or loss
```

## Usage

The game is started from the command line by providing the index of the secret word:

```bash
go run . 10
```

The application then asks for a username and starts the game.

After the game, the player can choose to view their statistics, including:

* Games played
* Games won
* Average attempts per game

## Implementation

The application uses `bufio.Scanner` for reading user input from standard input.

The game logic handles:

* Guess validation
* Letter feedback
* Attempt tracking
* Remaining letter management
* ANSI terminal colors
* Word list loading
* Statistics persistence

The project is organized as a Go command-line application with `main.go` at the repository root.

## What I Practiced

* Go CLI applications
* `bufio.Scanner`
* Command-line arguments
* String and rune manipulation
* Slices and data structures
* Game state management
* File I/O
* CSV data handling
* Persistent state
* ANSI escape codes
* Input validation
* Error handling
