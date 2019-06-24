package pagination

import (
	"net/url"
	"testing"
)

func TestNewOptionsPage(t *testing.T) {

	query := url.Values{}
	query["page"] = []string{"5"}

	paginator := New(query)

	result := struct {
		got  int
		want int
	}{paginator.Options.Page, 5}

	if result.got != result.want {
		t.Errorf("Limit was incorrect, got: %v, want: %v.", result.got, result.want)
	}
}

func TestNewOptionsLimit(t *testing.T) {

	query := url.Values{}
	query["limit"] = []string{"5"}

	paginator := New(query)

	result := struct {
		got  int
		want int
	}{paginator.Options.Limit, 5}

	if result.got != result.want {
		t.Errorf("Limit was incorrect, got: %v, want: %v.", result.got, result.want)
	}
}

func TestNewOptionsOrder(t *testing.T) {

	query := url.Values{}
	query["order"] = []string{"id,ASC"}

	paginator := New(query)

	result := struct {
		got  string
		want string
	}{paginator.Options.Order, "id ASC"}

	if result.got != result.want {
		t.Errorf("Limit was incorrect, got: %v, want: %v.", result.got, result.want)
	}
}

// func TestBuild(t *testing.T) {
// 	t.Errorf("Not implemented")
// }
