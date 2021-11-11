package model

type EventType uint8

type EventStatus uint8

const (
	InvalidType EventType = iota
	Created
	Updated
	Removed

	InvalidStatus EventStatus = iota
	Idle
	Deferred
	Processed
)
