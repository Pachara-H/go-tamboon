// Package errorcode is a internal service error code for specification
package errorcode

// define error code for config package
const (
	// FailToLoadOmiseConfigPublicKey represent error code when loading Omise config for public key failed
	FailToLoadOmiseConfigPublicKey = iota + 1001
	// FailToLoadOmiseConfigSecretKey represent error code when loading Omise config for secret key failed
	FailToLoadOmiseConfigSecretKey
)

// define error code for validator package
const (
	// FailFileNotExisted represent error code when file isn't existed
	FailFileNotExisted = iota + 2001
	// FailCheckingFile represent error code when file checking failed eg. permission
	FailCheckingFile
	// FailFileNotCSV represent error code when file extension is not .csv
	FailFileNotCSV
)

// define error code for cipher package
const (
	// FailEmptyCipherData represent error code when cipher was null or empty
	FailEmptyCipherData = iota + 3001
	// FailRot128InitReader represent error code when initial reader
	FailRot128InitReader
	// FailRot128Decryption represent error code when rot128 decryption failed
	FailRot128Decryption
)

// define error code for parser package
const (
	// FailEmptyCSVContent represent error code when CSV content was null or empty
	FailEmptyCSVContent = iota + 4001
	// FailReadingCSVRecord represent error code when reading CSV failed/error
	FailReadingCSVRecord
	// FailReadingCSVTimeout represent error code when reading CSV timeout
	FailReadingCSVTimeout
	// FailMissingCSVColumnName represent error code when cannot found some expected column
	FailMissingCSVColumnName
	// FailConvertingCSVAmount represent error code when converting CSV amount failed
	FailConvertingCSVAmount
	// FailConvertingCSVExpMonth  represent error code when converting CSV exp month failed
	FailConvertingCSVExpMonth
	// FailConvertingCSVExpYear represent error code when converting CSV exp year failed
	FailConvertingCSVExpYear
)
