package main

import (
	"strconv"
	"strings"
	"sync"
)

// function to process each line from a csv file
func processLine(line string, w *sync.WaitGroup) {
	unformatedLine := strings.Split(line, " ")
	req := Request{}
	req.Ip = unformatedLine[0]

	data := strings.Join(unformatedLine[3:5], " ")
	req.Data = strings.ReplaceAll(strings.ReplaceAll(data, "[", ""), "]", "")
	unformatedRequest := unformatedLine[5:]
	unformatedRequestStr := strings.Join(unformatedRequest, " ")
	httpData := strings.Split(unformatedRequestStr, "\"")
	req.Method = strings.Split(httpData[1], " ")[0]
	statusBytes := strings.Split(httpData[2], " ")
	req.Status = statusBytes[1]
	req.Bytes, _ = strconv.Atoi(statusBytes[2])
	req.Url = httpData[3]
	req.UserAgent = httpData[5]
	w.Done()

}
