package day2

type Direction = int

const (
	Forward Direction = iota
	Up      Direction = 2
	Down    Direction = 3
)

type Command struct {
	Direction Direction
	Length    int
}

type Location struct {
	Horizontal int
	Depth      int
}

type LocationWithAim struct {
	Horizontal int
	Depth      int
	Aim        int
}
