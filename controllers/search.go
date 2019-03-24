package controllers
import(
	"fmt"
	"github.com/gin-gonic/gin"
	service "github.com/testgolang/service"

)
type SearchControllers struct {

}

func (searchControllers *SearchControllers) Search (c *gin.Context){
	fmt.Println("---------------------")
	text,err := c.GetQuery("text")
	fmt.Println(text)

	if err== false  {
		panic("err")
	}
	searchService := service.SearchService{}
	searchService.Search(text)
}