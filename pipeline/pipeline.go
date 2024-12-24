package pipeline

// Stage defines a single processing step in the pipeline.
type Stage interface {
	Process(data map[string]interface{}) (map[string]interface{}, error)
}

// Pipeline orchestrates multiple stages.
type Pipeline struct {
	stages []Stage
}

// AddStage adds a new stage to the pipeline.
func (p *Pipeline) AddStage(stage Stage) {
	p.stages = append(p.stages, stage)
}

// Execute runs the pipeline on the input data.
func (p *Pipeline) Execute(data map[string]interface{}) (map[string]interface{}, error) {
	var err error
	for _, stage := range p.stages {
		data, err = stage.Process(data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

// NewPipeline creates a new empty pipeline.
func NewPipeline() *Pipeline {
	return &Pipeline{}
}
