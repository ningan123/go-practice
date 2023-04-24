package model

type student struct {
	Name  string
	score float64
}

func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		score: score,
	}
}

func (s *student) GetScore() float64 {
	return s.score // 这个地方可以直接访问，因为在本包
}
