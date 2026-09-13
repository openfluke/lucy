package lucy_test

import (
	"testing"
	"time"

	"github.com/openfluke/lucy/lucy"
)

func TestLucyScoreEquation(t *testing.T) {
	s := lucy.Snapshot{
		TotalOutputs:   1000,
		TotalCorrect:   800,
		InferMs:        800,
		TrainMs:        200,
		Duration:       time.Second,
		SoftAcc:        75,
		AccuracyPulses: 2,
		Windows: []lucy.Window{
			{SoftAcc: 80, Accuracy: 80},
			{SoftAcc: 70, Accuracy: 70},
		},
	}
	lucy.Finalize(&s)
	// thru=1000, avail=80, Acc=80 (TotalCorrect/TotalOutputs) → 1000*80*80/10000 = 640
	if s.Throughput != 1000 {
		t.Fatalf("throughput %v", s.Throughput)
	}
	if s.Availability != 80 {
		t.Fatalf("availability %v", s.Availability)
	}
	if s.SoftAcc != 75 {
		t.Fatalf("softAcc %v", s.SoftAcc)
	}
	if s.AvgAccuracy != 80 {
		t.Fatalf("acc %v want 80", s.AvgAccuracy)
	}
	if s.Score != 640 {
		t.Fatalf("score %v want 640 (T x Avail x Acc / 10000)", s.Score)
	}
}

func TestSoftAccOne(t *testing.T) {
	if lucy.SoftAccOne(0.5, 0.5) != 100 {
		t.Fatalf("exact match")
	}
	a := lucy.SoftAccOne(0.55, 0.5)
	if a < 49.9 || a > 50.1 {
		t.Fatalf("softAcc %v want ~50", a)
	}
}

func TestSoftAccProb(t *testing.T) {
	a := lucy.SoftAccProb(0.8, 1.0)
	if a < 79.9 || a > 80.1 {
		t.Fatalf("softAccProb %v want ~80", a)
	}
}
