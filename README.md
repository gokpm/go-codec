# go-codec

A Go package providing base64 encoding/decoding and DEFLATE compression/decompression utilities with structured logging.

## Installation

```bash
go get github.com/gokpm/go-codec
```

## Usage

```go
import (
    "context"
    "github.com/gokpm/go-codec"
)

ctx := context.Background()

// Base64 encoding/decoding
data := []byte("hello world")
encoded := codec.Encode(ctx, data)
decoded, err := codec.Decode(ctx, encoded)

// DEFLATE compression/decompression
compressed, err := codec.Compress(ctx, data, flate.DefaultCompression)
decompressed, err := codec.Decompress(ctx, compressed)
```

## Functions

- `Encode(ctx, input)` - Base64 encode bytes to string
- `Decode(ctx, input)` - Base64 decode string to bytes
- `Compress(ctx, input, level)` - DEFLATE compress bytes
- `Decompress(ctx, input)` - DEFLATE decompress bytes

All functions include structured logging via [go-sig](https://github.com/gokpm/go-sig).

## License

MIT