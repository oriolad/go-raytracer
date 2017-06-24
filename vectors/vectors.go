package vectors

type Point struct{
	X, Y, Z float64
}

type Vector Point

type Line struct{
	P Point
	V Vector	
}

type Plane struct{
	N Vector
	V1 Vector
	V2 Vector
}
 
func (p1 Point) sub(p2 Point) Vector {
	return Vector{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
		Z: p1.Z - p2.Z,
	} 
}

func NewLine(p1 Point, p2 Point) Line{
	var vector = p1.sub(p2)
	return Line{
		P: p1,
		V: vector,
	}		
}

//func NewPlane(n Vector, v1 Vector, v2 Vector) Plane{
// }
