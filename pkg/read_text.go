package pkg

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/chunck"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/extract"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/load_page"
	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type ReadText struct {
	Url     string
	Limit   int64
	Sources *goquery.Document
	Data    DataReadText
}
type DataReadText struct {
	Text           string   `json:"text"`
	TotalCaracters int64    `json:"total_characters"`
	CountChunck    int64    `json:"-"`
	Chuncks        []string `json:"-"`
	Url            string   `json:"url"`
}

type ResponseReadtext struct {
	Failure   bool           `json:"failure"`
	Success   bool           `json:"success"`
	Message   string         `json:"message"`
	Data      []DataReadText `json:"data"`
	Timestamp float64        `json:"ts"`
}

func NewReadText(url string, limit int64, source *goquery.Document) ReadText {
	return ReadText{
		Url:     url,
		Limit:   limit,
		Sources: source,
	}
}

func (c *ReadText) Call() (ResponseReadtext, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := ResponseReadtext{
			Failure:   true,
			Success:   false,
			Message:   err.Error(),
			Timestamp: ts.GetTime(),
		}
		return responseErro, err
	}

	var page load_page.LoadPage
	var err error

	if c.Sources == nil {
		page = load_page.NewLoadPage(c.Url)
		err = page.Call()
	} else {
		page = load_page.LoadPage{
			Url:    c.Url,
			Source: c.Sources,
		}
		err = nil
	}

	if err != nil {
		ts.End()
		responseErro := ResponseReadtext{
			Failure:   true,
			Success:   false,
			Message:   err.Error(),
			Timestamp: ts.GetTime(),
		}
		return responseErro, err
	}

	informatin := extract.NewText(page.Source)
	extractext := informatin.Call()
	chuncks := chunck.NewChunck(extractext.Text, c.Limit)
	chuncks.Call()

	ts.End()

	data := DataReadText{
		Text:           extractext.Text,
		TotalCaracters: int64(len(extractext.Text)),
		CountChunck:    chuncks.CountChunck,
		Chuncks:        chuncks.ListChuncks,
		Url:            c.Url,
	}
	c.Data = data
	datas := []DataReadText{data}

	responseSuccess := ResponseReadtext{
		Failure:   false,
		Success:   true,
		Message:   "Success",
		Data:      datas,
		Timestamp: ts.GetTime(),
	}

	return responseSuccess, nil
}
