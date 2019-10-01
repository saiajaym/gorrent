package handles

import (
	"PFS/common"
	"encoding/gob"
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
func (cli *Client) RegisterReply() {
	enc := gob.NewEncoder(cli.Con)
	file := []common.FileShare{}
	msg := &common.MsgRep{Success: true, List: file}
	enc.Encode(msg)
}

//FileListReply Includes the number of files in the list (uint16);
//and for each file, a file name (string) and a file length (uint32)
func (cli *Client) FileListReply(list []common.FileShare) {
	enc := gob.NewEncoder(cli.Con)
	msg := &common.MsgRep{Success: true, List: list}
	enc.Encode(msg)
}

//FileLocationsReply Includes number of endpoints (uint16); then for each endpoint,
//chunks of the file it has, an IP address (uint32) and port (uint16).
func (cli *Client) FileLocationsReply(list []common.FileShare) {
	enc := gob.NewEncoder(cli.Con)
	msg := &common.MsgRep{Success: true, List: list}
	enc.Encode(msg)
}

//ChunkRegisterReply Advises if the chunk registration was a success (Boolean).
func (cli Client) ChunkRegisterReply() {

}

//FileChunkReply A stream of bytes repre- senting the requested chunk (array of bytes).
func (cli *Client) FileChunkReply() {

}
