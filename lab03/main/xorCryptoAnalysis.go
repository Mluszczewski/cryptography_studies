package main

import (
	"strings"
)

var freq = [26]float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015,
	0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749,
	0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758,
	0.00978, 0.02360, 0.00150, 0.01974, 0.00074,
}

func sum(a []float64) (sum float64) {
	for _, f := range a {
		sum += f
	}
	return
}

func bestMatch(a []float64) int {
	sum := sum(a)
	bestFit, bestRotate := 1e100, 0
	for rotate := 0; rotate < 26; rotate++ {
		fit := 0.0
		for i := 0; i < 26; i++ {
			d := a[(i+rotate)%26]/sum - freq[i]
			fit += d * d / freq[i]
		}
		if fit < bestFit {
			bestFit, bestRotate = fit, rotate
		}
	}
	return bestRotate
}

func freqEveryNth(msg []int, key []byte) float64 {
	l := len(msg)
	interval := len(key)
	out := make([]float64, 26)
	accu := make([]float64, 26)
	for j := 0; j < interval; j++ {
		for k := 0; k < 26; k++ {
			out[k] = 0.0
		}
		for i := j; i < l; i += interval {
			out[msg[i]]++
		}
		rot := bestMatch(out)
		key[j] = byte(rot + 65)
		for i := 0; i < 26; i++ {
			accu[i] += out[(i+rot)%26]
		}
	}
	sum := sum(accu)
	ret := 0.0
	for i := 0; i < 26; i++ {
		d := accu[i]/sum - freq[i]
		ret += d * d / freq[i]
	}
	return ret
}

func decrypt(text, key string) string {
	var sb strings.Builder
	ki := 0
	for _, c := range text {
		if c < 'A' || c > 'Z' {
			continue
		}
		ci := (c ^ rune(key[ki]) + 26) % 26
		sb.WriteRune(ci + 65)
		ki = (ki + 1) % len(key)
	}
	return sb.String()
}
