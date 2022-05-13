package main

type Uuid uint32

type User struct {
	id   Uuid
	Name string
}

type QaCoin uint

type Reward struct {
	From   User
	To     User
	Amount QaCoin
	Note   string
}

type Users interface {
	all() []User
}

type Rewards interface {
	all() []Reward
	last(limit uint) []Reward
	store(reward Reward)
}
