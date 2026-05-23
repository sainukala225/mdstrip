# mdstrip
A CLI tool to convert Markdown files to plain text by stripping all Markdown formatting.

## Requirements
- Go 1.26 or higher

## Installation

### Using go install
```bash
go install github.com/sainukala225/mdstrip@latest
```

### Manual
Clone the repository, build the binary, and move it to your local bin:
```bash
git clone https://github.com/sainukala225/mdstrip
cd mdstrip
go build -o mdstrip .
cp mdstrip ~/.local/bin/
```

## Usage
```bash
mdstrip [--output/-o <file.txt>] <file.md>
```

Read from stdin:
```bash
cat file.md | mdstrip
```

You can also pass the input file via flag:
```bash
mdstrip --filename <file.md> [--output <file.txt>]
```

## Options

| Flag | Short | Description | Default |
|---|---|---|---|
| `--filename` | `-f` | Input markdown file | positional arg |
| `--output` | `-o` | Output text file | stdout |
| `--version` | `-v` | Print version | |
| `--help` | `-h` | Print help message | |

## Examples

Print to stdout:
```bash
mdstrip README.md
```

Write to file:
```bash
mdstrip --output out.txt README.md
```

Using short flags:
```bash
mdstrip -o out.txt README.md
```

Using flags for both:
```bash
mdstrip --filename README.md --output out.txt
mdstrip -f README.md -o out.txt
```

Read from stdin:
```bash
echo "# Hello **world**" | mdstrip
```

## What gets stripped
- Headings (`#`, `##`, etc.)
- Bold and italic (`**bold**`, `*italic*`)
- Strikethrough (`~~text~~`)
- Inline code and code blocks
- Links — URL removed, text kept
- Images — removed entirely
- Blockquotes
- List markers (`-`, `*`, `+`, `1.`)
- Tables — formatting removed, content kept
- Horizontal rules