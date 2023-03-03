package providers

import (
	queuecontract "gopkg.in/go-mixed/framework.v1/contracts/queue"
	"gopkg.in/go-mixed/framework.v1/facades/queue"
)

type QueueServiceProvider struct {
}

func (sp *QueueServiceProvider) Register() {

	queue.Register(sp.Jobs()...)
}

func (sp *QueueServiceProvider) Boot() {

}

func (sp *QueueServiceProvider) Jobs() []queuecontract.IJob {
	return []queuecontract.IJob{}
}
