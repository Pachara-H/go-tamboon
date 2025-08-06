package utilities

// SecureByte provides secure handling of sensitive byte data with clearing
type SecureByte struct {
	data []byte
	size int
}

// NewSecureByte creates a new SecureByte from a byte slice
func NewSecureByte(b []byte) *SecureByte {
	data := make([]byte, len(b))
	copy(data, b)
	return &SecureByte{
		data: data,
		size: len(b),
	}
}

// Bytes returns a copy of the underlying byte slice
func (sb *SecureByte) Bytes() []byte {
	if sb.data == nil {
		return nil
	}
	result := make([]byte, sb.size)
	copy(result, sb.data[:sb.size])
	return result
}

// String returns the string representation of the secure byte data
func (sb *SecureByte) String() string {
	if sb.data == nil {
		return ""
	}
	return string(sb.data[:sb.size])
}

// Len returns the length of the secure byte data
func (sb *SecureByte) Len() int {
	return sb.size
}

// IsEmpty returns true if the secure byte data is empty or cleared
func (sb *SecureByte) IsEmpty() bool {
	return sb.data == nil || sb.size == 0
}

// Clear securely wipes the sensitive byte data from memory
func (sb *SecureByte) Clear() {
	if sb.data == nil {
		return
	}

	for i := range sb.data {
		sb.data[i] = 0
	}

	// Clear the slice reference
	sb.data = nil
	sb.size = 0
}
