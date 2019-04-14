package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/find-job-server-golang/config"
	"github.com/olivere/elastic"
	uuid "github.com/satori/go.uuid"
)

type Returnd struct {
	_index  string
	_type   string
	_id     string
	_score  float64
	_source struct {
		companyId    string
		companyType  string
		technologies []string
		location     string
	}
}

type SearchService struct {
}

type DocumentResponse struct {
	CompanyId uuid.UUID `json:"companyId"`
	Name      string    `json:"name"`
	// Technologies []string `gorm: "technologies type:string[]"`
	Technologies string `json:"technologies"`
	CompanyType  string `json:"companyType"`
	Location     string `json:"location"`
}

func (searchService *SearchService) Search(value string, page, limit int) []DocumentResponse {
	ctx := context.Background()
	client := config.GetElactic(ctx)
	termQuery := elastic.NewMultiMatchQuery(value, "companyType", "location", "technologies", "name")
	result, err := client.Search().
		Index("data-work").
		Query(termQuery).
		From(page*limit).Size(limit).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	// Transform search results before returning them
	docs := make([]DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc DocumentResponse
		json.Unmarshal(*hit.Source, &doc)
		fmt.Println(doc)
		docs = append(docs, doc)
	}
	return docs
}
