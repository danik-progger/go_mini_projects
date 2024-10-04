package internal

import (
	"movies_crud/api"
)

var MoviesDB = []api.Movie{
	{
		Id:    "1",
		Isbn:  "438227",
		Title: "Moonrising kingdom",
		Director: &api.Director{
			FirstName: "Paul Thomas",
			LastName:  "Anderson",
		},
	},
	{
		Id:    "2",
		Isbn:  "45455",
		Title: "Gentlemen",
		Director: &api.Director{
			FirstName: "Guy",
			LastName:  "Ritchie",
		},
	},
}
