package main

type MovementDirector struct {
	Direction string
	Symbol    byte
	NextX     int
	NextY     int
}

func (m *MovementDirector) ChangeDirection() {
	switch m.Direction {
	case "UP":
		m.Direction = "RIGHT"
		m.Symbol = '>'
		m.NextX = 1
		m.NextY = 0
	case "RIGHT":
		m.Direction = "DOWN"
		m.Symbol = 'v'
		m.NextX = 0
		m.NextY = 1
	case "DOWN":
		m.Direction = "LEFT"
		m.Symbol = '<'
		m.NextX = -1
		m.NextY = 0
	case "LEFT":
		m.Direction = "UP"
		m.Symbol = '^'
		m.NextX = 0
		m.NextY = -1
	}
}

func (m *MovementDirector) GetNextCoordinates(x, y int) (int, int) {
	return (x + m.NextX), (y + m.NextY)
}

func InitMovementDirector() MovementDirector {
	return MovementDirector{
		Direction: "UP",
		Symbol:    '^',
		NextX:     0,
		NextY:     -1,
	}
}
