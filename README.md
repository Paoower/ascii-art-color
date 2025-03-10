# ASCII Art Color

A command-line utility that generates colorized ASCII art text.

## Description

This program extends the basic ASCII art functionality by adding color support. It converts input strings into colorful ASCII art patterns using predefined banner templates and ANSI color codes.

## Features

- All base functionality of the ascii-art program
- Support for colored output using ANSI color codes
- Color can be applied to the entire text or specific characters

## Usage

```bash
go run . [STRING] [BANNER] [COLOR]
go run . [STRING] [BANNER] [COLOR=<color>]
go run . [STRING] [BANNER] [LETTER:POSITION:COLOR]
```

Examples:
```bash
go run . "Hello World" standard red
go run . "Hello World" standard color=blue
go run . "Hello World" standard letter:0-5:red
```

## Implementation Details

- Uses ANSI escape sequences for terminal coloring
- Preserves original banner formatting
- Supports selective coloring of specific characters or ranges

## Requirements

- Go 1.13 or higher
- Terminal with ANSI color support
