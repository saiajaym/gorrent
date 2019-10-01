package handles

import (
	"PFS/common"
	"encoding/gob"
	"fmt"
)

//RequestHandle handles client request/ acts as router
func (cli *Client) RequestHandle() {
	defer cli.Con.Close()
	dec := gob.NewDecoder(cli.Con)
	msg := &common.MsgReq{}
	dec.Decode(msg)
	fmt.Println("Received req from client of type:" + msg.MessageType)

	switch msg.MessageType {
	case "Register":
		req := msg.Reg
		suc := cli.registerReq(req.IPAddr, req.List)
		if suc {
			cli.RegisterReply()
		}
	case "FileList":
		cli.fileListReq()
	case "FileLocation":
		req := msg.Loc
		cli.fileLocReq(req.File)
	default:
		fmt.Println("RequestHandle invalid message tyoe")
	}

}

//Tells the server what files the peer wants to share with the network.
//takes in the IP address (uint32) and port (uint16)files to register (uint16)
//the number of files to register (uint16);
//and for every file, a file name (string) and its length (uint32)
func (cli *Client) registerReq(ipAdd string, list []common.FileShare) bool {
	flag := true

	err := cli.AddList(list)
	if err != nil {
		fmt.Println("Register req at addlist failed" + err.Error())
		flag = false
	}

	err = cli.AddIPAdd(list, ipAdd)

	if err != nil {
		fmt.Println("Register req at addlist failed" + err.Error())
		flag = false
	}

	return flag
}

func (cli *Client) fileListReq() ([]common.FileShare, error) {

	list, err := cli.GetList()
	if err != nil {
		fmt.Println("fileList req failed " + err.Error())
		return list, nil
	}
	cli.FileListReply(list)
	return list, err
}

func (cli *Client) fileLocReq(file common.FileShare) bool {
	if len(file.Name) == 0 {
		fmt.Println("FileLocation empty Request")
		return false
	}
	flag := true
	list, err := cli.FileLoc(file)
	if err != nil {
		fmt.Println("filelocreq " + err.Error())
		return false
	}
	cli.FileLocationsReply(list)
	return flag
}
