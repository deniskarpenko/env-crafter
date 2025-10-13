package valueobjects

type EnvFile struct {
	fileName string
	file     []byte
}

func (ef EnvFile) FileName() string {
	return ef.fileName
}

func (ef EnvFile) File() []byte {
	return ef.file
}

func (ef EnvFile) ToYaml() string {
	return string(ef.file)
}

func NewEnvFile(fileName string, fileData []byte) EnvFile {
	return EnvFile{fileName, fileData}
}
