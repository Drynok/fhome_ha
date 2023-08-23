package workerpool

import (
	"errors"
	"log"
	"time"
)

type Worker struct {
	ID                int `json:"id"`
	ProcessedMessages int `json:"processedMessages"`
	Task              TaskHandler
	stopChan          chan struct{} `json:"-"`
}

// Start worker.
func (w *Worker) Start(dataChan <-chan string) {
	for {
		select {
		case <-w.stopChan:
			log.Printf("Worker %d stopped", w.ID)
			return
		case data := <-dataChan:
			w.Task(data)
			log.Printf("%v worker finished with his task", w.ID)
		case <-time.After(2 * time.Minute):
			log.Printf("Worker %d reached the maximum lifetime of 2 minutes", w.ID)
			return
		}
	}
}

// Stop worker.
func (w *Worker) Stop() {
	close(w.stopChan)
}

func NewWorker(id int) *Worker {
	return &Worker{
		ID:                id,
		ProcessedMessages: 0,
		stopChan:          make(chan struct{}),
	}
}

// TaskHandler function callbacks
type TaskHandler func(...interface{}) (interface{}, error)

type WorkerPool struct {
	workersCount int
	Workers      []Worker `json:"" binding:"dive"`
	task         chan TaskHandler
	timeout      time.Duration // max timeout
}

// GetStats returns current status of the worker pool.
func (p *WorkerPool) GetStats() []Worker {
	return p.Workers
}

func (p *WorkerPool) Submit(tsk TaskHandler) {

}

func (p *WorkerPool) Start() {
}

func (p *WorkerPool) Stop() {
}

func NewWorkerPool(initialWorkers int, maxOperationsInQueue int) (*WorkerPool, error) {
	// wrong data entry
	if initialWorkers < 0 || maxOperationsInQueue <= 0 {
		log.Print("not able to init worker pool")
		return nil, errors.New("not able to init worker pool")
	}

	ret := &WorkerPool{}

	return ret, nil
}
