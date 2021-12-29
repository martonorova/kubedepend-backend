package execution

import (
	"fmt"
	"log"

	"github.com/martonorova/kubedepend-backend/pkg/services"
)

type defaultExecutionService struct {
	dispatcher *Dispatcher
	collector  services.CollectorService
}

func NewDefaultExecutionService(nworkers uint64, jobQueueSize uint64, collector services.CollectorService) services.ExecutionService {

	dispatcher := NewDispatcher(nworkers, jobQueueSize)

	return &defaultExecutionService{
		dispatcher: dispatcher,
		collector:  collector,
	}
}

func (ws *defaultExecutionService) SubmitJob(jobSubmit *services.SubmitJobDTO) error {
	ws.dispatcher.Submit(*jobSubmit)
	return nil
}

func (ws *defaultExecutionService) Start() {
	ws.collector.Start()
	ws.dispatcher.Start(ws.collector)
}

// register itself to WorkerQueue (pool) and execute Jobs
type Worker struct {
	ID          uint64
	JobC        chan services.SubmitJobDTO
	WorkerQueue chan chan services.SubmitJobDTO
	//ResultQueue *chan dto.JobResultDTO
	collector services.CollectorService
	QuitC     chan bool
}

// receives jobsubmits, forwards it to next available worker
type Dispatcher struct {
	JobQueue     chan services.SubmitJobDTO
	WorkerQueue  chan chan services.SubmitJobDTO
	nworkers     uint64
	jobQueueSize uint64
	workers      map[uint64]*Worker
}

func NewWorker(id uint64, workerQueue chan chan services.SubmitJobDTO, collector services.CollectorService) *Worker {
	worker := &Worker{
		ID:          id,
		JobC:        make(chan services.SubmitJobDTO),
		WorkerQueue: workerQueue,
		collector:   collector,
		QuitC:       make(chan bool),
	}

	return worker
}

func NewDispatcher(nworkers uint64, jobQueueSize uint64) *Dispatcher {
	dispatcher := &Dispatcher{
		JobQueue:     make(chan services.SubmitJobDTO, jobQueueSize),
		WorkerQueue:  make(chan chan services.SubmitJobDTO, nworkers),
		nworkers:     nworkers,
		jobQueueSize: jobQueueSize,
		workers:      make(map[uint64]*Worker),
	}

	return dispatcher
}

func (w *Worker) Start() {
	go func() {
		for {
			// add worker to worker queue
			w.WorkerQueue <- w.JobC

			select {
			case job := <-w.JobC:
				log.Printf("worker%d: Received job request with ID: %d\n", w.ID, job.ID)

				result := fibonacci(job.Input)

				// TODO send to collector
				go func() {
					w.collector.CollectJob(&services.JobResultDTO{ID: job.ID, Result: result})
				}()

				log.Printf("worker%d: Finished job with ID: %d \n", w.ID, job.ID)

			case <-w.QuitC:
				log.Printf("worker%d: Stopping\n", w.ID)
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

func (d *Dispatcher) Submit(jobSubmit services.SubmitJobDTO) {

	// send jobsubmit in goroutine in case job queue is full
	go func() {
		d.JobQueue <- jobSubmit
	}()
}

func (d *Dispatcher) Start(collector services.CollectorService) {
	// Create workers
	// TODO start with a few workers, and increase the pool if needed
	fmt.Println("Starting dispatcher...")
	for i := 0; i < int(d.nworkers); i++ {
		worker := NewWorker(uint64(i+1), d.WorkerQueue, collector)

		// register worker in dispatcher
		d.workers[worker.ID] = worker

		// start worker
		worker.Start()
	}

	// listen for job submits and forward to workers
	go func() {
		for {
			select {
			case jobSubmit := <-d.JobQueue:
				log.Printf("Received job submit with ID: %d and Input: %d\n", jobSubmit.ID, jobSubmit.Input)
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
