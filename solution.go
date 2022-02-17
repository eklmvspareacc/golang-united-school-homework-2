package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sidesQuantity int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
const (
	SidesCircle   sidesQuantity = 0
	SidesTriangle sidesQuantity = 3
	SidesSquare   sidesQuantity = 4
)

func CalcSquare(sideLen float64, sidesNum sidesQuantity) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	case SidesTriangle:
		return (math.Sqrt(3) / 4) * sideLen * sideLen
	case SidesSquare:
		return sideLen * sideLen
	default:
		return 0
	}
}
