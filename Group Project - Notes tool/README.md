# Notes Tool

Notes Tool is a simple command-line application written in Go for managing short, single-line notes.

The tool allows users to:

* Show notes
* Add a note
* Delete a note
* Exit the application

## Usage

The tool requires exactly one argument, which is the name of the collection.

```bash
./notestool coding_ideas
```

If no argument is provided, or more than one argument is provided, the tool displays the usage message:

```text
Usage: ./notestool [TAG]
```

After starting the tool, the user can select an operation from the menu:

```text
Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
```

### Show notes

Select `1` to display the notes in the collection.

```text
Notes:
001 - note one
002 - note two
```

### Add a note

Select `2` and enter the note text.

```text
Enter the note text:
note three
```

The note is then added to the collection.

### Delete a note

Select `3` and enter the number of the note to remove.

```text
Enter the number of note to remove or 0 to cancel:
3
```

Entering `0` cancels the delete operation.

### Exit

Select `4` to exit the application.

## Data Storage

Each collection is stored in a separate plain text file. The file name is the same as the collection name.

For example:

```bash
./notestool coding_ideas
```

uses a file named:

```text
coding_ideas
```

Each note is stored on a separate line:

```text
learn Docker
practice Linux
study Go
```

When the tool is started again with the same collection name, the existing notes are loaded from the file.

## Contributors

- Ville Myllyniemi
- Ozan Arikan
- Malaika Ahmed
