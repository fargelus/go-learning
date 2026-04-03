package main

import (
	"fmt"
	"maps"
	"slices"
)

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstTeamName string, firstTeamScore int, secondTeamName string, secondTeamScore int) {
	if firstTeamScore == secondTeamScore {
		return
	}

	if firstTeamScore > secondTeamScore {
		l.Wins[firstTeamName] += 1
	} else {
		l.Wins[secondTeamName] += 1
	}
}

func (l *League) Ranking() []string {
	return slices.Collect(maps.Keys(l.Wins))
}

func main() {
	league := League{}
	arsenal := Team{
		Name:        "Arsenal",
		PlayerNames: []string{"Robin Van Persie", "Alexis Sanchez", "Thierry Henry"},
	}
	manchesterCity := Team{
		Name:        "Manchester City",
		PlayerNames: []string{"Sergio Agüero", "David Silva", "Vincent Kompany"},
	}
	manchesterUnited := Team{
		Name:        "Manchester United",
		PlayerNames: []string{"Cristiano Ronaldo", "Wayne Rooney", "Gary Neville"},
	}
	bavaria := Team{
		Name:        "Bayern Munich",
		PlayerNames: []string{"Arjen Robben", "Frank Ribery", "Manuel Neuer"},
	}

	league.Teams = []Team{arsenal, manchesterCity, manchesterUnited}
	league.Wins = make(map[string]int)

	league.MatchResult(arsenal.Name, 3, manchesterCity.Name, 2)
	league.MatchResult(arsenal.Name, 1, manchesterUnited.Name, 0)
	league.MatchResult(manchesterCity.Name, 5, manchesterUnited.Name, 4)
	league.MatchResult(bavaria.Name, 0, manchesterUnited.Name, 1)
	league.MatchResult(manchesterCity.Name, 2, bavaria.Name, 1)
	league.MatchResult(arsenal.Name, 3, bavaria.Name, 0)

	fmt.Println(league.Ranking())
}
