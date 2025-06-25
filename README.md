# byte-utils

Command-line tools to parse and format byte sizes.

## bytefmt

Parse and format byte sizes from the command line.

### Installation

```go
go install gabe565.com/byte-utils/cmd/bytefmt@latest
```

### Usage

Accepts one or more parameters and parses or formats the value dynamically:

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

An additional utility that pretty-prints file sizes or the size of data piped to stdin

### Installation

```shell
go install gabe565.com/byte-utils/cmd/bytect@latest
```

### Usage

1. Files
   ```shell
   $ bytect index.html
   1.23KiB  index.html
   ```
2. Stdin
   ```shell
   $ curl -s example.com | bytect
   1.23KiB
   ```
