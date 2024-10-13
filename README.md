# Advanced Text Processor

## Description

This project is an advanced text processing tool implemented in Go. It performs various text transformations and formatting operations on input text files. The processor can handle multiple types of modifications, including numeric base conversions, case changes, punctuation formatting, and more.

## Features

- Numeric Conversions:
  - Hexadecimal to decimal conversion
  - Binary to decimal conversion
- Text Case Modifications:
  - Uppercase conversion
  - Lowercase conversion
  - Capitalization
- Grammar Improvements:
  - "a" to "an" conversion before words starting with vowels or 'h'
- Complex Modifiers:
  - Support for modifiers like "(up, 3)" to affect multiple words
- Punctuation Formatting:
  - Correct spacing around punctuation marks
  - Special handling for ellipsis (...) and interrobang (!?)
- Quote Formatting:
  - Proper placement of single quotes

## Requirements

- Go (version 1.16 or later recommended)

## Installation

1. Clone the repository:
   ```
   git clone https://platform.zone01.gr/git/gpapadopoulos/go-reloaded.git
   ```
2. Navigate to the project directory:
   ```
   cd go-reloaded 
   ```

## Usage

Run the program from the command line, specifying input and output files:

```
go run . <input_file> <output_file>
```

Example:
```
go run . sample.txt result.txt
```

This will process the content of `sample.txt` and write the result to `result.txt`.

## Input Format

The input text can include various modifiers and formatting instructions:

- `(hex)` after a number to convert from hexadecimal to decimal
- `(bin)` after a number to convert from binary to decimal
- `(up)`, `(low)`, `(cap)` for case modifications
- `(up, n)`, `(low, n)`, `(cap, n)` to affect multiple words
- Regular text with punctuation for formatting

## Examples

Input:
```
The 42 (hex) cats saw 1101 (bin) mice. let's MAKE (low, 2) it 'awesome' ... no?
```

Output:
```
The 66 cats saw 13 mice. let's make it 'awesome'... no?
```

## Contributing

Contributions to improve the text processor are welcome. Please feel free to submit pull requests or open issues to discuss potential enhancements.

## License

[Coding Academy Zone01 Athens,Greece.]

