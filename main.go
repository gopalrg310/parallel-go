package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	log "github.com/sirupsen/logrus"
)

// main function
func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02T15:04:05.000-07:00"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
	requrls := []string{"https://facebook.com", "https://gmail.com", "https://instagram.com", "https://golang.com", "https://tutorialedge.net", "https://github.com/gopalrg310"}
	allfunc := make([]func(string), 0)
	for range requrls {
		allfunc = append(allfunc, doApiCall)
	}
	Parallelize(requrls, allfunc...)
}
//This is to do API Call
/*input
	- string
*/
func doApiCall(str string) {
	log.WithFields(log.Fields{}).Info(str)
	currenttime := time.Now()
	req, err := http.NewRequest("GET", str, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	log.Info("Completed ", str, " - ", time.Since(currenttime))
}
// This function will be used to execute functions parallely in GO
/*
	input-
		*urlstr - url to do call
		*function - function that defined
*/
func Parallelize(urlstr []string, functions ...func(string)) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))
	defer waitGroup.Wait()
	for i, function := range functions {
		go func(funct func(string), value string) {
			defer waitGroup.Done()
			funct(value)
		}(function, urlstr[i])
	}
}
