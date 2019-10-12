# gorrent
Implemetation of peer to peer file sharing based on paper [Incentives Build Robustness in BitTorrent](http://bittorrent.org/bittorrentecon.pdf "")

Most of the stratergies described in the paper on peer to peer file sharing such as rarest first, fault tolerance, etc. are implemented in go lang. Go threads are leveraged to handle multiple seeds and leeches concurrently. 

The project is meant for learning and undertanding  [Incentives Build Robustness in BitTorrent](http://bittorrent.org/bittorrentecon.pdf "").

## Usage:

### Server/Tracker

Or Tracker in this case, runs on port 7777. All the peer data is stored in tracker.db file.
Tracker can be started by using the below commad.

```bash
go run tracker.go

```

### Seeds/Leeches

The client code in this case acts as both seed and leeche. 
 ```bash
 go run peer.go [PORT NO]
  ```
  
  The common folder has message interface file for a clean communication between trackers and peers, and also between peers and peers.
  
  BoltDB is used both in tracker and peer code for persistant storage of information.
