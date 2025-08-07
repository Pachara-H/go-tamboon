package cipher

import (
	"bytes"
	"context"
	"log"
	"os"

	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
)

// Rot128DecryptFileContent decrypt file content with rot128 (Caesar) method
func (a *agent) Rot128DecryptFileContent(_ context.Context, path string) (*utilities.SecureByte, error) {
	contentByte, err := os.ReadFile(path) //nolint
	if err != nil {
		log.Print("[ERROR]: read file content failed")
		return nil, Error.NewInternalServerError(Code.FailReadFileContent)
	}

	if len(contentByte) <= 0 {
		log.Print("[ERROR]: content of encrypted file is empty")
		return nil, Error.NewNotFoundError(Code.FailEmptyCipherData)
	}

	reader, err := utilities.NewRot128Reader(bytes.NewBuffer(contentByte))
	if err != nil {
		log.Print("[ERROR]: initial rot128 reader failed")
		return nil, Error.NewInternalServerError(Code.FailRot128InitReader)
	}

	buf := make([]byte, len(contentByte))
	n, err := reader.Read(buf)
	if err != nil || n <= 0 {
		log.Print("[ERROR]: decrypt rot128 failed")
		return nil, Error.NewInternalServerError(Code.FailRot128Decryption)
	}

	return utilities.NewSecureByte(buf[:n]), nil
}
