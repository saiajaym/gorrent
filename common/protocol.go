package common

const (
	//ChunkSize a cont chunk value for files 1MB
	ChunkSize = 1048576
)

//FileShare common structure for file metadata
type FileShare struct {
	Name string
	Size string
}

//FileReg Format for file registration with server
type FileReg struct {
	IPAddr string
	List   []FileShare
}
