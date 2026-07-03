package requests

import (
	"fmt"
	"folion-server/src/network/session"
	"folion-server/src/system/x"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type Response struct {
	Payload interface{}
}

type RequestContext struct {
	id        uuid.UUID
	request   Request
	timestamp time.Time
	session   session.Session
}

func NewRequestContext(_request Request, _session session.Session) *RequestContext {
	return &RequestContext{
		id:        uuid.New(),
		request:   _request,
		timestamp: time.Now(),
		session:   _session,
	}
}

func (requestContext *RequestContext) Respond(_response interface{}) error {
	err := requestContext.session.Write(_response)
	if err != nil {
		return x.ErrSessionRespondFailed.Message(
			fmt.Sprintf("Failed to send response to session %s", requestContext.session.ID.String())).Wrap(err)
	}

	return nil
}

type Worker struct {
	WorkerPool     chan chan Request
	RequestChannel chan Request
	quit           chan bool
}

func NewWorker(workerPool chan chan Request) *Worker {
	return &Worker{
		WorkerPool:     workerPool,
		RequestChannel: make(chan Request),
		quit:           make(chan bool),
	}
}

func (worker *Worker) Start() {
	go func() {
		slog.Debug("Worker started!")
		for {
			worker.WorkerPool <- worker.RequestChannel

			select {
			case request := <-worker.RequestChannel:
				slog.Debug(request.Type) // Process the request here
				return

			case <-worker.quit:
				return

			}
		}
	}()
}

func (worker *Worker) Stop() {
	go func() {
		worker.quit <- true
	}()
}

var RequestQueueChan chan Request

type RequestQueue struct {
	MaxWorkers int
	WorkerPool chan chan Request
}

func NewRequestQueue(maxWorkers int) *RequestQueue {
	workerPool := make(chan chan Request, maxWorkers)
	return &RequestQueue{
		MaxWorkers: maxWorkers,
		WorkerPool: workerPool,
	}
}

func (requestQueue *RequestQueue) Run() {
	slog.Debug("Starting Request queue")
	for i := 0; i < requestQueue.MaxWorkers; i++ {
		worker := NewWorker(requestQueue.WorkerPool)
		slog.Debug(fmt.Sprintf("Created worker %d", i))
		worker.Start()
	}

	go requestQueue.process()
}

func (requestQueue RequestQueue) process() {
	for {
		slog.Debug("Started processing Request queue")
		select {
		case request := <-RequestQueueChan:
			go func(request Request) {
				requestChannel := <-requestQueue.WorkerPool
				requestChannel <- request
			}(request)
		}
	}
}
