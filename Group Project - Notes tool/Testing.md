### Testing

Testing ensures that the software works as expected by validating its features against the requirements. It helps catch bugs early, improves reliability, and maintains quality.

## How to Test

1. Clone the repository, then build and run the submitted code.
2. Agree on how testing responsibilities are divided between reviewers.
3. Test functionality and check compliance with the requirements.
4. Provide feedback in the group chat and request fixes if necessary.
5. Clearly state which changes are mandatory and which are optional.
6. Repeat the testing cycle after requested changes are made.

## Mandatory

### Description File

Is there a description file with a brief explanation of the tool in Markdown?

Required sections:

- What does the tool do?
- Tool usage with examples.
- Explanation of how the data is stored.

### Command-Line Arguments

Is an error message displayed if the number of arguments is not 1?

### Command-Line Execution

Can the tool be run from the command line without errors?

### User Greeting

Does the tool greet the user?

### Collection Persistence

Can collections be created and do they persist between uses?

Test by:

1. Opening or creating a collection.
2. Adding notes.
3. Exiting the application.
4. Running the tool again.
5. Opening the same collection and checking that the notes are still present.

### Collection Isolation

Do collections work independently?

Open or create multiple collections and verify that manipulating one collection does not affect another.

### Creating and Deleting Notes

Can notes be added and removed correctly?

Test by:

1. Opening or creating a collection.
2. Adding multiple notes.
3. Displaying the notes.
4. Removing notes.
5. Displaying the notes again.
6. Exiting and restarting the application.
7. Checking that the collection was stored correctly.

### Menu

Does the tool provide the user with the required menu options?

The menu should include:

1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

### General Functionality

Can the user perform operations such as adding, removing, and displaying notes in any order and for any amount of time without restarting the tool?

## Extra

### Invalid Input Handling

Does the tool check for invalid input and provide helpful error messages?

Test cases:

- Invalid input in the menu.
- Empty note.
- Non-existing note ID when removing a note.

### Documentation Quality

How well is the description file designed?

Check:

- Is Markdown used effectively?
- Is the description easy to read and understand?
- Are the examples useful and clear?

### Code Structure and Readability

How well is the program structured and how readable is the code?

Consider:

- Are functions and variables named clearly?
- Are important parts of the code explained?
- How did the students structure the code?
- Could the structure or implementation be improved?

Grade:

**0 1 2**

Poor → Great

### User Experience Features

Are there extra features that make using the tool more comfortable?

Examples:

- Terminal is cleared when moving between menus.
- Colors or text decorations are used.
- Menu navigation uses arrow keys.

### Additional Features

Are there additional features that make the tool more useful?

Examples:

- Password protection and encryption of the data.
- Timestamps on notes.
- Tags, titles, or other extra data on notes.
