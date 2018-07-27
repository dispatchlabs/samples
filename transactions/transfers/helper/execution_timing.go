package helper

import (
	"github.com/dispatchlabs/disgo/commons/types"
	"time"
	"fmt"
	"sort"
)

var timingMap = map[string]*ExecutionTiming{}

type ExecutionTiming struct {
	CreateIndex     int
	Tx   			*types.Transaction
	Receipt 		*types.Receipt
	TxSubmitTime 	time.Time
	SubmitDelta     int64
	ExecuteDelta    int64
}

func NewExecutionTiming(index int, submitTime time.Time, tx *types.Transaction) *ExecutionTiming {
	return &ExecutionTiming{index,tx, nil, submitTime, 0, 0}
}

func (this *ExecutionTiming) GetTiming() string {
	return fmt.Sprintf("CreateIndex = %d Tx Create Time: %v Submit time: %v Execute time: %v",
		this.CreateIndex,
		this.Tx.ToTime().Format(time.StampMilli),
		this.TxSubmitTime.Format(time.StampMilli),
		this.Receipt.Created.Format(time.StampMilli),
	)
	//return fmt.Sprintf("CreateIndex = %s Tx Create Time: %v Submit time: %v SubmitDelta: %d Execute time: %v ExecuteDelta: %d",
	//	this.CreateIndex,
	//	this.Tx.ToTime(),
	//	this.TxSubmitTime,
	//	this.Receipt.Created,
	//	this.SubmitDelta,
	//	this.ExecuteDelta,
	//)
}

func AddTx(index int, tx *types.Transaction) {
	timingMap[tx.Hash] = NewExecutionTiming(index, time.Now(), tx)
}

func AddReceipt(hash string, receipt *types.Receipt) {
	timingMap[hash].Receipt = receipt
	//timingMap[hash].SubmitDelta = (timingMap[hash].TxSubmitTime.UnixNano() - timingMap[hash].Tx.Time*1000)
	//timingMap[hash].ExecuteDelta = (receipt.Created.UnixNano() - timingMap[hash].TxSubmitTime.UnixNano())
}

func GetTiming(hash string) *ExecutionTiming {
	return timingMap[hash]
}

func AddSubmitDelta(hash string, value int64) {
	timingMap[hash].SubmitDelta = value
}

func PrintTiming() {
	timingList := make([]*ExecutionTiming, 0)
	for _, v := range timingMap {
		timingList = append(timingList, v)
	}
	sorter := &ExecutionTimingSorter{timingList}
	sort.Sort(sorter)
	for _, timing := range sorter.Timings {
		fmt.Printf("%s\n", timing.GetTiming())
	}
}

// TransactionSorter joins a By function and a slice of Transaction to be sorted.
type ExecutionTimingSorter struct {
	Timings 	[]*ExecutionTiming
}

// Len is part of sort.Interface.
func (this *ExecutionTimingSorter) Len() int {
	return len(this.Timings)
}

// Swap is part of sort.Interface.
func (this *ExecutionTimingSorter) Swap(i, j int) {
	this.Timings[i], this.Timings[j] = this.Timings[j], this.Timings[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (this *ExecutionTimingSorter) Less(i, j int) bool {
	return this.Timings[i].Tx.Time < this.Timings[j].Tx.Time
}
