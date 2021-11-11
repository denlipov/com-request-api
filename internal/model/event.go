package model

// EventType ...
type EventType uint8

// EventStatus ...
type EventStatus uint8

const (
	// Created ...
	Created EventType = iota

	// Updated ...
	Updated

	// Removed ...
	Removed

	// Idle ...
	Idle EventStatus = iota

	// Deferred ...
	Deferred

	// Processed ...
	Processed
)
