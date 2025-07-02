# go-codec

A lightweight Go package for base64 encoding/decoding and DEFLATE compression/decompression.

## Installation

```bash
go get github.com/gokpm/go-codec
```

## Usage

```go
import (
    "compress/flate"
    "github.com/gokpm/go-codec"
    "log"
)

// Base64 operations
data := []byte("hello world")
encoded := codec.Encode(data)
decoded, err := codec.Decode(encoded)
if err != nil {
    log.Fatal(err)
}

// DEFLATE compression
compressed, err := codec.Compress(data, flate.DefaultCompression)
if err != nil {
    log.Fatal(err)
}
decompressed, err := codec.Decompress(compressed)
if err != nil {
    log.Fatal(err)
}
```

## Functions

- `Encode(input []byte) string` - Base64 encode bytes to string
- `Decode(input string) ([]byte, error)` - Base64 decode string to bytes  
- `Compress(input []byte, level int) ([]byte, error)` - DEFLATE compress bytes
- `Decompress(input []byte) ([]byte, error)` - DEFLATE decompress bytes

## License

MIT