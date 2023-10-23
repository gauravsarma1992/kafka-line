package pipeline

type (
	WorkerStatusT uint8

	PollerManager struct {
		Pollers []*Poller `json:"pollers"`
	}

	Poller struct {
		Pipeline *Pipeline `json:"pipeline"`
		Workers  []*Worker `json:"worker"`
	}

	Worker struct {
		Uuid   UuidT         `json:"uuid"`
		Status WorkerStatusT `json:"status"`
	}
)

func NewPollerMgr() (pMgr *PollerManager, err error) {
	pMgr = &PollerManager{}
	return
}

func (pMgr *PollerManager) Run() (err error) {
	pipelines := []*Pipeline{}
	for _, pipeline := range pipelines {
		poller, _ := NewPoller(pipeline)
		go poller.Run()
	}
	return
}

func NewPoller(pipeline *Pipeline) (poller *Poller, err error) {
	poller = &Poller{
		Pipeline: pipeline,
	}
	return
}

func (poller *Poller) Run() (err error) {
	return
}
