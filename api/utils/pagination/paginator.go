package pagination

import (
	"log"
	"math"
	"net/url"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

const Page = 1
const Limit = 5
const Order = "id DESC"

const ASC = "ASC"
const DESC = "DESC"

var Paginator Paginate

type Paginate struct {
	Data       interface{} `json:"data"`
	Limit      int         `json:"items_per_page"`
	Page       int         `json:"page"`
	TotalPages float64     `json:"total_pages"`
	Count      int         `json:"count"`
}

// Set Pagination data based on the given page nr
func Build(cllctn *mgo.Collection, q url.Values, m interface{}) {

	p := getNumberQueryParam(q.Get("page"), Page)
	l := getNumberQueryParam(q.Get("limit"), Limit)
	o := getMultipleQueryParam(q.Get("order"), Order)

	if p < Page {
		p = Page
	}

	offset := l * (p - 1)

	Paginator.Limit = l
	Paginator.Page = p

	// Count the total amount of records of this Model
	count, err := cllctn.Count()
	if err != nil {
		log.Fatal(err)
	}

	Paginator.Count = count

	Paginator.TotalPages = math.Round(float64(Paginator.Count) / float64(l))

	if err := cllctn.Find(nil).Sort(o, "true").Skip(offset).Limit(l).All(m); err != nil {
		log.Fatal(err)
	}

	Paginator.Data = m
}

func getNumberQueryParam(q string, d int) int {

	if l, err := strconv.Atoi(q); err == nil {
		return l
	}

	return d
}

func getMultipleQueryParam(q string, d string) string {

	v := strings.Split(q, ",")

	// check has 2 values
	if len(v) != 2 {
		return d
	}

	direction := strings.ToUpper(v[1])

	// check 2nd value is ASC or DESC
	if direction == ASC || direction == DESC {
		return v[0] + " " + v[1]
	}

	return d
}
