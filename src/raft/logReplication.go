package raft

import "time"

func (rf *Raft) resetHeartBeatTimer() {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.lastHeartBeat = time.Now()
}

func (rf *Raft) pastHeartBeatTimeOut() bool {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return time.Since(rf.lastHeartBeat) > rf.heartBeatTimer
}
