package service	

import(
	"github.com/testgolang/config"
	"github.com/olivere/elastic"
	"context"
	"fmt"
	"encoding/json"
	uuid "github.com/satori/go.uuid"

)


type Returnd struct {
	_index string
	_type  string
	_id string
	_score float64
	_source  struct{
		companyId string
		companyType string
		technologies string
		location string
	}
}


type SearchService struct {
}


type DocumentResponse struct {
	CompanyId    uuid.UUID  `json:"companyId"`
	Name         string   `json:"name"`
	// Technologies []string `gorm: "technologies type:string[]"`
	Technologies string `json:"technologies"`
	CompanyType  string   `json:"companyType"`
	Location     string   `json:"location"`
}

type SearchResponse struct {
	Time      string             `json:"time"`
	Hits      string             `json:"hits"`
	Documents []DocumentResponse `json:"documents"`
}
func (searchService *SearchService) Search(value string) {
	fmt.Println("hih")
	ctx := context.Background()
	client := config.GetElactic(ctx)
	termQuery := elastic.NewMultiMatchQuery(value,"companyType","location","technologies","name")
	result, err := client.Search().
		Index("job").
		Query(termQuery).
		From(0).Size(10).
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

}