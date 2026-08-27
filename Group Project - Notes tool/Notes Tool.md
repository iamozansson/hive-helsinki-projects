# Notes Tool

A command-line tool for managing short, single-line notes.

## Learning Objectives

* Practice problem-solving by handling different user inputs and managing notes.
* Practice input validation and handling invalid arguments or input.
* Create an interactive command-line interface with a simple menu.
* Practice working with files to store and persist data.
* Collaborate with team members and integrate different parts of a project.

## Features

The tool allows users to:

1. Display notes from a collection.
2. Add a new note.
3. Remove an existing note.
4. Exit the program.

Each collection is stored separately, and notes persist between program runs.

## Usage

The tool requires exactly one argument: the name of the collection.

```bash
./notestool coding_ideas
```

If the collection does not exist, it will be created automatically. If it already exists, its notes will be loaded.

### Help

If no argument is provided, more than one argument is provided, or the argument is `help`, the application displays a help message.

```bash
./notestool
Usage: ./notestool [TAG]
```

## Data Storage

Each collection is stored as a plain text file with the same name as the collection.

For example:

```bash
./notestool coding_ideas
```

creates or loads a file named:

```text
coding_ideas
```

Notes are stored in separate rows in the file. This allows the notes to persist between different runs of the application.

## Example

```text
$ ./notestool testtag

Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2

Enter the note text:
note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two
003 - note three

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
3

Enter the number of note to remove or 0 to cancel:
3

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1

Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
```

## Allowed Packages

The following Go packages are allowed:

* `bufio`
* `fmt`
* `os`
* `strconv`
* `strings`
