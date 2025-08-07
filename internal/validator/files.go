package validator

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
)

// IsFileExist check exiting of file
func (a *agent) IsFileExist(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Print("[ERROR]: file doesn't exist")
			return Error.NewNotFoundError(Code.FailFileNotExisted)
		}
		log.Print("[ERROR]: file checking failed")
		return Error.NewInternalServerError(Code.FailCheckingFile)
	}
	return nil
}

// IsCSVExtension check is file extension .csv or not
func (a *agent) IsCSVExtension(filePath string) error {
	ext := filepath.Ext(filePath)
	if ext != ".csv" {
		log.Print("[ERROR]: invalid expected sub file extension .csv")
		return Error.NewUnsupportedMediaTypeError(Code.FailFileNotCSV)
	}
	return nil
}

// IsCSVExtension check is file extension .csv.rot128 or not
func (a *agent) IsCSVRot128Extension(filePath string) error {
	ext := filepath.Ext(filePath)
	if ext != ".rot128" {
		log.Print("[ERROR]: invalid expected file extension .rot128")
		return Error.NewUnsupportedMediaTypeError(Code.FailFileNotRot128)
	}
	return a.IsCSVExtension(strings.TrimSuffix(filePath, ".rot128"))
}
