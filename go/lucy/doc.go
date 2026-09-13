// Package lucy is the source-of-truth measuring math for live-fit boards.
//
// Score, Availability, SoftAcc, Q, LPD (Lucy Pareto density), gold / lean / trap
// bands — portable so Tide / Ocean / River hosts do not each reinvent the ruler.
//
// Defaults match the synthetic-organism benchmark (serve + train in a small box),
// but thresholds and duty-cycle definitions are meant to be tunable: wrap your own
// experiment, swap Acc keep floors, redefine Availability, feed Samples from any
// backend.
//
// Status: placeholder. Formulas currently live in welvet/lucy; this module will
// own them (and later wasm / native binary builds for npm + PyPI wrappers).
//
// Planned later: optional chart/JPG generation from ranked boards.
package lucy
