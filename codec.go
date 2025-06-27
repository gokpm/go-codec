package codec

import (
	"bytes"
	"compress/flate"
	"context"
	"encoding/base64"
	"io"

	"github.com/gokpm/go-sig"
)

func Encode(ctx context.Context, input []byte) string {
	log := sig.Start(ctx)
	defer log.End()
	return base64.StdEncoding.EncodeToString(input)
}

func Decode(ctx context.Context, input string) ([]byte, error) {
	log := sig.Start(ctx)
	defer log.End()
	bytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Compress(ctx context.Context, input []byte, level int) ([]byte, error) {
	log := sig.Start(ctx)
	defer log.End()
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

func Decompress(ctx context.Context, input []byte) ([]byte, error) {
	log := sig.Start(ctx)
	defer log.End()
	reader := flate.NewReader(bytes.NewReader(input))
	defer reader.Close()
	output, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return output, nil
}
