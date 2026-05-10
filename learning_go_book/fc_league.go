package main

/*
1. You have been asked to manage a basketball league and are going to write a program to help you.
Define two types. The first one, called Team, has a field for the name of the team and a field for the player names.
The second type is called League and has a field called Teams for the teams in the league and a field called Wins that maps a team’s name to its number of wins.

2. Add two methods to League. The first method is called MatchResult. It takes four parameters: the name of the first team, its score in the game, the name of the second team, and its score in the game.
This method should update the Wins field in League.
Add a second method to League called Ranking that returns a slice of the team names in order of wins.
Build your data structures and call these methods from the main function in your program using some sample data.

3. Define an interface called Ranker that has a single method called Ranking that returns a slice of strings.
Write a function called RankPrinter with two parameters, the first of type Ranker and the second of type io.Writer.
Use the io.WriteString function to write the values returned by Ranker to the io.Writer, with a newline separating each result.
Call this function from main.
*/

import (
	"fmt"
	"io"
	"os"
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

func (l League) Ranking() []string {
	ranks := make([]string, len(l.Teams))
	for i, team := range l.Teams {
		ranks[i] = team.Name
	}

	slices.SortFunc(ranks, func(a, b string) int {
		if l.Wins[a] > l.Wins[b] {
			return -1
		} else {
			return 1
		}
	})

	return ranks
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

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for i, rank := range r.Ranking() {
		io.WriteString(w, fmt.Sprintf("%d. %s\n", i+1, rank))
	}
}

func prepareLeagueData() League {
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
	barcelona := Team{
		Name:        "Barcelona",
		PlayerNames: []string{"Lionel Messi", "Xavi Hernandez", "Andres Iniesta"},
	}

	league.Teams = []Team{barcelona, manchesterCity, arsenal, bavaria, manchesterUnited}
	league.Wins = make(map[string]int)

	league.MatchResult(arsenal.Name, 3, manchesterCity.Name, 2)
	league.MatchResult(arsenal.Name, 1, manchesterUnited.Name, 0)
	league.MatchResult(manchesterCity.Name, 5, manchesterUnited.Name, 4)
	league.MatchResult(bavaria.Name, 0, manchesterUnited.Name, 1)
	league.MatchResult(manchesterCity.Name, 2, bavaria.Name, 1)
	league.MatchResult(arsenal.Name, 3, bavaria.Name, 0)

	return league
}

func main() {
	l := prepareLeagueData()

	fmt.Println("Calling Ranking() method:")
	fmt.Println(l.Ranking())

	fmt.Println("----------")

	fmt.Println("Calling RankPrinter func:")
	RankPrinter(l, os.Stdout)
}
