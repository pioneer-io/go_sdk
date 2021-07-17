package models

type Flag struct {
	Id          int
	Title       string
	Description string
	Is_Active   bool
	Rollout     int
}