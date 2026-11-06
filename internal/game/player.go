package game

type Player struct {
	id    int
	name  string
	cards Cards
}

type Cards []Card
