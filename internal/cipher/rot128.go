package cipher

import (
	"bytes"

	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
)

// Rot128Decrypt decrypt cipherText with rot128 (Caesar) method
func (a *agent) Rot128Decrypt(cipherByte []byte) ([]byte, error) {
	if len(cipherByte) <= 0 {
		return nil, Error.NewInternalServerError(Code.FailEmptyCipherData)
	}

	reader, err := utilities.NewRot128Reader(bytes.NewBuffer(cipherByte))
	if err != nil {
		return nil, Error.NewInternalServerError(Code.FailRot128InitReader)
	}

	buf := make([]byte, len(cipherByte))
	n, err := reader.Read(buf)
	if err != nil || n <= 0 {
		return nil, Error.NewInternalServerError(Code.FailRot128Decryption)
	}

	return buf[:n], nil
}
