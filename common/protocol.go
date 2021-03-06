package common

const (
	//ChunkSize a cont chunk value for files 1MB
	ChunkSize = 1000000
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
	Name   string
	IpAddr string
	Chunks string
}

//MsgReq enables coomunication between client and server with type of req
//Register to registerfiles
//FileList to list files
type MsgReq struct {
	MessageType   string
	Reg           FileReg
	Loc           FileLocation
	ChunkRegister ChunkRegReq
}

//MsgRep struct for msg reply for MsgReq
type MsgRep struct {
	Success bool
	List    []FileShare
	Loc     []FileLocation
}

//ChunkReq to req chunk from a client
type ChunkReq struct {
	File  string
	Chunk int
}

//ChunkRep reply for chunk request peer to peer
type ChunkRep struct {
	File  string
	Chunk int
	Glob  []byte
}

//ChunkRegReq Registers chunk with server
type ChunkRegReq struct {
	File   string
	IPAddr string
	Chunk  int
}
