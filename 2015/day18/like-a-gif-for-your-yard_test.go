// https://adventofcode.com/2015/day/18

package day18

import "testing"

func Test_ShouldCalculateLightsOnAfterXSteps(t *testing.T) {
	result := calculateLightsOnAfterXSteps(100)
	if result != 821 {
		t.Fatalf(`calculateLightsOnAfterXSteps(100) = %v, want %v`, result, 821)
	}

}

func Test_ShouldCalculateLightsOnAfterXStepsWithCornersAlwaysOn(t *testing.T) {
	result := calculateLightsOnAfterXStepsWithCornersAlwaysOn(100)
	if result != 886 {
		t.Fatalf(`calculateLightsOnAfterXStepsWithCornersAlwaysOn(100) = %v, want %v`, result, 886)
	}
}
