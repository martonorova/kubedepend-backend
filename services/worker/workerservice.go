package worker

import (
	"fmt"

	"github.com/martonorova/kubedepend-backend/dto"
)

// collect results from Workers
type Collector struct {
	ResultC chan dto.JobResultDTO
}

// register itself to WorkerQueue (pool) and execute Jobs
type Worker struct {
	ID          uint64
	JobC        chan dto.SubmitJobDTO
	WorkerQueue chan chan dto.SubmitJobDTO
	ResultQueue chan dto.JobResultDTO
	QuitC       chan bool
}

// receives jobsubmits, forwards it to next available worker
type Dispatcher struct {
	JobQueue     chan dto.SubmitJobDTO
	WorkerQueue  chan chan dto.SubmitJobDTO
	nworkers     uint64
	jobQueueSize uint64
	workers      map[uint64]*Worker
}

func NewCollector(resultQueueSize uint64) *Collector {
	return &Collector{
		ResultC: make(chan dto.JobResultDTO, resultQueueSize),
	}
}

func NewWorker(id uint64, workerQueue chan chan dto.SubmitJobDTO, resultQueue chan dto.JobResultDTO) *Worker {
	worker := &Worker{
		ID:          id,
		JobC:        make(chan dto.SubmitJobDTO),
		WorkerQueue: workerQueue,
		ResultQueue: resultQueue,
		QuitC:       make(chan bool),
	}

	return worker
}

func NewDispatcher(nworkers uint64, jobQueueSize uint64) *Dispatcher {
	dispatcher := &Dispatcher{
		JobQueue:     make(chan dto.SubmitJobDTO, jobQueueSize),
		WorkerQueue:  make(chan chan dto.SubmitJobDTO, nworkers),
		nworkers:     nworkers,
		jobQueueSize: jobQueueSize,
		workers:      make(map[uint64]*Worker),
	}

	return dispatcher
}

func (c *Collector) Start() {
	go func() {
		for {
			select {
			case result := <-c.ResultC:
				fmt.Printf("Collected Job result ID: %d, Result: %d", result.ID, result.Result)
			}
		}
	}()
}

func (w *Worker) Start() {
	go func() {
		for {
			// add worker to worker queue
			w.WorkerQueue <- w.JobC

			select {
			case job := <-w.JobC:
				fmt.Printf("worker%d: Received job request with ID: %d\n", w.ID, job.ID)

				result := fibonacci(job.Input)

				// TODO send to collector
				w.ResultQueue <- dto.JobResultDTO{ID: job.ID, Result: result}

				fmt.Printf("worker%d: Finished job with ID: %d \n", w.ID, job.ID)

			case <-w.QuitC:
				fmt.Printf("worker%d: Stopping\n", w.ID)
			}
		}
	}()
}

func (w *Worker) Stop() {
	// do not want to block
	// worker will only stop after it finished its current job
	go func() {
		w.QuitC <- true
	}()
}

func (d *Dispatcher) Submit(jobSubmit dto.SubmitJobDTO) error {

	// send jobsubmit in goroutine in case job queue is full
	go func() {
		d.JobQueue <- jobSubmit
	}()
	return nil
}

func (d *Dispatcher) Start(collector *Collector) {
	// Create workers
	// TODO start with a few workers, and increase the pool if needed
	fmt.Println("Starting dispatcher...")
	for i := 0; i < int(d.nworkers); i++ {
		worker := NewWorker(uint64(i+1), d.WorkerQueue, collector.ResultC)

		// register worker in dispatcher
		d.workers[worker.ID] = worker

		// start worker
		worker.Start()
	}

	// Start collector
	fmt.Println("Start collector")
	collector.Start()

	// listen for job submits and forward to workers
	go func() {
		for {
			select {
			case jobSubmit := <-d.JobQueue:
				fmt.Printf("Received job submit with ID: %d and Input: %d\n", jobSubmit.ID, jobSubmit.Input)
				go func() {
					worker := <-d.WorkerQueue
					worker <- jobSubmit
				}()
			}
		}
	}()
}

func (d *Dispatcher) StopWorkers() {
	fmt.Println("Stopping workers...")
	for _, worker := range d.workers {
		worker.Stop()
	}
}

func fibonacci(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
