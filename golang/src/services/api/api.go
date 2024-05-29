package api

type Api struct {
	repo       *any
	httpserver *any
}

func (a *Api) Start() {

	// <-a.feeder.Chan

	// TODO :
}

func NewApi() *Api {
	return &Api{
		nil, // new
		nil, // new
	}
}
