package raft

import (
	"math/rand"
	"time"
)

const baseElectionTime = 300
const None = -1

func (rf *Raft) StartElection() {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.resetElectionTime()
	rf.becomeCandidate()

	done := false
	votes := 1
	term := rf.currentTerm
	args := RequestVoteArgs{
		Term:        rf.currentTerm,
		CandidateId: rf.me,
	}

	for i := range rf.peers {
		if i == rf.me {
			continue
		}

		go func(serverId int) {
			reply := RequestVoteReply{}
			ok := rf.sendRequestVote(serverId, &args, &reply)
			if !ok || !reply.VoteGranted {
				return
			}
			rf.mu.Lock()
			defer rf.mu.Unlock()

			if reply.Term < rf.currentTerm {
				return
			}
			votes++

			if done || votes <= len(rf.peers)/2 {
				return
			}
			done = true
			if rf.state != CANDIDATE || rf.currentTerm != term {
				return
			}
			rf.state = LEADER
			go rf.StartAppendEntries(true)
		}(i)
	}
}

func (rf *Raft) pastElectionTimeOut() bool {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	return time.Since(rf.lastElection) > rf.electionTimer
}

func (rf *Raft) resetElectionTime() {
	electionTimeOut := baseElectionTime + (rand.Int63() % baseElectionTime)
	rf.electionTimer = time.Duration(electionTimeOut) * time.Millisecond
	rf.lastElection = time.Now()
}

func (rf *Raft) becomeCandidate() {
	rf.state = CANDIDATE
	rf.currentTerm++
	rf.votedFor = rf.me

	// DPrintf("开始新一轮的投票，成为候选人....", "Success\n")
}

func (rf *Raft) toFollower() {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	rf.state = FOLLOWER
	rf.votedFor = None
}
