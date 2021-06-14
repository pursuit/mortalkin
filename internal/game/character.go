package game

type Character struct {
	ID       int
	UserID   int
	Name     string
	Position Position
}

func (this Character) GetID() int {
	return this.ID
}

func (this Character) GetName() string {
	return this.Name
}

func (this Character) GetPosition() Position {
	return this.Position
}

type Position struct {
	X int
	Y int
}

func (this Position) GetX() int {
	return this.X
}

func (this Position) GetY() int {
	return this.Y
}
