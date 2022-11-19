package broker

import (
	sm "github.com/lni/dragonboat/v4/statemachine"
	"github.com/tecbot/gorocksdb"
	"io"
	"sync/atomic"
	"unsafe"
)

var processedIndex = []byte("processedIndex")

type EventStateMachine struct {
	// 快照目录
	checkpointDir string
	shardID       uint64
	replicaID     uint64
	db            unsafe.Pointer
}

func (s *EventStateMachine) Open(stopc <-chan struct{}) (uint64, error) {
	//db, err := gorocksdb.OpenDb(gorocksdb.NewDefaultOptions(), "")
	//if err != nil {
	//	return 0, err
	//}
	//atomic.SwapPointer(&s.db, unsafe.Pointer(db))
	//slice, err := db.Get(gorocksdb.NewDefaultReadOptions(), processedIndex)
	//if slice.Size() == 0 {
	//	return 0, nil
	//}

	//return binary.LittleEndian.Uint64(slice.Data()), nil
	return 0, nil

}

func (s *EventStateMachine) Update(entries []sm.Entry) ([]sm.Entry, error) {

	println("update")
	//if s.closed {
	//	panic("update called after Close()")
	//}
	//db := (*gorocksdb.TransactionDB)(atomic.LoadPointer(&s.db))
	//
	//begin := db.TransactionBegin(gorocksdb.NewDefaultWriteOptions(), gorocksdb.NewDefaultTransactionOptions(), nil)
	//for i := range entries {
	//	entry := entries[i]
	//	var event protocol.Event
	//	err := proto.Unmarshal(entry.Cmd, &event)
	//	if err != nil {
	//		panic("unmarshal fail")
	//	}
	//	if jobCreate, ok := event.Value.(*protocol.Event_JobCreate); ok {
	//
	//		job := jobCreate.JobCreate
	//
	//		println("create", job.Name)
	//
	//		continue
	//	}
	//
	//	if _, ok := event.Value.(*protocol.Event_JobActive); ok {
	//
	//		continue
	//	}
	//	if _, ok := event.Value.(*protocol.Event_JobCompleted); ok {
	//
	//		continue
	//	}
	//
	//}

	//result := <-requestState.AppliedC()
	//if result.Committed() {
	//
	//}
	//begin.Commit()

	return entries, nil
}

func (s *EventStateMachine) Lookup(key interface{}) (interface{}, error) {
	//db := (*gorocksdb.TransactionDB)(atomic.LoadPointer(&s.db))
	//if db != nil {
	//	v, err := db.Get(gorocksdb.NewDefaultReadOptions(), key.([]byte))
	//	if err == nil && s.closed {
	//		panic("lookup returned valid result when EventStateMachine is already closed")
	//	}
	//	if err == pebble.ErrNotFound {
	//		return v, nil
	//	}
	//	return v, err
	//}
	//return nil, errors.New("db closed")
	return key, nil
}

func (s *EventStateMachine) Sync() error {
	return nil
}

func (s *EventStateMachine) PrepareSnapshot() (interface{}, error) {
	db := (*gorocksdb.TransactionDB)(atomic.LoadPointer(&s.db))
	checkpoint, err := db.NewCheckpoint()
	if err != nil {
		return nil, err
	}
	if err = checkpoint.CreateCheckpoint(s.checkpointDir, 64); err != nil {
		return nil, err
	}
	return checkpoint, nil

}

func (s *EventStateMachine) SaveSnapshot(checkpoint interface{}, writer io.Writer, notify <-chan struct{}) error {
	//TODO implement me
	panic("implement me")

}

func (s *EventStateMachine) RecoverFromSnapshot(reader io.Reader, i <-chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *EventStateMachine) Close() error {
	db := (*gorocksdb.TransactionDB)(atomic.LoadPointer(&s.db))
	db.Close()
	return nil
}
