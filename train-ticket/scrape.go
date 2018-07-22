package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Course stores information about a coursera course
type Ticket struct {
	Name       string
	Departure string
	Arrival     string
	DepartTime       string
	Price         string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: coursera.org, www.coursera.org
		//colly.AllowedDomains("coursera.org", "www.coursera.org"),

		// Cache responses to prevent multiple download of pages
		// even if the collector is restarted
		//colly.CacheDir("./coursera_cache"),
	)

	tickets := make([]Ticket, 0, 200)





	// Extract details of the course
	c.OnHTML(`tbody[id=tbody_depart]`, func(e *colly.HTMLElement) {
		//log.Println("Course found", e.Request.URL)

		e.ForEach("tr.item-list", func(el *colly.HTMLElement) {
			ticketAvailable := e.ChildText(".orange-button")
			if ticketAvailable != "" {
				ticket := Ticket{
					Name:       e.Attr("data-name"),
					Departure:         e.Attr("data-departure-station-name"),
					Arrival: e.Attr("data-arrival-station-name"),
					DepartTime:     e.Attr("data-depart"),
					Price: e.Attr("data-price"),
				}

				tickets = append(tickets, ticket)
			}
		})
	})

	// Start scraping on http://coursera.com/browse
	c.Visit("https://www.tiket.com/kereta-api/cari?d=SLO&dt=STATION&a=BANDUNG&at=CITY&date=2018-06-17&ret_date=&adult=1&infant=0")

	if len(tickets) > 0 {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")

		// Dump json to the standard output
		enc.Encode(tickets)
	}

}
