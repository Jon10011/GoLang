package model

type student struct {
	Name  string
	score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

//student 中score首字母小写，则在其他包中无法引用
//解决
func (s *student) getScore() float64 {
	return s.score
}
