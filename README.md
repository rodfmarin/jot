# Jot
<img src="./resources/writing-svgrepo-com.svg" alt="Jot" width="200" height="200"><br>
---
Jot (not to be confused with unix utility [jot](http://man.freebsd.org/cgi/man.cgi?jot(1))) is a simple command-line tool written in Go for quickly appending notes to your daily note file in [Obsidian](https://obsidian.md/).

## Features
- Append a note to your daily note file with a single command.
- Integrates seamlessly with Obsidian's daily note format (`YYYY-MM-DD.md`).
- If the daily note file exists, your note is appended; if not, the file is created automatically.

## Usage

1. **Set environment Variables**
   
   Before using Jot, set the `OBSIDIAN_VAULT_PATH` environment variable to the full path where your Obsidian daily notes are stored:
   
   ```bash
   export OBSIDIAN_VAULT_PATH=/path/to/your/obsidian/vault
   ```

2. **Add a Note**
   
   Use the `jot` command followed by your note:
   
   ```bash
   jot "Here's a new note that will be appended to my daily note file in Obsidian"
   ```
   
   This will append your note to today's daily note file (format: `YYYY-MM-DD.md`). If the file does not exist, it will be created.

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/rodfmarin/jot.git
   cd jot
   ```
2. Build the binary:
   ```bash
   go build -o jot
   ```
3. Set your OBSIDIAN_VAULT_PATH environment variable:
   ```bash
   export OBSIDIAN_VAULT_PATH=/path/to/your/obsidian/vault
   ```

   Alternatively, you can set this variable permanently in your shell configuration file (e.g., `.bashrc`, `.zshrc`):
   ```bash
   echo 'export OBSIDIAN_VAULT_PATH="/path/to/your/obsidian/vault"' >> ~/.bashrc
   source ~/.bashrc
   ```
4. (Optional) Move the binary to your PATH:
   ```bash
   mv jot /usr/local/bin/
   ```
5. (Optional) Add an alias to your shell configuration file (e.g., `.bashrc`, `.zshrc`); call it something else if you 
   Wish to continue using the original `jot` command:
   ```bash
   echo 'alias jot="/path/to/jot"' >> ~/.bashrc
   source ~/.bashrc
   ```
   or if you want to use the original `jot` command:
   ```bash
   echo 'alias anotherjot="/path/to/jot"' >> ~/.bashrc
   source ~/.bashrc
    ```

## Requirements
- Go 1.18 or newer
- Obsidian (for daily note integration), or any path where you want to keep .md note files!

## License

This project is licensed under the GNU General Public License (GPL). See the LICENSE file for details.
