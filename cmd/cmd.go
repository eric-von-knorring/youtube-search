package cmd

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var Format string

type VideoLink struct {
	Title string
	BaseUrl string
	Href  string
}

var rootCmd = &cobra.Command{
	Use:   "youtube-search <search query>",
	Short: "Search youtube for given text",
	Long:  `Search youtube for given text and outputs the title and link on standard out`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		argWords := make([]string, 0)
		for _, arg := range args {
			argWords = append(argWords, strings.Fields(arg)...)
		}
		mergedArgs := strings.Join(argWords, "+")
		searchQuery := strings.ReplaceAll(mergedArgs, " ", "+")

		response, err := http.Get("https://www.youtube.com/results?search_query=" + searchQuery)
		if err != nil {
			os.Exit(127)
		}
		defer response.Body.Close()

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("Error open search result response body. ", err)
		}

		template, err := template.New("output").Parse(Format)
		if err != nil {
			log.Fatal("Could not parse format flag from as go template")
		}

		document.Find("a").Each(func(_ int, element *goquery.Selection) {
			title, titleExist := element.Attr("title")
			href, hrefExist := element.Attr("href")

			if titleExist && hrefExist && strings.HasPrefix(href, "/watch?") {
				 _ = template.Execute(os.Stdout, VideoLink{
					title,
					"https://www.youtube.com",
					href,
				})
				fmt.Println()
			}
		})
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(
		&Format,
		"format",
		"f",
		"{{ .Title }} - {{ .BaseUrl }}{{ .Href }}",
		"Go template to format output.")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
