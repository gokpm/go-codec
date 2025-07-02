package codec

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

// Encode converts byte slice to base64 encoded string
func Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// Decode converts base64 encoded string back to byte slice
func Decode(input string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Compress compresses byte slice using DEFLATE algorithm with specified compression level
// level: compression level (flate.NoCompression, flate.BestSpeed, flate.BestCompression, flate.DefaultCompression)
func Compress(input []byte, level int) ([]byte, error) {
	var buf bytes.Buffer
	// Create DEFLATE writer with specified compression level
	writer, err := flate.NewWriter(&buf, level)
	if err != nil {
		return nil, err
	}
	// Write input data to compressor
	_, err = writer.Write(input)
	if err != nil {
		// Close writer to flush any remaining data
		writer.Close()
		return nil, err
	}
	// Close writer to flush any remaining data
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress decompresses DEFLATE compressed byte slice back to original data
func Decompress(input []byte) ([]byte, error) {
	// Create DEFLATE reader from compressed data
	reader := flate.NewReader(bytes.NewReader(input))
	defer reader.Close()
	// Read all decompressed data
	output, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return output, nil
}
