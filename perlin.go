/*
  To the extent possible under law, Ian Davis has waived all copyright
  and related or neighboring rights to this Source Code file.
  This work is published from the United Kingdom. 
*/
package perlin

// Based on algorithm given at http://freespace.virgin.net/hugo.elias/models/m_perlin.htm

import (
  "math"
)	


var primes = [5][3]int64 {
							[3]int64{15731, 789221, 1376312589},
							[3]int64{99877, 103867, 9576890767},
							[3]int64{27653, 222437, 2860486313},
							[3]int64{30181, 326083, 3367900313},
							[3]int64{52709, 450299, 5915587277},

						}


func noise(x, y int64, seed int64, octave int) float64 {
	n := x + y*57 + seed*7
	fn := (n << 13) ^ n
	return (1.0 - float64((fn*(fn*fn*15731+789221)+1376312589)&0x7fffffff)/float64(0x40000000))
}

func smoothedNoise(x float64, y float64, seed int64, octave int) float64 {
	xint := int64(math.Trunc(x))
	yint := int64(math.Trunc(y))

	corners := (noise(xint-1, yint-1, seed, octave) + noise(xint+1, yint-1, seed, octave) + noise(xint-1, yint+1, seed, octave) + noise(xint+1, yint+1, seed, octave)) / 16
	sides := (noise(xint-1, yint, seed, octave) + noise(xint+1, yint, seed, octave) + noise(xint, yint-1, seed, octave) + noise(xint, yint+1, seed, octave)) / 8
	center := noise(xint, yint, seed, octave) / 4
	return corners + sides + center
}

func interpolate(a, b, x float64) float64 {
	ft := x * math.Pi
	f := (1 - math.Cos(ft)) * 0.5
	return a*(1-f) + b*f
}

func interpolatedNoise(x, y float64, seed int64, octave int) float64 {

	xint := math.Trunc(x)
	xfrac := x - xint

	yint := math.Trunc(y)
	yfrac := y - yint

	v1 := smoothedNoise(xint, yint, seed, octave)
	v2 := smoothedNoise(xint+1, yint, seed, octave)
	v3 := smoothedNoise(xint, yint+1, seed, octave)
	v4 := smoothedNoise(xint+1, yint+1, seed, octave)

	i1 := interpolate(v1, v2, xfrac)
	i2 := interpolate(v3, v4, xfrac)

	return interpolate(i1, i2, yfrac)

}

func Noise2D(x, y float64, seed int64, persistence float64, octaves int) (value float64) {

	for i := 0; i < octaves; i++ {
		frequency := math.Pow(2, float64(i))
		amplitude := math.Pow(persistence, float64(i))
		value += interpolatedNoise(x*frequency, y*frequency, seed, i) * amplitude
	}
	return
}
