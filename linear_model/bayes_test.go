package linear_model

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestBayesianRidge(t *testing.T) {

	var X [][]float = make([][]float, 10000)
	Y := make([]float, len(X))
	f := func(X []float) float { return 1. + 2.*X[0] + 3.*X[1] + 4.*X[2] }
	for i := range X {
		X[i] = make([]float, 3, 3)
		for j := range X[i] {
			X[i][j] = rand.Float64()*20. - 10.
		}
		Y[i] = f(X[i]) + (rand.Float64()-.5)/2
	}
	for _, normalize := range []bool{false, true} {
		m := NewBayesianRidge()
		m.Normalize = normalize
		//m.Verbose = true
		m.ComputeScore = true
		m.Fit(X, Y)
		fmt.Printf("TestBayesianRidge normalize=%v score:%.4g\n", normalize, m.Score(X, Y, nil))
		eps := .1
		Xp := [][]float{{7., 8., 9.}}
		y_true := []float{f(Xp[0])}
		Yp := m.Predict(Xp)
		//fmt.Println(Yp[0], " expected: ", y_true)
		if math.Abs(Yp[0]-y_true[0]) > eps {
			fmt.Printf("TestBayesianRidge Yp[0]-y_true[0]=%g\n", Yp[0]-y_true[0])
			t.Fail()
		}
		// Output:
		// 75.
	}
}

func unused(...interface{}) {}

func ExampleBayesianRidge() {
	var X [][]float = make([][]float, 10000)
	Y := make([]float, len(X))
	f := func(X []float) float { return 1. + 2.*X[0] + 3.*X[1] + 4.*X[2] }
	for i := range X {
		X[i] = make([]float, 3, 3)
		for j := range X[i] {
			X[i][j] = rand.Float64()
		}
		Y[i] = f(X[i]) + (rand.Float64()-.5)/2
	}

	m := NewBayesianRidge()
	m.Normalize = true
	//m.Verbose = true
	m.ComputeScore = true
	m.Fit(X, Y)
	//fmt.Printf("Scores: %#v\n", m.Scores_)
	Xp := [][]float{{7., 8., 9.}}
	y_true := []float{f(Xp[0])}
	Yp := m.Predict(Xp)
	unused(y_true, Yp)

	//fmt.Println(Yp[0], " expected: ", y_true)
	fmt.Printf("TestBayesianRidge score:%.2g\n", m.Score(X, Y, nil))
	// Output:
	// TestBayesianRidge score:0.99
}