# Hulu

Hulu is a tool for converting LaTeX code within Markdown files to image links for easy rendering on websites. This is achieved by converting the LaTeX code to a Zhihu LaTeX image link.

## Installation

To install Hulu, use the following command:

```
go install github.com/levinion/hulu
```

## Usage

Hulu is used as a command-line tool. It has two modes of operation: 

- `-c` : This mode outputs the parsed file to a new file with a different name, in the same directory as the original file. 

- `-w` : This mode overwrites the original file with the parsed output.

In both cases, a path to the input file must be provided. Additionally, a path for the output file can be provided as a second parameter.

Here are some examples:

```
hulu /path/to/input/file.md /path/to/output/file.md
```

This will parse the LaTeX code in the input file and output the results to a new file at `/path/to/output/file.md`.

```
hulu -w /path/to/input/file.md
```

This will parse the LaTeX code in the input file and overwrite the original file.

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on the GitHub repository. 

## License

Hulu is released under the MIT License. See the LICENSE file for more information.
