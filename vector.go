package ebitil

import (
	"math"
)

type Vector2D struct {
	X, Y float64
}

func (v *Vector2D) Add(v2 *Vector2D) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vector2D) Sub(v2 *Vector2D) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vector2D) Dot(v2 *Vector2D) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v *Vector2D) Cross(v2 *Vector2D) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v *Vector2D) Mul(v2 *Vector2D) {
	v.X *= v2.X
	v.Y *= v2.Y
}

func (v *Vector2D) Div(v2 *Vector2D) {
	v.X /= v2.X
	v.Y /= v2.Y
}

func (v *Vector2D) DivScalar(v2 *Vector2D) {
	v.X /= v2.X
	v.Y /= v2.Y
}

func (v *Vector2D) MulScalar(v2 *Vector2D) {
	v.X *= v2.X
	v.Y *= v2.Y
}

func (v *Vector2D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vector2D) Normalized() Vector2D {
	mag := v.Magnitude()
	return Vector2D{v.X / mag, v.Y / mag}
}
