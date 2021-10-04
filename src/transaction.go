package main

type Transaction struct {
	issuer         *Shard
	id             uint
	timestamp      float64
	time_validated float64
	is_proof       bool
	validator      *Transaction
	scenario       uint
	time_seen      float64
}
