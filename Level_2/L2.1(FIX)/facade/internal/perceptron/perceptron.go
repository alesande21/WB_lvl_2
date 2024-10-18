package perceptron

type BaseModel interface {
	init()
	setInput()
	forwardFeed()
	backPropogation()
	weightUpdate()
}

type PerceptronFacade struct {
	baseModel BaseModel
}

func GetPerceptronFacade(model string) *PerceptronFacade {
	var bm BaseModel
	if model == "graph" {
		bm = GetGraphModel()
	} else {
		bm = GetMatrixModel()
	}
	pf := &PerceptronFacade{baseModel: bm}
	pf.baseModel.init()
	return pf
}

func (pf *PerceptronFacade) SwitchPerceptron(model string) {
	var bm BaseModel
	if model == "graph" {
		bm = GetGraphModel()
	} else {
		bm = GetMatrixModel()
	}
	pf.baseModel = bm
	pf.baseModel.init()
}

func (pf *PerceptronFacade) Run() {
	pf.baseModel.setInput()
	pf.baseModel.forwardFeed()
	pf.baseModel.backPropogation()
	pf.baseModel.weightUpdate()
}
