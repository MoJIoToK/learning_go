package main

import "sync"

type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry // slice of entries sorted from longest to shortest.
	es    []muxEntry          // whether any patterns contain hostnames
	hosts bool
}
