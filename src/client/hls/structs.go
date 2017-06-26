package hls

import (
	"sync"
	"time"

	"fmt"

	"github.com/grafov/m3u8"
)

type HLSFetcher struct {
	URL            string
	TimeStart      time.Time
	MasterPlaylist m3u8.Playlist

	waitGroup *sync.WaitGroup
	stopChan  chan bool
}

func NewHLSFetcher(URL string) *HLSFetcher {
	s := &HLSFetcher{
		URL:       URL,
		TimeStart: time.Now(),
		stopChan:  make(chan bool),
		waitGroup: &sync.WaitGroup{},
	}
	return s
}

func (self *HLSFetcher) Start() {
	go self.doUpdateMasterPlaylist()
}

func (self *HLSFetcher) Stop() {
	fmt.Println("Closing stop chan")
	close(self.stopChan)

	fmt.Println("Waiting")
	self.waitGroup.Wait()

	fmt.Println("Done")
}

func (self *HLSFetcher) doUpdateMasterPlaylist() {
	self.waitGroup.Add(1)
	defer self.waitGroup.Done()

	timer := time.NewTimer(time.Second * 2) // ToDo: kill this magick timestamp
	defer timer.Stop()

	for {
		select {
		case <-self.stopChan:
			return
		case <-timer.C:
			fmt.Println("timer tick")
		default:

		}

	}
}
