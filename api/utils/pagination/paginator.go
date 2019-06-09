package pagination

import (
	"log"
	"math"
	"net/url"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

const DefaultPage = 1
const DefaultLimit = 5
const DefaultOrder = "id DESC"

const ASC = "ASC"
const DESC = "DESC"

type (
	options struct {
		Limit      int     `json:"items_per_page"`
		Page       int     `json:"page"`
		TotalPages float64 `json:"total_pages"`
		Count      int     `json:"count"`
		Order      string  `json:"order"`
	}

	Paginator struct {
		Data    interface{} `json:"data"`
		Options options     `json:"options"`
	}
)

func New(urlParameter url.Values) *Paginator {

	page := getNumberQueryParam(urlParameter.Get("page"), DefaultPage)
	limit := getNumberQueryParam(urlParameter.Get("limit"), DefaultLimit)
	order := getMultipleQueryParam(urlParameter.Get("order"), DefaultOrder)

	if page < DefaultPage {
		page = DefaultPage
	}

	options := options{
		Limit: limit,
		Page:  page,
		Order: order,
	}

	return &Paginator{Options: options}
}

// Set Pagination data based on the given page nr
func (paginator *Paginator) Build(collection *mgo.Collection, resource interface{}) {

	count, countError := collection.Count()
	if countError != nil {
		log.Fatal(countError)
	}

	order := paginator.Options.Order
	page := paginator.Options.Page
	limit := paginator.Options.Limit

	paginator.Options.Count = count
	paginator.Options.TotalPages = math.Round(float64(count) / float64(limit))

	offset := limit * (page - 1)

	queryError := collection.Find(nil).
		Sort(order, "true").
		Skip(offset).
		Limit(limit).
		All(resource)

	if queryError != nil {
		log.Fatal(queryError)
	}

	paginator.Data = resource
}

func getNumberQueryParam(urlParameter string, d int) int {

	if l, err := strconv.Atoi(urlParameter); err == nil {
		return l
	}

	return d
}

func getMultipleQueryParam(urlParameter string, d string) string {

	v := strings.Split(urlParameter, ",")

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
