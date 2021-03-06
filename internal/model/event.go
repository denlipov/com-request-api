package model

// EventType ...
type EventType uint8

// EventStatus ...
type EventStatus uint8

const (
	// InvalidType ...
	InvalidType EventType = iota

	// Created ...
	Created

	// Updated ...
	Updated

	// Removed ...
	Removed

	// InvalidStatus ...
	InvalidStatus EventStatus = iota

	// Idle ...
	Idle

	// Deferred ...
	Deferred

	// Processed ...
	Processed
)
