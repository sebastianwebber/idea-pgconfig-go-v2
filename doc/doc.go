package doc

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// ParamDoc is foo
type ParamDoc struct {
	Title          string
	ShortDesc      string
	Text           []string
	DocURL         string
	ConfURL        string
	Recomendations string
}

func formatVer(ver float32) string {
	if ver < 10 {
		return fmt.Sprintf("%.1f", ver)
	}

	return fmt.Sprintf("%.0f", ver)
}

// Get does foo
func Get(param string, ver float32) (ParamDoc, error) {

	var out ParamDoc
	out.ConfURL = fmt.Sprintf("https://postgresqlco.nf/en/doc/param/%s/%s/", param, formatVer(ver))

	res, err := http.Get(out.ConfURL)

	if err != nil {
		return out, fmt.Errorf("could not get URL: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return out, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// title
	sel := doc.Find("body > div.wrapper > div > section.content-header > div > div.col-md-8 > h1.parameter-title")
	for i := range sel.Nodes {
		out.Title = sel.Eq(i).Text()
	}

	// short desc
	sel = doc.Find("body > div.wrapper > div > section.content > div > div.col-md-8 > div.box.box-solid.box-primary > div:nth-child(1) > strong")
	for i := range sel.Nodes {
		out.ShortDesc = sel.Eq(i).Text()
	}

	// doc text
	sel = doc.Find("body > div.wrapper > div > section.content > div > div.col-md-8 > div.box.box-solid.box-primary > div.box-body > p")
	for i := range sel.Nodes {

		out.Text = append(out.Text, sel.Eq(i).Text())
	}

	// doc url?
	sel = doc.Find("body > div.wrapper > div > section.content > div > div.col-md-8 > div.box.box-solid.box-primary > div:nth-child(3) > span:nth-child(1) > a")
	for i := range sel.Nodes {
		single, e := sel.Eq(i).Attr("href")

		if e {
			out.DocURL = single
		}
	}

	// recomendations
	sel = doc.Find("body > div.wrapper > div > section.content > div > div.col-md-8 > div:nth-child(3) > div.box-body")
	for i := range sel.Nodes {
		out.Recomendations = sel.Eq(i).Text()
	}

	return out, nil
}
