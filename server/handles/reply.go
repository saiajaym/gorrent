package handles

import (
	"PFS/common"
	"net"

	"github.com/boltdb/bolt"
)

//Client interface
type Client struct {
	Con net.Conn
	Db  *bolt.DB
}

//RegisterReply for each file
//it advises if the file registration was a success (Boolean).
func (con *Client) RegisterReply() {

}

//FileListReply Includes the number of files in the list (uint16);
//and for each file, a file name (string) and a file length (uint32)
func (con *Client) FileListReply(list []common.FileShare) {

}

//FileLocationsReply Includes number of endpoints (uint16); then for each endpoint,
//chunks of the file it has, an IP address (uint32) and port (uint16).
func (con *Client) FileLocationsReply(list []common.FileShare) {

}

//ChunkRegisterReply Advises if the chunk registration was a success (Boolean).
func (con Client) ChunkRegisterReply() {

}

//FileChunkReply A stream of bytes repre- senting the requested chunk (array of bytes).
func (con *Client) FileChunkReply() {

}
