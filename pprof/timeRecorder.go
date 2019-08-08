package pprof

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

type TimeRecorder struct {
	mutex sync.RWMutex
	records map[string]*timeRecord
}

type timeRecord struct {
	Times int64
	TotalUsedTime int64
	MaxUsedTime   int64
	MinUsedTime   int64
}

type avgTimeRecord struct {
	Name          string
	Times         int64
	AvgUsedTime   int64
	MinUsedTime   int64
	MaxUsedTime   int64
	TotalUsedTime int64
}
type avgTimeRecords []*avgTimeRecord

func (this avgTimeRecords) Len() int {
	return len(this)
}

func (this avgTimeRecords) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this avgTimeRecords) Less(i, j int) bool {
	return this[i].AvgUsedTime > this[j].AvgUsedTime || (this[i].AvgUsedTime == this[j].AvgUsedTime && this[i].Times < this[j].Times)
}

func NewTimeRecorder()*TimeRecorder{
	return &TimeRecorder{
		records:make(map[string]*timeRecord),
	}
}

func (tr *TimeRecorder) getRecord(name string)*timeRecord{
	tr.mutex.Lock()
	defer tr.mutex.Unlock()
	r,exists := tr.records[name]
	if !exists {
		r = new(timeRecord)
		tr.records[name] = r
	}
	return r
}

func (tr *TimeRecorder) getRecords() avgTimeRecords{
	tr.mutex.RLock()
	defer tr.mutex.RUnlock()
	results := make(avgTimeRecords, 0, len(tr.records))

	for name, d := range tr.records {
		results = append(results, &avgTimeRecord{
			Name:          name,
			Times:         d.Times,
			AvgUsedTime:   d.TotalUsedTime / d.Times,
			MaxUsedTime:   d.MaxUsedTime,
			MinUsedTime:   d.MinUsedTime,
			TotalUsedTime: d.TotalUsedTime,
		})
	}

	return results
}

func (tr *TimeRecorder) Record(name string,elapsed time.Duration){
	r := tr.getRecord(name)
	usedNano := elapsed.Nanoseconds()
	atomic.AddInt64(&r.Times, 1)
	atomic.AddInt64(&r.TotalUsedTime, usedNano)

	for {
		old := atomic.LoadInt64(&r.MaxUsedTime)
		if old >= usedNano || atomic.CompareAndSwapInt64(&r.MaxUsedTime, old, usedNano) {
			break
		}
	}

	for {
		old := atomic.LoadInt64(&r.MinUsedTime)
		if old <= usedNano || atomic.CompareAndSwapInt64(&r.MinUsedTime, old, usedNano) {
			break
		}
	}
}

func (tr *TimeRecorder) WriteCSV(writer io.Writer) error {
	results:= tr.getRecords()
	sort.Sort(results)
	buf := bufio.NewWriter(writer)
	if _,err := fmt.Fprintln(writer,"name,times,avg,min,max,total");err!=nil{
		return err
	}

	for _, r := range results {
		if _, err := fmt.Fprintf(writer,
			"%s,%d,%dns,%dns,%dns,%d nanosecond\n",
			r.Name,
			r.Times,
			r.AvgUsedTime,
			r.MinUsedTime,
			r.MaxUsedTime,
			r.TotalUsedTime,
		); err != nil {
			return err
		}
	}

	return buf.Flush()
}

func (tr *TimeRecorder) SaveCSV(filename string) error{
	file,err := os.Create(filename)
	if err != nil{
		return err
	}
	defer file.Close()
	return tr.WriteCSV(file)
}