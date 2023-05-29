package main

import (
	"log"
	"os"
	"os/exec"
// 	"fmt"
	"runtime"
	"sort"
)

func handleErr(err error, info string) {
	if err != nil {
		log.Panic(info, err.Error())
	}
}

func logErr(err error, info string) {
	if err != nil {
		log.Println(info, err.Error())
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if runtime.GOOS == "linux" {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}


func sortIPs(bandwidthMap map[string]*IPStruct, sortvale int) PairList {
	pl := make(PairList, len(bandwidthMap))
	i := 0
	for k, v := range bandwidthMap {
	   // fmt.Printf("%v,", v.OutBytes)
// 		pl[v.OutBytes] = Pair{k, v}
		pl[i] = Pair{k, v}
		i++
	}
// 	sort.Sort(sort.Reverse(pl))
	//排序
	switch sortvale {
	    case 1: 
            sort.Slice(pl, func(i, j int) bool {
                return pl[i].Value.OutBytes > pl[j].Value.OutBytes
            })
	    break
	    case 2: 
            sort.Slice(pl, func(i, j int) bool {
                return pl[i].Value.InBytes > pl[j].Value.InBytes
            })
	    break
	    case 3: 
            sort.Slice(pl, func(i, j int) bool {
                return pl[i].Value.TotalBytes > pl[j].Value.TotalBytes
            })
	    break
	    default:
	    sort.Sort(sort.Reverse(pl))
	}
	
	// sort.Sort(pl)
	return pl
}
