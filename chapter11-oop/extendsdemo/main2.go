package main

import "fmt"

type Student struct {
	Name  string
	Score float64
}

type Pupil struct {
	Student
}

type Graduate struct {
	Student
}

func (s *Student) SetScore(score float64) {
	s.Score = score
}

func (s *Student) ShowInfo() {
	fmt.Println("ShowInfo:", s)
}

func (p *Graduate) GraduateBook(book string) {
	fmt.Println("GraduateBook:", book)
}

func main() {
	// 小学生
	p := &Pupil{}
	p.Name = "小学生"
	p.Score = 10

	fmt.Println(p)

	p.SetScore(90)
	p.ShowInfo()
	fmt.Println()

	// 大学生
	g := &Graduate{}
	g.Name = "da学生"
	g.Score = 90
	fmt.Println(g)
	g.GraduateBook("高数")
}
