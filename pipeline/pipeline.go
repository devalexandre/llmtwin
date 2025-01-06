package pipeline

type Stage interface {
	ShouldExecute(data map[string]interface{}) bool
	Process(data map[string]interface{}) (map[string]interface{}, error)
}

type Pipeline struct {
	stages []Stage
}

func (p *Pipeline) AddStage(stage Stage) {
	p.stages = append(p.stages, stage)
}

func (p *Pipeline) Execute(data map[string]interface{}) (map[string]interface{}, error) {
	var err error
	for _, stage := range p.stages {
		if stage.ShouldExecute(data) {
			data, err = stage.Process(data)
			if err != nil {
				return nil, err
			}
		}
	}
	return data, nil
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}
