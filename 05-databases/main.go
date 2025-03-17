package main

import "fmt"

var movies = []Row{
	{
		Entries: []Entry{
			{"id", "1"},
			{"title", "Toy Story (1995)"},
			{"genre", "Adventure|Animation|Children|Comedy|Fantasy"},
		},
	},
	{
		Entries: []Entry{
			{"id", "2"},
			{"title", "Jumanji (1995)"},
			{"genre", "Adventure|Children|Fantasy"},
		},
	},
	{
		Entries: []Entry{
			{"id", "3"},
			{"title", "Grumpier Old Men (1995)"},
			{"genre", "Comedy|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "4"},
			{"title", "Waiting to Exhale (1995)"},
			{"genre", "Comedy|Drama|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "5"},
			{"title", "Father of the Bride Part II (1995)"},
			{"genre", "Comedy"},
		},
	},
	{
		Entries: []Entry{
			{"id", "6"},
			{"title", "Heat (1995)"},
			{"genre", "Action|Crime|Thriller"},
		},
	},
	{
		Entries: []Entry{
			{"id", "7"},
			{"title", "Sabrina (1995)"},
			{"genre", "Comedy|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "8"},
			{"title", "Tom and Huck (1995)"},
			{"genre", "Adventure|Children"},
		},
	},
	{
		Entries: []Entry{
			{"id", "9"},
			{"title", "Sudden Death (1995)"},
			{"genre", "Action"},
		},
	},
	{
		Entries: []Entry{
			{"id", "10"},
			{"title", "GoldenEye (1995)"},
			{"genre", "Action|Adventure|Thriller"},
		},
	},
	{
		Entries: []Entry{
			{"id", "11"},
			{"title", "American President, The (1995)"},
			{"genre", "Comedy|Drama|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "12"},
			{"title", "Dracula: Dead and Loving It (1995)"},
			{"genre", "Comedy|Horror"},
		},
	},
	{
		Entries: []Entry{
			{"id", "11"},
			{"title", "American President, The (1995)"},
			{"genre", "Comedy|Drama|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "11"},
			{"title", "American President, The (1995)"},
			{"genre", "Comedy|Drama|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "13"},
			{"title", "Balto (1995)"},
			{"genre", "Adventure|Animation|Children"},
		},
	},
	{
		Entries: []Entry{
			{"id", "14"},
			{"title", "Nixon (1995)"},
			{"genre", "Drama"},
		},
	},
	{
		Entries: []Entry{
			{"id", "15"},
			{"title", "Cutthroat Island (1995)"},
			{"genre", "Action|Adventure|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "16"},
			{"title", "Casino (1995)"},
			{"genre", "Crime|Drama"},
		},
	},
	{
		Entries: []Entry{
			{"id", "17"},
			{"title", "Sense and Sensibility (1995)"},
			{"genre", "Drama|Romance"},
		},
	},
	{
		Entries: []Entry{
			{"id", "18"},
			{"title", "Four Rooms (1995)"},
			{"genre", "Comedy"},
		},
	},
	{
		Entries: []Entry{
			{"id", "19"},
			{"title", "Ace Ventura: When Nature Calls (1995)"},
			{"genre", "Comedy"},
		},
	},
	{
		Entries: []Entry{
			{"id", "20"},
			{"title", "Money Train (1995)"},
			{"genre", "Action|Comedy|Crime|Drama|Thriller"},
		},
	},
}

func main() {
	// scanner := NewScannerOperator(movies)
	genreDrama := ContainsGenreExpression{Column: "genre", Substr: "Drama"}

	selectionOperator := NewSelectionOperator(movies, genreDrama)

	for selectionOperator.Next() {
		fmt.Println(selectionOperator.Execute())
	}
}
