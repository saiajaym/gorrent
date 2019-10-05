package handlers

import (
	"PFS/common"
	"encoding/gob"
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
)

type pair struct {
	ipaddr string
	chunk  int
}

type kv struct {
	key pair
	val int
}

func buildTree(list []common.FileLocation) []kv {
	var tree map[pair]int
	for _, l := range list {
		ip := l.IpAddr
		ch := strings.Split(l.Chunks, ",")
		for _, c := range ch {
			cint, _ := strconv.Atoi(c)
			tmp := pair{ip, cint}
			tree[tmp]++
		}
	}
	kvl := make([]kv, len(tree))
	for k, v := range tree {
		kvl = append(kvl, kv{k, v})
	}

	sort.Slice(kvl, func(i, j int) bool {
		return kvl[i].val < kvl[j].val
	})

	return kvl
}

func get(ip string, chunk int, file string) error {
	con, err := net.Dial(ip, "tcp")
	defer con.Close()
	if err != nil {
		fmt.Printf("Failed to get chunk %d from %s\n", chunk, ip)
		return err
	}
	req := &common.ChunkReq{}
	req.File = file
	req.Chunk = chunk
	enc := gob.NewEncoder(con)
	enc.Encode(req)
	rep := &common.ChunkRep{}
	dec := gob.NewDecoder(con)
	dec.Decode(rep)

}

//Download Manages doenloads after getting the file Locations
func Download(list []common.FileLocation, file string) {
	if len(list) == 0 {
		return
	}

	fmt.Println("Starting Download...")
	chunks := buildTree(list)

	var toDownload kv

	for len(chunks) > 0 {
		toDownload = chunks[0]
		chunks = chunks[1:]
		get(toDownload.key.ipaddr, toDownload.key.chunk, file)
	}

}
