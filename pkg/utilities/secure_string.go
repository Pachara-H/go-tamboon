package utilities

// SecureString provides secure handling of sensitive data with clearing
type SecureString struct {
	data []byte
	size int
}

// NewSecureString creates a new SecureString from a regular string
func NewSecureString(s string) *SecureString {
	data := make([]byte, len(s))
	copy(data, []byte(s))
	return &SecureString{
		data: data,
		size: len(s),
	}
}

// NewSecureStringFromByte creates a new SecureString from a regular byte
func NewSecureStringFromByte(b []byte) *SecureString {
	data := make([]byte, len(b))
	copy(data, b)
	return &SecureString{
		data: data,
		size: len(b),
	}
}

// String returns the string representation of the secure data
func (s *SecureString) String() string {
	if s.data == nil {
		return ""
	}
	return string(s.data[:s.size])
}

// Bytes returns a copy of the underlying byte slice
func (s *SecureString) Bytes() []byte {
	if s.data == nil {
		return nil
	}
	result := make([]byte, s.size)
	copy(result, s.data[:s.size])
	return result
}

// Len returns the length of the secure string
func (s *SecureString) Len() int {
	return s.size
}

// IsEmpty returns true if the secure string is empty or cleared
func (s *SecureString) IsEmpty() bool {
	return s.data == nil || s.size == 0
}

// Clear securely wipes the sensitive data from memory
func (s *SecureString) Clear() {
	if s.data == nil {
		return
	}

	for i := range s.data {
		s.data[i] = 0
	}

	// Clear the slice reference
	s.data = nil
	s.size = 0
}
