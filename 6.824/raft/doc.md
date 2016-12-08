PACKAGE DOCUMENTATION

package raft
    import "."


CONSTANTS

const Debug = 0
    Debugging

FUNCTIONS

func DPrintf(format string, a ...interface{}) (n int, err error)

TYPES

type ApplyMsg struct {
    Index       int
    Command     interface{}
    UseSnapshot bool   // ignore for lab2; only used in lab3
    Snapshot    []byte // ignore for lab2; only used in lab3
}
    as each Raft peer becomes aware that successive log entries are
    committed, the peer should send an ApplyMsg to the service (or tester)
    on the same server, via the applyCh passed to Make().

type Persister struct {
    // contains filtered or unexported fields
}

func MakePersister() *Persister

func (ps *Persister) Copy() *Persister

func (ps *Persister) RaftStateSize() int

func (ps *Persister) ReadRaftState() []byte

func (ps *Persister) ReadSnapshot() []byte

func (ps *Persister) SaveRaftState(data []byte)

func (ps *Persister) SaveSnapshot(snapshot []byte)

type Raft struct {
    // contains filtered or unexported fields
}
    A Go object implementing a single Raft peer.

func Make(peers []*labrpc.ClientEnd, me int,
    persister *Persister, applyCh chan ApplyMsg) *Raft
    the service or tester wants to create a Raft server. the ports of all
    the Raft servers (including this one) are in peers[]. this server's port
    is peers[me]. all the servers' peers[] arrays have the same order.
    persister is a place for this server to save its persistent state, and
    also initially holds the most recent saved state, if any. applyCh is a
    channel on which the tester or service expects Raft to send ApplyMsg
    messages. Make() must return quickly, so it should start goroutines for
    any long-running work.

func (rf *Raft) GetState() (int, bool)
    return currentTerm and whether this server believes it is the leader.

func (rf *Raft) Kill()
    the tester calls Kill() when a Raft instance won't be needed again. you
    are not required to do anything in Kill(), but it might be convenient to
    (for example) turn off debug output from this instance.

func (rf *Raft) RequestVote(args RequestVoteArgs, reply *RequestVoteReply)
    example RequestVote RPC handler.

func (rf *Raft) Start(command interface{}) (int, int, bool)
    the service using Raft (e.g. a k/v server) wants to start agreement on
    the next command to be appended to Raft's log. if this server isn't the
    leader, returns false. otherwise start the agreement and return
    immediately. there is no guarantee that this command will ever be
    committed to the Raft log, since the leader may fail or lose an
    election.

    the first return value is the index that the command will appear at if
    it's ever committed. the second return value is the current term. the
    third return value is true if this server believes it is the leader.

type RequestVoteArgs struct {
}
    example RequestVote RPC arguments structure.

type RequestVoteReply struct {
}
    example RequestVote RPC reply structure.


