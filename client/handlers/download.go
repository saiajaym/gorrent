package handlers

import (
	"PFS/common"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

type kv struct {
	chunk  int
	ipAddr []string
}

func buildTree(list []common.FileLocation) []kv {
	//var tree map[int][]string
	tree := make(map[int][]string)
	for _, l := range list {
		fmt.Println(l)
		ip := l.IpAddr
		l.Chunks = l.Chunks[0 : len(l.Chunks)-1]
		chunks := strings.Split(l.Chunks, ",")
		for _, c := range chunks {
			ch, _ := strconv.Atoi(c)

			tree[ch] = append(tree[ch], ip)

		}
	}
	var kvl []kv
	for k, v := range tree {
		fmt.Println(k, v)
		kvl = append(kvl, kv{k, v})
	}
	sort.Slice(kvl, func(i, j int) bool {
		return len(kvl[i].ipAddr) < len(kvl[j].ipAddr)
	})

	return kvl
}

func get(ip string, chunk int, file string) ([]byte, error) {
	con, err := net.Dial("tcp", ip)
	//	defer con.Close()
	if err != nil {
		fmt.Printf("Failed to get chunk %d from %s\n", chunk, ip)
		return nil, err
	}
	req := &common.ChunkReq{}
	req.File = file
	req.Chunk = chunk
	enc := gob.NewEncoder(con)
	enc.Encode(req)
	rep := &common.ChunkRep{}
	dec := gob.NewDecoder(con)
	dec.Decode(rep)
	con.Close()
	return rep.Glob, err

}

func save(glob []byte, chunk int, f *os.File) {
	f.WriteAt(glob, int64(chunk)*int64(common.ChunkSize))
}

//Download Manages doenloads after getting the file Locations
func Download(list []common.FileLocation, file string) {
	if len(list) == 0 {
		return
	}
	f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	fmt.Println("Starting Download...")
	chunks := buildTree(list)

	var toDownload kv

	for len(chunks) > 0 {
		toDownload = chunks[0]
		chunks = chunks[1:]
		fmt.Printf("Downloading ... %d from %s \n", toDownload.chunk, toDownload.ipAddr[0])
		glob, err := get(toDownload.ipAddr[0], toDownload.chunk, file)
		if err != nil {
			fmt.Println("Failed chunk download ... retrying ..." + err.Error())
		} else {
			save(glob, toDownload.chunk, f)
		}
	}

}
