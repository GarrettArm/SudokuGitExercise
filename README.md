# SudokuGitExercise

A tool for teaching Git.

## How a learning session works

Each person in the group is at a computer with git installed.  We'll only be using the Git program, the File Explorer, and a WebBrowser.

In the process, we'll learn to:

- grab a shared repo
- create a branch to work in
- share our work back to the main repo
- compare our branch with other branches
- merge branches

We do this by playing a game of Sudoku as a team.  Each person will be pushing a few squares at a time to the shared repo, and over time we'll combine everyone's work.

Solving a sudoku puzzle is a metaphor for writing code.  In both, you write in text files a solution to the puzzle.  Both you find yourself trying various approaches, and sometimes you take dead-ends.

Git lets you create a bread crumbs path of all the steps and missteps you took while writing a solution.  It also communicates all those paths so that others may follow, branch off in their own direction, and sometimes share back to you the path they found.

In this game, the file "board.txt" is analogous to code.  For code, you may use VSCode to edit your code files.  In a direct analogy, we are using this sudoku webapp to edit the file "board.txt".  This one file represents the project's codebase.

It allows collaboration by sharing a history of text files.  The important text file in this project is "board.txt".

## Setting it up

Pick a folder on your computer where you'd like to put this repo.  Then, Git Clone this repo

  - via the git app
  - or via `git clone https://github.com/GarrettArm/SudokuGitExercise`

Checkout a new branch of the repo, for your unique work.  Give it a memorable name, like your name.

  - via the git app
  - or via `git checkout -b YourName`

Git Push your new branch to the repo.

  - via the git app
  - or via `git push -set-upstream origin YourName`

Now your version of the repo is linked to the shared repo.  We can all share our unique work, without affecting anyone else's work.

Open this folder and run the program: sudoku-windows.exe or sudoku-mac-intel.

Open your webbrowser to [http://localhost:8100](localhost:8100)






## [Admin only] Rebuilding the executables

with Go installed on your computer:

```
git clone https://github.com/GarrettArm/SudokuGitExercise
cd SudokuGitExercise
```

```
GOOS=darwin GOARCH=arm64 go build -o sudoku-mac-M1 .
GOOS=windows GOARCH=amd64 go build -o sudoku-windows.exe .
GOOS=linux GOARCH=amd64 go build -o sudoku-linux .
```

