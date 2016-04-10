package chief

import (
	"testing"
	"time"
)

func Test_StartStopWorker(t *testing.T) {
	done := make(chan bool)
	fn := func(_ Job) {
		done <- true
	}

	pool := make(PoolChannel, 1)
	w := newWorker(pool, fn)
	w.Start()

	jobs := <-pool
	jobs <- Job{Order: true}

	<-done

	w.Stop()
}

func Test_StartSendJobStopChief(t *testing.T) {
	done := make(chan bool)
	fn := func(_ Job) {
		done <- true
	}

	c := New(1, fn)
	c.Start()

	c.Jobs <- Job{Order: true}

	<-done

	c.Stop()
}

func Test_StartStopChief(t *testing.T) {
	fn := func(_ Job) {
	}

	c := New(1, fn)
	c.Start()
	c.Stop()
}

func Test_StartManyWorkersSendJobsStopChief(t *testing.T) {
	fn := func(_ Job) {
	}

	c := New(10, fn)
	c.Start()
	c.Jobs <- Job{Order: true}
	c.Jobs <- Job{Order: true}
	c.Jobs <- Job{Order: true}
	c.Stop()
}

func Test_StartStopChiefZero(t *testing.T) {
	fn := func(_ Job) {
	}

	c := New(0, fn)
	c.Start()
	c.Jobs <- Job{Order: true}
	c.Stop()
	// We have to wait that we get the 100% coverage :)
	// We have to wait that the goroutine in ctrl func
	// receive the close pool signal
	time.Sleep(1 * time.Second)
}
