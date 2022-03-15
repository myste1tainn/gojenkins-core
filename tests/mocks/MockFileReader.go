package mocks

type MockFileReader struct {
	FileContent string
}

func (m MockFileReader) ReadFile(path string) ([]byte, error) {
	return []byte(m.FileContent), nil
}
