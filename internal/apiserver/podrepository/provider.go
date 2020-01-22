package podrepository

type PodRepository struct {
}

func New() *PodRepository {
	return &PodRepository{}
}

func (pr *PodRepository) Get() string {
	return ""
}

func (pr *PodRepository) Add() string {
	return ""
}

func (pr *PodRepository) Remove() string {
	return ""
}

func (pr *PodRepository) Update() string {
	return ""
}
