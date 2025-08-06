package validator

import (
	"os"
	"path/filepath"

	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
)

// IsFileExist check exiting of file
func (a *agent) IsFileExist(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return Error.NewNotFoundError(Code.FailFileNotExisted)
		}
		return Error.NewInternalServerError(Code.FailCheckingFile)
	}
	return nil
}

// IsCSVExtension check is file extension .csv or not
func (a *agent) IsCSVExtension(filePath string) error {
	ext := filepath.Ext(filePath)
	if ext != ".csv" {
		return Error.NewUnsupportedMediaTypeError(Code.FailFileNotCSV)
	}
	return nil
}
