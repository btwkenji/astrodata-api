package core

import "github.com/kenjitheman/magneconn_api/vars"

func CalculateDst(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	minVal := data[0]
	for _, value := range data {
		if value < minVal {
			minVal = value
		}
	}
	Dst := minVal * vars.CONSTANT
	return Dst
}

func SplitData(data []float64) [][]float64 {
	if len(data) < 144 {
		return nil
	}

	var parts [][]float64

	part1 := data[:72]
	parts = append(parts, part1)

	part2 := data[72:144]
	parts = append(parts, part2)

	part3 := data[144:216]
	parts = append(parts, part3)

	part4 := data[216:]
	parts = append(parts, part4)

	for i := len(parts); i < 4; i++ {
		parts = append(parts, make([]float64, 0))
	}

	return parts
}

func CalculateDstForParts(parts [][]float64) []float64 {
	var results []float64

	for _, part := range parts {
		minVal := part[0]
		for _, value := range part {
			if value < minVal {
				minVal = value
			}
		}
		Dst := CalculateDst([]float64{minVal})
		results = append(results, Dst)
	}

	return results
}
