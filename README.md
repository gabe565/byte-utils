# byte-utils

Command-line tools to parse and format byte sizes.

## bytefmt

Parse and format byte sizes from the command line.

### Installation

```go
go install gabe565.com/byte-utils/cmd/bytefmt@latest
```

### Usage

Run the command with one or more byte size parameters.

[View flag documentation](docs/bytefmt.md)

#### Examples:

1. Raw size
    ```shell
    $ bytefmt 1048576
    1MiB
    ```
2. Formatted size
    ```shell
    $ bytefmt 1MiB
    1048576
    ```
3. Size from stdin
    ```shell
    $ curl -s example.com | wc -c | bytefmt
    1.23KiB
    ```

## bytect

An additional utility that pretty-prints file sizes or the size of data piped to stdin.

### Installation

```shell
go install gabe565.com/byte-utils/cmd/bytect@latest
```

### Usage

Run the command with a file, list of files, or in a pipe. If the tool detects stdout is a pipe, it will automatically pass the data through.

[View flag documentation](docs/bytect.md)

1. Files
   ```shell
   $ bytect index.html
   1.23KiB  index.html
   ```
2. Stdin with no stdout pipe
   ```shell
   $ curl -s example.com | bytect
   1.23KiB
   ```
3. Stdin with stdout pipe
   ```shell
   $ curl -s example.com | bytect > index.html
   1.23KiB
   $ ls -l index.html
   -rw-r--r--  1 root  root  1256 Jun 25 12:19 index.html
   ```
