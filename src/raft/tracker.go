package raft

// type PeerTracker struct {
// 	nextIndex  uint64
// 	matchIndex uint64

// 	lastAck time.Time
// }

// const activeWindowWidth = 2 * baseElectionTime * time.Millisecond

// func (rf *Raft) quorumActive() bool {
// 	activePeers := 1
// 	for i, tracker := range rf.peerTrackers {
// 		if i != rf.me && time.Since(tracker.lastAck) <= activeWindowWidth {
// 			activePeers++
// 		}
// 	}

// 	return activePeers > len(rf.peers)/2
// }
