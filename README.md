# wnote (Work Notes) 
> Note: This is work in progress. The description in the readme is a reference of capabilities for the features to be implemented

`wnote` is a minimalist, context-aware CLI tool built in Go. For developers who need to capture mental state and track tasks without breaking their workflow.

---

## 1. Core Principles

- **Context-First**: Automatically associates notes with the current Git repository and branch.
- **CLI Native**: Designed to be piped, filtered, and integrated with tools like `fzf` and `nushell`.
- **Markdown-Centric**: Uses Markdown for note content to support rich formatting (checklists, code blocks) in `$EDITOR`.
- **Single Binary**: Compiles to a static Go binary for instant startup and portability.

---

## 2. Usage Guide

### Note Creation
Capture thoughts instantly. The tool identifies the current Git root and branch automatically.

```bash
wn todo "Refactor the .NET auth middleware"
wn checkpoint "Completed GRD resource mapping"
wn followup "Check status of gdb symbol issues in Chromium"
```
### Retrieval & Filtering
Retrieve notes based on scope and type.

```bash
# Get active todos for the current repo (Default)
wn get todo

# Get all notes of a specific type across all repositories
wn get todo -a

# Paginate or limit output
wn get checkpoint --last 10
Advanced Management
Use the power of Unix piping and your preferred text editor.

Bash
# Multi-select notes to delete using fzf
wn get --ids | fzf -m | xargs wn prune

# Open a note in your $EDITOR as a temporary Markdown file
wn edit <id>
```

### Structured Output
Integrate with Nushell or other automation tools.

```bash
wn get --json | from json | where type == "todo"
```

## 3. Data Architecture

Storage
Engine: SQLite (Single file located at ~/.local/share/wn/notes.db).

Reasoning: Provides structured querying (filtering by repo/type) while remaining zero-config and local-first.

Metadata Schema
Each entry tracks:

- ID: Unique integer for reference.
- Type: todo, checkpoint, followup, or idea.
- Content: Raw Markdown text.
- Repository: Absolute path to the Git root.
- Branch: The active Git branch at creation.
- Status: active, done, or archived.
- Timestamps: created_at and updated_at.

## 4. Operational Workflows
The "Edit" Buffer
When `wn edit <id>` is invoked:
1. wn fetches the note from SQLite.
2. It generates a temporary .md file with YAML frontmatter for metadata.
3. It opens the file in the user's $EDITOR.
4. Upon save, wn parses the file, validates the frontmatter/structure, and updates the database.

### Git Detection Logic
The tool traverses upward from the current working directory until it finds a .git directory. If no .git directory is found, the note is filed under a global or cwd context.
Neovim: Potential to export todo types into a Quickfix list for immediate code navigation.

FZF: Native support for interactive selection via ID-only output flags.
