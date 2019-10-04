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

//FileLocation message type for location request
type FileLocation struct {
	File FileShare
}

//MsgReq enables coomunication between client and server with type of req
//Register to registerfiles
//FileList to list files
type MsgReq struct {
	MessageType string
	Reg         FileReg
	Loc         FileLocation
}

//MsgRep struct for msg reply for MsgReq
type MsgRep struct {
	Success bool
	List    []FileShare
}
