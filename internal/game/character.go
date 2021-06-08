package game

type character struct {
	id       int
	userID   int
	name     string
	position position
}

func (this character) GetID() int {
	return this.id
}

func (this character) GetName() string {
	return this.name
}

func (this character) GetPosition() position {
	return this.position
}

type position struct {
	x int
	y int
}

func (this position) GetX() int {
	return this.x
}

func (this position) GetY() int {
	return this.y
}
