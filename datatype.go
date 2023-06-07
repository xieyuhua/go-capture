package main

import "time"

type IPStruct struct {
	InBytes    int       `json:"inbytes"`
	OutBytes   int       `json:"outbytes"`
	TotalBytes int       `json:"totalbytes"`
	Num        int       `json:"num"`
	Country    string    `json:"country"`
	City       string    `json:"city"`
	Longitude  float64   `json:"longitude"`
	Latitude   float64   `json:"latitude"`
	LastActive time.Time `json:"lastactive"`
}

// 配置选项
type Option struct {
	deviceName      string
	flushInterval   int
	ifWritePcap     int
	ifReverseResult bool
	pcapFilename    string
}

// 管道传输的流量统计信息结构
type BandwidthData struct {
	BandwidthStatisticStr string   `json:"bandwidthstr"`
	BandwidthList         PairList `json:"bandwidthlist"`
}

// 用于实现map的排序输出(先转为slice，并使用自定义个排序接口)
type Pair struct {
	Key   string    `json:"key"`
	Value *IPStruct `json:"value"`
}

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value.TotalBytes < p[j].Value.TotalBytes }

type PairList []Pair

