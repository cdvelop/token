package token

import (
	"bytes"
	"compress/flate"
	"io"
)

func compress(data []byte) (out []byte, err string) {
	var compressedData bytes.Buffer

	writer, er := flate.NewWriter(&compressedData, flate.BestCompression)
	if er != nil {
		return nil, er.Error()
	}

	_, er = writer.Write(data)
	if er != nil {
		return nil, er.Error()
	}

	er = writer.Close()
	if er != nil {
		return nil, er.Error()
	}

	return compressedData.Bytes(), ""
}

func decompress(compressedData []byte) (out, err string) {
	reader := flate.NewReader(bytes.NewReader(compressedData))
	defer reader.Close()

	decompressedData, er := io.ReadAll(reader)
	if er != nil {
		return "", "decompress error " + er.Error()
	}

	return string(decompressedData), ""
}
