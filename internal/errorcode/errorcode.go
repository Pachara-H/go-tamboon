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
