# wnote (Work Notes) 

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
