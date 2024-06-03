package main

import "fmt"

type magainze struct {
	publication
}

func (m magainze) String() string {
	return fmt.Sprintf("This is a magazine named %s, it has %d pages and it is published by %s", m.name, m.pages, m.publisher)
}

func createMagazine(name string, pg int, pub string) iPublication {
	return &magainze{
		publication{
			name:      name,
			pages:     pg,
			publisher: pub,
		},
	}
}
