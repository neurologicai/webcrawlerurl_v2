package pkg

import (
	"errors"

	"github.com/gildemberg-santos/webcrawlerurl_v2/util/timestamp"
)

type ReadText struct {
	Url string
}

type responseSuccessReadText struct {
	ReadText struct {
		Text string `json:"text"`
	} `json:"readtext"`
	Url        string  `json:"url"`
	Timestamp  float64 `json:"ts"`
	StatusCode int     `json:"status_code"`
}

type responseErroReadtext struct {
	Erro       string      `json:"erro"`
	ReadText   interface{} `json:"readtext"`
	Url        string      `json:"url"`
	Timestamp  float64     `json:"ts"`
	StatusCode int         `json:"status_code"`
}

func (c *ReadText) Call() (interface{}, error) {
	ts := timestamp.NewTimestamp().Start()

	if c.Url == "" {
		err := errors.New("url is empty")
		ts.End()
		responseErro := responseErroReadtext{
			Erro:       err.Error(),
			ReadText:   nil,
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: 500,
		}
		return responseErro, err
	}

	pagina := LoadPage{
		Url: c.Url,
	}

	err := pagina.Load()
	if err != nil {
		ts.End()
		responseErro := responseErroReadtext{
			Erro:       err.Error(),
			ReadText:   nil,
			Url:        c.Url,
			Timestamp:  ts.GetTime(),
			StatusCode: pagina.StatusCode,
		}
		return responseErro, err
	}

	informatin := ExtractText{}
	informatin.Init(pagina.Source)
	extractext, _ := informatin.Call()

	ts.End()
	responseSuccess := responseSuccessReadText{
		Url:        c.Url,
		Timestamp:  ts.GetTime(),
		StatusCode: pagina.StatusCode,
	}
	responseSuccess.ReadText.Text = extractext.(ExtractText).Text

	return responseSuccess, nil
}
