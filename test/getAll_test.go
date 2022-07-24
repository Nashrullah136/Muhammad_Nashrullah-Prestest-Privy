package test

import (
	"Pretests/app"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = app.Initialize()
	log.Print("test....")
	//TODO: make migrateDB function util
	result := m.Run()
	ClearTable("cakes")
	os.Exit(result)
}

func TestGetAllWithoutData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cakes", nil)
	response := ExecuteRequest(router, req)
	err := CheckResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetAllWithData(t *testing.T) {
	var cakes = []Cake{
		{
			ID:          1,
			Title:       "Amandine",
			Description: "Chocolate layered cake filled with chocolate, caramel and fondant cream",
			Rating:      9,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Amandine_cake.jpg/800px-Amandine_cake.jpg",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:          2,
			Title:       "Amygdalopita",
			Description: "Almond cake made with ground almonds, flour, butter, egg and pastry cream",
			Rating:      7,
			Image:       "https://images.culinarybackstreets.com/wp-content/uploads/cb_athens_amygdalopita_recipe_cd_final3.jpg?lossy=1&ssl=1",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:          3,
			Title:       "Batik cake",
			Description: "A non-baked cake dessert made by mixing broken Marie biscuits, combined with a chocolate sauce or runny custard.",
			Rating:      9,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/Malaysian_batik_cake.jpg/552px-Malaysian_batik_cake.jpg",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:          4,
			Title:       "Kremówka",
			Description: "A Polish type of cream pie. It is made of two layers of puff pastry, filled with whipped cream, creamy buttercream, vanilla pastry cream (custard cream) or sometimes egg white cream",
			Rating:      11,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3f/Napoleon_cake_02.JPG/563px-Napoleon_cake_02.JPG",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:          5,
			Title:       "Gooey butter cake",
			Description: "Butter",
			Rating:      15,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/e/ed/Gooey_Pumpkin_Butter_Cake.jpg/800px-Gooey_Pumpkin_Butter_Cake.jpg",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:          6,
			Title:       "Panettone",
			Description: "Raisins, orange peel, and lemon peel",
			Rating:      19,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/Panettone_-_Nicolettone_2017_-_IMG_7085_%2831752542285%29.jpg/800px-Panettone_-_Nicolettone_2017_-_IMG_7085_%2831752542285%29.jpg",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
		{
			ID:          7,
			Title:       "Amandine",
			Description: "Chocolate layered cake filled with chocolate, caramel and fondant cream",
			Rating:      10,
			Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Amandine_cake.jpg/800px-Amandine_cake.jpg",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	for _, cake := range cakes {
		_, err := insertCake(cake)
		if err != nil {
			t.Error(err.Error())
			return
		}
	}
	req, _ := http.NewRequest("GET", "/cakes", nil)
	response := ExecuteRequest(router, req)
	ClearTable("cakes")
	err := CheckResponseCode(http.StatusOK, response.Code)
	if err != nil {
		t.Error(err.Error())
	}
	var responseStruct []Cake
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error(fmt.Errorf("Failed to read response. Error: " + err.Error()))
	}
	err = json.Unmarshal(responseByte, &responseStruct)
	if err != nil {
		t.Error(fmt.Errorf("Failed to convert response to construct. Error: " + err.Error()))
	}
	if len(responseStruct) != len(cakes) {
		t.Errorf("Doesn't get all data in database. Data should be %d, got %d data", len(cakes), len(responseStruct))
	}
	for i, cake := range responseStruct {
		if i == 0 {
			continue
		}
		if cake.Rating > responseStruct[i-1].Rating {
			t.Errorf("Rating not in sequence. Cake rating now: %f, Cake rating before: %f", cake.Rating, responseStruct[i-1].Rating)
			return
		} else if cake.Rating == responseStruct[i-1].Rating {
			if cake.Title < responseStruct[i-1].Title {
				t.Errorf("Title not in sequence. Cake title now: %s, Cake title before: %s", cake.Title, responseStruct[i-1].Title)
				return
			}
		}
	}
}
