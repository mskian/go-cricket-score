package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

type PostResponse struct {
	Title                string `json:"title"`
	Update               string `json:"update"`
	Current              string `json:"current"`
	Batsman              string `json:"batsman"`
	Batsmanrun           string `json:"batsmanrun"`
	Ballsfaced           string `json:"ballsfaced"`
	Fours                string `json:"fours"`
	Sixes                string `json:"sixes"`
	Sr                   string `json:"sr"`
	Batsmantwo           string `json:"batsmantwo"`
	Batsmantworun        string `json:"batsmantworun"`
	Batsmantwoballsfaced string `json:"batsmantwoballsfaced"`
	Batsmantwofours      string `json:"batsmantwofours"`
	Batsmantwosixes      string `json:"batsmantwosixes"`
	Batsmantwosr         string `json:"batsmantwosr"`
	Bowler               string `json:"bowler"`
	Bowlerover           string `json:"bowlerover"`
	Bowlerruns           string `json:"bowlerruns"`
	Bowlerwickets        string `json:"bowlerwickets"`
	Bowlermaiden         string `json:"bowlermaiden"`
	Bowlertwo            string `json:"bowlertwo"`
	Bowletworover        string `json:"bowletworover"`
	Bowlertworuns        string `json:"bowlertworuns"`
	Bowlertwowickets     string `json:"bowlertwowickets"`
	Bowlertwomaiden      string `json:"bowlertwomaiden"`
	Partnership          string `json:"partnership"`
	Recentballs          string `json:"recentballs"`
	Lastwicket           string `json:"lastwicket"`
	Runrate              string `json:"runrate"`
	Commentary           string `json:"commentary"`
}

func main() {
	resp, err := http.Get("https://cricket-api.vercel.app/live")
	s := spinner.New(spinner.CharSets[7], 50*time.Millisecond)
	s.Prefix = "Fetching Live Cricket score "
	s.Color("green", "bold")
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	var post PostResponse

	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		print(err)
	}

	s.Start()
	time.Sleep(1 * time.Second)
	s.Stop()
	if post.Current == "Data Not Found" {
		fmt.Printf(color.MagentaString("sorry Currently No Live Match \n"))
	} else {
		fmt.Printf(color.MagentaString("Match: %v\n", post.Title))
		fmt.Printf(color.GreenString("Score: %v\n", post.Current))
		fmt.Printf(color.BlueString("%v\n", post.Runrate))
		fmt.Printf(color.YellowString("Current Batsman: %v\n", post.Batsman))
		fmt.Printf(color.YellowString("Current Batsman Run: %v - ", post.Batsmanrun))
		fmt.Printf(color.YellowString("Batsman Run: %v - ", post.Batsmanrun))
		fmt.Printf(color.YellowString("Balls: %v\n", post.Ballsfaced))
		fmt.Printf(color.CyanString("Current Bowler: %v\n", post.Bowler))
		fmt.Printf(color.CyanString("Current Bowler Over: %v - ", post.Bowlerover))
		fmt.Printf(color.CyanString("Bowler Runs: %v - ", post.Bowlerruns))
		fmt.Printf(color.CyanString("Wicket: %v\n", post.Bowlerwickets))
	}
}
