# ASCII-ART-Generator

A command-line tool written in Go that takes a string and renders it in large ASCII art using banner fonts.

```
 _   _          _   _          
| | | |   ___  | | | |   ___  
| |_| |  / _ \ | | | |  / _ \ 
|  _  | |  __/ | | | | | (_) |
|_| |_|  \___| |_| |_|  \___/ 
```

---

## How It Works

Each character in the input is looked up in a banner file — a `.txt` file where every printable ASCII character (space through `~`) is represented as 8 rows of ASCII art. The program builds a character map from the file, then renders the input row by row, stitching each character's art horizontally to produce the final output.

---

## Usage

```bash
go run . [STRING]
go run . [STRING] [BANNER]
```

- `STRING` — the text to render (use quotes for multi-word input)
- `BANNER` — optional banner style (default: `standard`)

### Available Banners

| Banner | Description |
|---|---|
| `standard` | Classic block letters |
| `shadow` | Letters with a shadow effect |
| `thinkertoy` | Lighter, symbolic style |

---

## Examples

```bash
go run . "Hello"
go run . "Hello World" shadow
go run . "Hello\nWorld" thinkertoy
```

Using `\n` in your string inserts a new line in the output:

```bash
go run . "Hi\n\nThere"
# renders "Hi", then a blank line, then "There"
```

---

## Edge Cases Handled

- **Unprintable characters** — anything outside ASCII 32–126 is replaced with a space
- **Newlines** — `\n` in the input is converted to a real line break before rendering
- **Blank lines** — `\n\n` produces a visible blank line between rendered text
- **Windows line endings** — `\r` is stripped from banner files automatically
- **Missing characters** — if a character has no art in the banner file, it falls back to a space

---

## Project Structure

```
.
├── main.go
├── standard.txt
├── shadow.txt
└── thinkertoy.txt
```

---

## Author

[Israel-light](https://github.com/Israel-light)
