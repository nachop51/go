package main

import "fmt"

type newspaper struct {
	publication
}

func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s, it has %d pages and it is published by %s", n.name, n.pages, n.publisher)
}

func createNewspaper(name string, pg int, pub string) iPublication {
	return &newspaper{
		publication{
			name:      name,
			pages:     pg,
			publisher: pub,
		},
	}
}
