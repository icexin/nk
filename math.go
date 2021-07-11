package nk

import "math"

func X__builtin_isnan(x float64) int32    { return Bool32(math.IsNaN(x)) }
func Xacos(x float64) float64             { return math.Acos(x) }
func Xacosh(x float64) float64            { return math.Acosh(x) }
func Xasin(x float64) float64             { return math.Asin(x) }
func Xasinh(x float64) float64            { return math.Asinh(x) }
func Xatan(x float64) float64             { return math.Atan(x) }
func Xatan2(x, y float64) float64         { return math.Atan2(x, y) }
func Xatanh(x float64) float64            { return math.Atanh(x) }
func Xceil(x float64) float64             { return math.Ceil(x) }
func Xceilf(x float32) float32            { return float32(math.Ceil(float64(x))) }
func Xcopysign(x, y float64) float64      { return math.Copysign(x, y) }
func Xcopysignf(x, y float32) float32     { return float32(math.Copysign(float64(x), float64(y))) }
func Xcos(x float64) float64              { return math.Cos(x) }
func Xcosf(x float32) float32             { return float32(math.Cos(float64(x))) }
func Xcosh(x float64) float64             { return math.Cosh(x) }
func Xexp(x float64) float64              { return math.Exp(x) }
func Xfabs(x float64) float64             { return math.Abs(x) }
func Xfabsf(x float32) float32            { return float32(math.Abs(float64(x))) }
func Xfloor(x float64) float64            { return math.Floor(x) }
func Xfmod(x, y float64) float64          { return math.Mod(x, y) }
func Xhypot(x, y float64) float64         { return math.Hypot(x, y) }
func Xisnan(x float64) int32              { return X__builtin_isnan(x) }
func Xisnanf(x float32) int32             { return Bool32(math.IsNaN(float64(x))) }
func Xisnanl(x float64) int32             { return Bool32(math.IsNaN(x)) } // ccgo has to handle long double as double as Go does not support long double.
func Xldexp(x float64, exp int32) float64 { return math.Ldexp(x, int(exp)) }
func Xlog(x float64) float64              { return math.Log(x) }
func Xlog10(x float64) float64            { return math.Log10(x) }
func Xround(x float64) float64            { return math.Round(x) }
func Xsin(x float64) float64              { return math.Sin(x) }
func Xsinf(x float32) float32             { return float32(math.Sin(float64(x))) }
func Xsinh(x float64) float64             { return math.Sinh(x) }
func Xsqrt(x float64) float64             { return math.Sqrt(x) }
func Xtan(x float64) float64              { return math.Tan(x) }
func Xtanh(x float64) float64             { return math.Tanh(x) }
func Xtrunc(x float64) float64            { return math.Trunc(x) }

func Xpow(x, y float64) float64 {
	r := math.Pow(x, y)
	if x > 0 && r == 1 && y >= -1.0000000000000000715e-18 && y < -1e-30 {
		r = 0.9999999999999999
	}
	return r
}
