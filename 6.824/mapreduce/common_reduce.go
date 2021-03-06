package mapreduce

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

// doReduce does the job of a reduce worker: it reads the intermediate
// key/value pairs (produced by the map phase) for this task, sorts the
// intermediate key/value pairs by key, calls the user-defined reduce function
// (reduceF) for each key, and writes the output to disk.
func doReduce(
	jobName string, // the name of the whole MapReduce job
	reduceTaskNumber int, // which reduce task this is
	nMap int, // the number of map tasks that were run ("M" in the paper)
	reduceF func(key string, values []string) string,
) {

	reduceMap := make(map[string][]string)

	var fileIdx []*os.File
	var fw []*json.Decoder
	for i := 0; i < nMap; i++ {
		f, err := os.Open(reduceName(jobName, i, reduceTaskNumber))
		if err != nil {
			log.Fatal("open file err: ", err)
		}
		fileIdx = append(fileIdx, f)
		fw = append(fw, json.NewDecoder(f))
	}

	defer func() {
		for i := 0; i < nMap; i++ {
			fmt.Println(fileIdx[i].Name())
			fileIdx[i].Close()
		}
	}()

	for _, dec := range fw {
		for {
			var item KeyValue
			if err := dec.Decode(&item); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal("json decode err: ", err)
			}

			if _, ok := reduceMap[item.Key]; !ok {
				reduceMap[item.Key] = make([]string, 0)
			}

			reduceMap[item.Key] = append(reduceMap[item.Key], item.Value)
		}
	}

	var keys []string

	for key := range reduceMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	log.Println("sorting keys")

	mf, _ := os.Create(mergeName(jobName, reduceTaskNumber))
	defer mf.Close()
	enc := json.NewEncoder(mf)

	for _, key := range keys {
		err := enc.Encode(&KeyValue{key, reduceF(key, reduceMap[key])})
		if err != nil {
			fmt.Println("err:", err)
		}
	}
	// TODO:
	// You will need to write this function.
	// You can find the intermediate file for this reduce task from map task number
	// m using reduceName(jobName, m, reduceTaskNumber).
	// Remember that you've encoded the values in the intermediate files, so you
	// will need to decode them. If you chose to use JSON, you can read out
	// multiple decoded values by creating a decoder, and then repeatedly calling
	// .Decode() on it until Decode() returns an error.
	//
	// You should write the reduced output in as JSON encoded KeyValue
	// objects to a file named mergeName(jobName, reduceTaskNumber). We require
	// you to use JSON here because that is what the merger than combines the
	// output from all the reduce tasks expects. There is nothing "special" about
	// JSON -- it is just the marshalling format we chose to use. It will look
	// something like this:
	//
	// enc := json.NewEncoder(mergeFile)
	// for key in ... {
	// 	enc.Encode(KeyValue{key, reduceF(...)})
	// }
	// file.Close()
}
