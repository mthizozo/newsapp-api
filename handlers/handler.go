package handlers

import (
	"Mobilebackend/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/barthr/newsapi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetHeadlines(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(err.Error()))
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in risk minder call, Please investigate")
		err := `{"risk_Data":{"faultstring":"Request body cannot be empty"}}`
		w.Header().Set("Content-Type", "application/json")

		w.Write([]byte(err))
		return

	}

	fmt.Println(string(body))
	UserInput := models.NewsData{}
	error := json.Unmarshal(body, &UserInput)

	if error == nil {
		c := newsapi.NewClient(os.Getenv("API_KEY"), newsapi.WithHTTPClient(http.DefaultClient))
		articles, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
			Country:  UserInput.Country,
		})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
        var response []byte
		response , _ = json.Marshal(articles)
		w.Write([]byte(response))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(error)
	}


}
