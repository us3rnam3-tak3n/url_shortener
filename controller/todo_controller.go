package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"url_shortening/service"
)

//Calls function to load HTML basic homepage.
func Home(c *gin.Context){
	service.Home(c)
}

type URLS[] string

func CreateTodo(c *gin.Context) {
	service.CreateTodo(c)
}

func FileHome(c *gin.Context){
	service.FileHom(c)
}

func FileParse(c *gin.Context){
	r := c.Request
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("file name => ", handler.Filename)

	tempFile, err := ioutil.TempFile("/tmp", "*.json")
	//fmt.Println(tempFile.Name())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
	fmt.Println("Done", tempFile.Name())

	parseJSON(tempFile.Name())

	c.String(http.StatusOK, "File Uploaded Successfully")

}

func parseJSON(fileName string) {
	path := fileName
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(jsonFile)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var fileURLs URLS
	err = json.Unmarshal([]byte(byteValue), &fileURLs)
	if err != nil {
		fmt.Println(err)
	}

	CreateTodoFile(fileURLs)
}

func CreateTodoFile(Urls [] string){
	var wg sync.WaitGroup
	for _, url := range Urls {
		wg.Add(1)
		go func(temp_url string) {
			service.CreateTodoUrl(temp_url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}