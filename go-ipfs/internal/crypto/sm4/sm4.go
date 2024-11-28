package sm4

import (
	"crypto/cipher"
	"io"

	"github.com/ipfs/kubo/internal/logger"
	"github.com/tjfoc/gmsm/sm4"
)

const (
	key   = "xinjianglihuasuo"
	ivKey = ""
)

var (
	iv = []byte{102, 42, 23, 108, 82, 23, 77, 9, 106, 48, 26, 255, 189, 181, 134, 220}
)

func Encrypt(reader io.Reader) (*cipherReader, error) {
	logger.GetCryptoLogger().Debugf("starting encrypt with key %v iv %v", key, iv)
	encBlock, err := sm4.NewCipher([]byte(key))
	if err != nil {
		logger.GetCryptoLogger().Errorf("创建SM4块失败 %v", err)
		return nil, err
	}

	encStream := cipher.NewCTR(encBlock, iv)

	return NewCipherReader(reader, encStream), nil
}

func Decrypt(reader io.Reader) (*decipherReader, error) {
	logger.GetCryptoLogger().Debugf("starting decrypt with key %v iv %v", key, iv)
	decBlock, err := sm4.NewCipher([]byte(key))
	if err != nil {
		logger.GetCryptoLogger().Errorf("创建SM4块失败 %v", err)
		return nil, err
	}
	decStream := cipher.NewCTR(decBlock, iv)
	return NewDecipherReader(reader, decStream), nil
}

func NewCipherReader(reader io.Reader, stream cipher.Stream) *cipherReader {
	return &cipherReader{
		reader: reader,
		stream: stream,
	}
}

func NewDecipherReader(reader io.Reader, stream cipher.Stream) *decipherReader {
	return &decipherReader{
		reader: reader,
		stream: stream,
	}
}

// cipherReader 实现了io.Reader接口，用于读取加密后的数据
type cipherReader struct {
	reader io.Reader
	stream cipher.Stream
	buffer []byte
}

// Read 实现io.Reader的Read方法，用于读取加密后的数据
func (cr *cipherReader) Read(p []byte) (int, error) {
	logger.GetCryptoLogger().Debugf("starting cipherReader Read ...")
	n, err := cr.reader.Read(p)
	if err != nil && err != io.EOF {
		return 0, err
	}

	if n > 0 {
		cr.stream.XORKeyStream(p[:n], p[:n])
	}
	return n, err
}

// decipherReader 实现了io.Reader接口，用于读取解密后的数据
type decipherReader struct {
	reader io.Reader
	stream cipher.Stream
	buffer []byte
}

// Read 实现io.Reader的Read方法，用于读取解密后的数据
func (dr *decipherReader) Read(p []byte) (int, error) {
	logger.GetCryptoLogger().Debugf("starting decipherReader Read ...")
	n, err := dr.reader.Read(p)
	if err != nil && err != io.EOF {
		return 0, err
	}

	if n > 0 {
		dr.stream.XORKeyStream(p[:n], p[:n])
	}

	return n, err
}
