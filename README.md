## ASCII-ART

ASCII-ART is a lightweight command-line tool built in Go that converts a given string into a graphical representation using ASCII characters. It takes a string as input and outputs the string in ASCII art format, where each character of the input string is represented by a corresponding ASCII art character.

### Features

- **String Conversion**: Converts any given string into ASCII art format.
- **Support for Special Characters**: Handles numbers, letters, spaces, and special characters in the input string.
- **Newline Support**: Supports newline characters, allowing for multi-line ASCII art output.
- **Simple to Use**: Easy-to-understand command-line interface for quick string conversion.

### Usage

To use ASCII Art Converter, simply provide a string as an argument when running the program. For example:

```bash
go run . "Hello"
```

This will output the ASCII representation of the string "Hello".

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
```

### Installation

To install ASCII Art Converter, clone the repository and navigate to the project directory:

```bash
git clone https://github.com/e-aub/ascii-art
cd ascii-art
```

Then, run the program:

```bash
go run . <string>
```

