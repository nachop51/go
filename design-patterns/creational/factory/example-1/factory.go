package main

import "errors"

func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	switch pubType {
	case "newspaper":
		return createNewspaper(name, pg, pub), nil
	case "magazine":
		return createMagazine(name, pg, pub), nil
	default:
		return nil, errors.New("Invalid publication type")
	}
}
