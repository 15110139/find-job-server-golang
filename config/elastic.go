package config
import (
	"fmt"
	"context"

	"github.com/olivere/elastic"
)
func GetElactic( ctx context.Context  ) *elastic.Client   {
	client,err := elastic.NewClient()
	if err != nil{
		panic(err)
	}
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	return client

}