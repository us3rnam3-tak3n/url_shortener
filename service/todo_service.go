package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"url_shortening/Model"
	"url_shortening/database_conn"
)

type URLS[] string

//A simple hashing function to create short links
func longtoshort(n uint)  string{
	enc := [62]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z","0","1","2","3","4","5","6","7","8","9"}
	var ans string
	for n!=0{
		ans+=enc[n%62]
		n/=62
	}
	return ans
}

//Loads basic HTML home
func Home(c *gin.Context){
	c.HTML(http.StatusOK,"index.html",gin.H{})
}

func FileHom(c *gin.Context){
	c.HTML(http.StatusOK,"file_index.html",gin.H{})
}
//Gets information from Post form, and searches if it is already present in the table.
//If not, then creates the short link for it, and adds the entry into the table.
//If already present, returns the previously stored short link.
func CreateTodo(c *gin.Context) {
	var todof Model.Todo
	todof.Long=c.PostForm("long")
	var cou int=0
	database_conn.Db.Model(&Model.Todo{}).Where("`long` LIKE ?", todof.Long).Count(&cou)
	if cou >= 1{
		database_conn.Db.Where("`long` = ?", c.PostForm("long")).First(&todof)
		c.HTML(http.StatusOK, "short.html", gin.H{
			"long" : todof.Long,
			"sho": todof.Short,
		})
		return
	}
	database_conn.Db.Save(&todof)
	todof.Short = longtoshort(todof.ID)
	database_conn.Db.Save(&todof)
	c.HTML(http.StatusOK,"short.html",gin.H{
		"sho" : todof.Short,
		"long": todof.Long,
	})
}
func CreateTodoUrl(d string){
	var todof Model.Todo
	todof.Long=d
	var cou int=0
	database_conn.Db.Model(&Model.Todo{}).Where("`long` LIKE ?", d).Count(&cou)
	if cou >= 1{
		database_conn.Db.Where("`long` = ?", d).First(&todof)
		return
	}
	database_conn.Db.Save(&todof)
	todof.Short = longtoshort(todof.ID)
	database_conn.Db.Save(&todof)
}

func FilePars(c *gin.Context){
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
			CreateTodoUrl(temp_url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}