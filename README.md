# pokedexcli

## Description

This project was an exercise for me to get a deeper understanding of developing a CLI that works with information 
requested from an API ([PokeAPI](https://pokeapi.co/)).

## Installation

To install the project on MacOS or Linux be sure to have Go installed.
The run the following commands:
```BASH
git clone https://github.com/langer-net/pokedexcli
cd pokedexcli
go build
```

## Usage

To run the project use:
```BASH
./pokedexcli
```

The CLI gives you the ability to view all the different maps of all the Pokemon games. You can use `map` to retrieve the 
next 20 cards from the API or cache and `mapb` to get the previous 20 cards. To see all Pokemon on a particular map, you 
can `explore` the map. Use `catch` with the name of the Pokemon to try to catch the given Pokemon. If you catch a 
Pokemon, it will be added to your Pokedex. You can see all entries with the command `pokedex`. When you have caught a 
Pokemon, you can also `inspect` it to see various statistics about it. To exit the application, simply use `exit`. With 
the command `help` it is not necessary to remember all these different commands, because it lists all commands with a 
description for you.

## Go Version

This project uses Go version 1.21.6. 

## Dependencies

There are no external dependencies for this project.

## Testing

To execute the tests, just run:
```BASH
go test
```

## License

This project is licensed under MIT licence.
