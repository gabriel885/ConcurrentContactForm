package main

import (
	"net/url"
	"strconv"
	"sync"
	"testing"
)

func testEmailRequest(i int, t *testing.T) {
	conf := LoadSmtpConfigurations("config.json")
	t.Logf("Go routine num. %v \n", strconv.Itoa(i))
	conf.HandleSendMail(url.Values{"name": {"Client"}, "email": {"client@example.com"}, "subject": {"Subject " + strconv.Itoa(i)}, "message": {"Test message "}})
}

func TestNonConcurrentSendEmail(t *testing.T) {

	for i := 0; i < 5; i++ {
		conf := LoadSmtpConfigurations("config.json")
		t.Logf("Go routine num. %v \n", strconv.Itoa(i))
		conf.HandleSendMail(url.Values{"name": {"Client"}, "email": {"client@example.com"}, "subject": {"Subject " + strconv.Itoa(i)}, "message": {"Test message "}})
	}
}

func TestConcurrentSendEmail(t *testing.T) {

	const nThreads = 5 // maximum go routines

	var ch = make(chan int, 5)
	var wg sync.WaitGroup

	wg.Add(nThreads) // add max num threads
	// start 6 go routins
	for i := 0; i < nThreads; i++ {
		go func() {
			for {
				i, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				testEmailRequest(i, t)
			}
		}()
	}

	// Now the jobs can be added to the channel, which is used as a queue
	for i := 0; i < 5; i++ {
		ch <- i // add i to the queue
	}
	close(ch) // close all go routines
	wg.Wait()

}
