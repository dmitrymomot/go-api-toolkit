package storage

// Filer is an interface of file structure
type Filer interface {
	FileName() string
	OriginFileName() string
	RelativePath() string
	Size() int64
	MIME() string
}

// File structure
type File struct {
	originName string
	name       string
	path       string
	size       int64
	mime       string
}

// FileName returns file name
func (f *File) FileName() string {
	return f.name
}

// OriginFileName returns origin file name
func (f *File) OriginFileName() string {
	return f.originName
}

// RelativePath returns relative file path
func (f *File) RelativePath() string {
	return f.path
}

// Size returns file size
func (f *File) Size() int64 {
	return f.size
}

// MIME returns file mime type
func (f *File) MIME() string {
	return f.mime
}
