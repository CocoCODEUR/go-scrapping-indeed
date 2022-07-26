package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Job struct {
	Salary []int
	Titles []string
	// Society []string
}

func CleanSalary(salary string) int {
	regex, err := regexp.Compile(`[^\dàÀ]+`)
	if err != nil {
		log.Fatal(err)
	}
	processedString := strings.ToLower(salary)
	processedString = strings.Replace(processedString, "k", "000", -1)
	processedString = regex.ReplaceAllString(processedString, "")

	switch {
	case strings.Contains(processedString, "à"):
		{
			salaries := strings.Split(processedString, "à")
			s1, err := strconv.Atoi(salaries[0])
			if err != nil {
				log.Printf("string one to int conversion failed : %v", err)
			}
			s2, err := strconv.Atoi(salaries[1])
			if err != nil {
				log.Printf("string two to int conversion failed : %v", err)

			}
			averageSalary := (s1 + s2) / 2

			return averageSalary
		}
	default:
		{
			averageSalary, err := strconv.Atoi(processedString)
			if err != nil {
				log.Printf(" string to int conversion failed : %v", err)
			}
			return averageSalary
		}

	}

}

func main() {
	var Jobs Job
	t := time.Now()

	c := colly.NewCollector(colly.AllowedDomains("fr.indeed.com", "www.indeed.com"))

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	// titre de l'annonce

	c.OnHTML("h2.jobTitle", func(e *colly.HTMLElement) {

		Jobs.Titles = append(Jobs.Titles, e.ChildText("span"))

	})
	// salaire de l'annonce
	c.OnHTML("div.salary-snippet-container", func(e *colly.HTMLElement) {
		salaryConverted := CleanSalary(e.ChildText("div"))

		Jobs.Salary = append(Jobs.Salary, salaryConverted)
	})

	// page suivante
	c.OnHTML("a[aria-label=Suivant]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(request *colly.Request) {
		// fmt.Println("Visiting site : ", request.URL.String())

	})

	research := "next%20js"
	c.Visit("https://fr.indeed.com/emplois?q=" + research + "&start=0&vjk=623df0a492ab31b7")
	defer log.Println(time.Since(t))

	data, err := json.MarshalIndent(Jobs, "", " ")
	_ = ioutil.WriteFile("test.json", data, 0644)

	if err != nil {
		log.Printf("Cannot convert data to JSON: %v", err)
	}
	fmt.Println(data)

}
