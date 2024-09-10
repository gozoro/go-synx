package synx

import (
	"sync"
	"time"
)

// The structure sync.WaitGroup is extended with WaitWithTimeout() function.
type WaitGroup struct {
	sync.WaitGroup
}

// WaitWithTimeout returns the value "true" when  the [WaitGroup] counter is zero.
// And returns the value "false" when the wait is completed by timeout.
//
//	func shutdownServices() {
//		wg := &synx.WaitGroup{}
//		wg.Add(1)
//		shutdownService1(wg)
//
//		wg.Add(1)
//		shutdownService2(wg)
//
//		if ok := wg.WaitWithTimeout(10 * time.Second); !ok {
//			log.Println("Force shutdown by timeout")
//		}
//		log.Println("Shutdown finished")
//		os.Exit(0)
//	}
func (wg *WaitGroup) WaitWithTimeout(timeout time.Duration) bool {

	timeoutChan := time.After(timeout)
	waitChan := make(chan struct{})

	go func() {
		wg.Wait()
		close(waitChan)
	}()

	select {
	case <-timeoutChan:
		return false
	case <-waitChan:
		return true
	}

}
