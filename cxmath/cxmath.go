package cxmath

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"

	"github.com/skycoin/cx-game/cxmath/math32"
)

func ConvertScreenCoordsToWorld(x, y float32, projection mgl32.Mat4) mgl32.Vec2 {
	homogenousClipCoords := mgl32.Vec4{x, y, -1.0, 1.0}
	cameraCoords := projection.Inv().Mul4x1(homogenousClipCoords)
	return cameraCoords.Vec2()
}

func Scale(factor float32) mgl32.Mat4 {
	return mgl32.Scale3D(factor, factor, factor)
}

func atan2f32(y, x float32) float32 {
	return float32(math.Atan2(float64(y), float64(x)))
}

func AngleTo(v1, v2 mgl32.Vec2) float32 {
	return atan2f32(v1.Y(), v1.X()) - atan2f32(v2.Y(), v2.X())
}

// non negative solution to x % d
func PositiveModulo(x, b int) int {
	x = x % b
	if x >= 0 {
		return x
	}
	return x + b
}

func Sign(value float32) float32 {
	if value < 0 {
		return -1
	} else {
		return 1
	}
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Abs(a float32) float32 {
	if a >= 0 {
		return a
	}
	return a * -1
}

func Sqrt(a float32) float32 {
	return float32(math.Sqrt(float64(a)))
}

func Floor(a float32) float32 {
	return float32(math.Floor(float64(a)))
}

func LerpVec2(v1, v2 mgl32.Vec2, alpha float32) mgl32.Vec2 {
	beta := 1 - alpha
	return v1.Mul(beta).Add(v2.Mul(alpha))
}

func RoundVec2(v1 mgl32.Vec2) (x, y int32) {
	x = int32(float32(math.Round(float64(v1.X()))))
	y = int32(float32(math.Round(float64(v1.Y()))))
	return
}

// linearly interpolate between start and target by a factor of alpha
func Lerp(start, target, alpha float32) float32 {
	alpha = math32.Clamp(alpha, 0, 1)
	return (alpha * target) + ((1 - alpha) * start)
}

func ClampF(x, min,max float32) float32 {
	return math32.Clamp(x,min,max)
}

func DegToRad(x float32) float32 {
	return x*math.Pi/180
}
