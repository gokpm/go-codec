package codec

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

func Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func Decode(input string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Compress(input []byte, level int) ([]byte, error) {
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, level)
	if err != nil {
		return nil, err
	}
	_, err = writer.Write(input)
	if err != nil {
		writer.Close()
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decompress(input []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(input))
	defer reader.Close()
	output, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return output, nil
}
