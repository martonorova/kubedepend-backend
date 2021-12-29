package collector

import (
	"log"

	"github.com/martonorova/kubedepend-backend/pkg/services"
)

type defaultCollectorService struct {
	jobService services.JobService
	ResultC    chan services.JobResultDTO
	QuitC      chan bool
}

func NewDefaultCollectorService(jobService services.JobService) services.CollectorService {

	dcs := &defaultCollectorService{
		jobService: jobService,
		ResultC:    make(chan services.JobResultDTO, 100),
	}

	return dcs
}

func (cs *defaultCollectorService) CollectJob(jobResult *services.JobResultDTO) {
	go func() {
		cs.ResultC <- *jobResult
	}()
}

func (cs *defaultCollectorService) Start() {
	cs.start()
}

func (cs *defaultCollectorService) start() {
	go func() {
		for {
			select {
			case result := <-cs.ResultC:
				log.Printf("Collected Job result ID: %d, Result: %d\n", result.ID, result.Result)

			case <-cs.QuitC:
				log.Println("Collector stopped")
				return
			}
		}
	}()
}

func (cs *defaultCollectorService) Stop() {
	go func() {
		cs.QuitC <- true
	}()
}
