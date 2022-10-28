package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexflint/go-arg"
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
	s := spinner.New(spinner.CharSets[7], 50*time.Millisecond)
	s.Prefix = "Fetching Live Cricket score "
	s.Color("green", "bold")
	var args struct {
		Live  bool   `arg:"-l,--live" help:"Display Current Live Cricket Score"`
		Match string `arg:"-m,--match" help:"Display Current Live Cricket Score from the Cricbuzz Live Match URL"`
	}
	arg.MustParse(&args)

	if args.Live == true {
		resp, err := http.Get("https://cricket-api.vercel.app/live")
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		var post PostResponse

		if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
			log.Fatal(err)
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
			fmt.Printf(color.YellowString("Batsman Run: %v - ", post.Batsmanrun))
			fmt.Printf(color.YellowString("Batsman SR: %v - ", post.Sr))
			fmt.Printf(color.YellowString("Balls: %v\n", post.Ballsfaced))
			fmt.Printf(color.CyanString("Current Bowler: %v\n", post.Bowler))
			fmt.Printf(color.CyanString("Current Bowler Over: %v - ", post.Bowlerover))
			fmt.Printf(color.CyanString("Runs: %v - ", post.Bowlerruns))
			fmt.Printf(color.CyanString("Wicket: %v\n", post.Bowlerwickets))
		}
	} else if args.Match != "" {
		client := &http.Client{}
		resp, err := http.NewRequest(http.MethodGet, "https://cricket-api.vercel.app/score", nil)
		if err != nil {
			log.Fatal(err)
		}
		q := resp.URL.Query()
		q.Add("url", args.Match)
		resp.URL.RawQuery = q.Encode()

		req, err := client.Do(resp)
		if err != nil {
			fmt.Println("Errored when sending request to the server")
			return
		}

		defer req.Body.Close()

		var post PostResponse

		if err := json.NewDecoder(req.Body).Decode(&post); err != nil {
			log.Fatal(err)
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
			fmt.Printf(color.YellowString("Batsman SR: %v - ", post.Sr))
			fmt.Printf(color.YellowString("Balls: %v\n", post.Ballsfaced))
			fmt.Printf(color.CyanString("Current Bowler: %v\n", post.Bowler))
			fmt.Printf(color.CyanString("Current Bowler Over: %v - ", post.Bowlerover))
			fmt.Printf(color.CyanString("Runs: %v - ", post.Bowlerruns))
			fmt.Printf(color.CyanString("Wicket: %v\n", post.Bowlerwickets))
		}
	} else {
		fmt.Printf(color.MagentaString("Please Enter a Valid Option to Fetch the Score \n"))
	}
}
