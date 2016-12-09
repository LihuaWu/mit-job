# raft
--
    import "github.com/LihuaWu/mit-job/6.824/raft"


## Usage

    const Debug = 0

Debugging

#### func  DPrintf

    func DPrintf(format string, a ...interface{}) (n int, err error)


#### type AppendEntriesArgs

    type AppendEntriesArgs struct {
    }



#### type AppendEntriesReply

    type AppendEntriesReply struct {
    }



#### type ApplyMsg

    type ApplyMsg struct {
    	Index       int
    	Command     interface{}
    	UseSnapshot bool   // ignore for lab2; only used in lab3
    	Snapshot    []byte // ignore for lab2; only used in lab3
    }


as each Raft peer becomes aware that successive log entries are committed, the
peer should send an ApplyMsg to the service (or tester) on the same server, via
the applyCh passed to Make().

#### type Entry

    type Entry struct {
    }



#### type Persister

    type Persister struct {
    }



#### func  MakePersister

    func MakePersister() *Persister


#### func (*Persister) Copy

    func (ps *Persister) Copy() *Persister


#### func (*Persister) RaftStateSize

    func (ps *Persister) RaftStateSize() int


#### func (*Persister) ReadRaftState

    func (ps *Persister) ReadRaftState() []byte


#### func (*Persister) ReadSnapshot

    func (ps *Persister) ReadSnapshot() []byte


#### func (*Persister) SaveRaftState

    func (ps *Persister) SaveRaftState(data []byte)


#### func (*Persister) SaveSnapshot

    func (ps *Persister) SaveSnapshot(snapshot []byte)


#### type Raft

    type Raft struct {
    }


A Go object implementing a single Raft peer.

#### func  Make

    func Make(peers []*labrpc.ClientEnd, me int,
    	persister *Persister, applyCh chan ApplyMsg) *Raft

the service or tester wants to create a Raft server. the ports of all the Raft
servers (including this one) are in peers[]. this server's port is peers[me].
all the servers' peers[] arrays have the same order. persister is a place for
this server to save its persistent state, and also initially holds the most
recent saved state, if any. applyCh is a channel on which the tester or service
expects Raft to send ApplyMsg messages. Make() must return quickly, so it should
start goroutines for any long-running work.

#### func (*Raft) GetState

    func (rf *Raft) GetState() (int, bool)

return currentTerm and whether this server believes it is the leader.

#### func (*Raft) Kill

    func (rf *Raft) Kill()

the tester calls Kill() when a Raft instance won't be needed again. you are not
required to do anything in Kill(), but it might be convenient to (for example)
turn off debug output from this instance.

#### func (*Raft) RequestVote

    func (rf *Raft) RequestVote(args RequestVoteArgs, reply *RequestVoteReply)

example RequestVote RPC handler.

#### func (*Raft) Start

    func (rf *Raft) Start(command interface{}) (int, int, bool)

the service using Raft (e.g. a k/v server) wants to start agreement on the next
command to be appended to Raft's log. if this server isn't the leader, returns
false. otherwise start the agreement and return immediately. there is no
guarantee that this command will ever be committed to the Raft log, since the
leader may fail or lose an election.

the first return value is the index that the command will appear at if it's ever
committed. the second return value is the current term. the third return value
is true if this server believes it is the leader.

#### type RequestVoteArgs

    type RequestVoteArgs struct {
    }


example RequestVote RPC arguments structure.

#### type RequestVoteReply

    type RequestVoteReply struct {
    }


example RequestVote RPC reply structure.

#### type State

    type State struct {
    }
