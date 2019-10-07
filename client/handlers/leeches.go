package handlers

import (
	"PFS/common"
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/boltdb/bolt"
)

//Leech interface
type Leech struct {
	Con net.Conn
	Db  *bolt.DB
}

func getchunk(file string, chunk int, leech Leech) []byte {
	path, _ := leech.GetPath(file)
	info, err := os.Stat(path)
	if os.ErrNotExist == err {
		fmt.Println("Failed to fetch chunk from local disk....")
		return nil
	}
	f, err := os.OpenFile(path, os.O_RDONLY, 0666)
	size := info.Size()
	if err != nil {
		fmt.Println("Failed to fetch chunk from local disk....")
		return nil
	}
	toread := size - int64(common.ChunkSize*(chunk+1))

	if toread < 0 {
		toread = size - int64(common.ChunkSize*chunk)
	} else {
		toread = common.ChunkSize
	}
	glob := make([]byte, toread)
	f.ReadAt(glob, int64(common.ChunkSize*chunk))
	return glob
}

//LeechHandle Handles leech connections
func LeechHandle(con net.Conn, db *bolt.DB) {
	fmt.Println("New Leech Connected: " + con.RemoteAddr().String())
	leech := Leech{con, db}
	dec := gob.NewDecoder(leech.Con)
	req := &common.ChunkReq{}
	dec.Decode(req)
	glob := getchunk(req.File, req.Chunk, leech)
	enc := gob.NewEncoder(con)
	rep := &common.ChunkRep{}
	rep.Glob = glob
	rep.File = req.File
	rep.Chunk = req.Chunk
	enc.Encode(rep)
	con.Close()
}
