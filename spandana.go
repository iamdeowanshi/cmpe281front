package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/unrolled/render"
	//"strconv"
	//"os"
	"log"

	"github.com/gorilla/mux"
	"github.com/mediocregopher/radix.v2/redis"
	//"io/ioutil"
	//"github.com/go-martini/martini"
	//"github.com/martini-contrib/binding"
	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/handlers"
	//"github.com/rs/cors"
)

type burgerData struct {
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

//Redis connection String
//var conn,err= redis.Dial("tcp", "18.144.9.107:6379")
var conn, err = redis.Dial("tcp", "localhost:6379")

//Test Redis connection
func Connection() {
	if err != nil {
		log.Fatal("failed to connect to Redis")
		log.Fatal(err)
	}
	//pong, err := client.Ping().Result()

	fmt.Println("Redis is now connected", conn)
	defer conn.Close()
}

//list of all the crud handlers
func initRoutes(mx *mux.Router, formatter *render.Render) {

	mx.HandleFunc("/ping", ping(formatter)).Methods("GET")
	//display Burger item
	mx.HandleFunc("/displayitem", getItem(formatter)).Methods("GET")
	//Create a new Burger item
	mx.HandleFunc("/createitem", createItem(formatter)).Methods("POST")
	mx.HandleFunc("/createitem", createItem(formatter)).Methods("OPTIONS")
	//update Burger item
	mx.HandleFunc("/item", updateItem(formatter)).Methods("PUT")
	mx.HandleFunc("/item", createItem(formatter)).Methods("OPTIONS")
	//delete Burger item
	mx.HandleFunc("/item", deleteItem(formatter)).Methods("DELETE")
	mx.HandleFunc("/item", createItem(formatter)).Methods("OPTIONS")
}

// func corsHandler(h http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 	  if (r.Method == "OPTIONS" || r.Method == "POST") {
// 		//handle preflight in here
// 	  } else {
// 		h.ServeHTTP(w,r)
// 	  }
// 	}
//   }

//Ping handler
func ping(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("Server running on port 30035")
		formatter.JSON(w, http.StatusOK, struct{ Burger string }{"Welcome to Counter Burger!"})
	}
}

//Display the burger
func getItem(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		if req.Method == "OPTIONS" {
			//fmt.Println("Ignore Options")
			return
		}
		//var m burgerData

		//var conn,err= redis.Dial("tcp", "18.144.9.107:6379")
		conn, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		fmt.Println("Server running on port 30035")

		//var getTitle="title"
		allburgers := conn.Cmd("HGETAll", "Burger")
		l, _ := allburgers.List()
		fmt.Println("List is", l)

		// get all the values in a list
		/*	for _,elemstr:= range l{
			fmt.Println("value is ",elemstr)
		} */
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("Burger data is:",title[0])

		formatter.JSON(w, http.StatusOK, l)
	}
}

//Create burger
func createItem(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		if req.Method == "OPTIONS" {
			//fmt.Println("Ignore Options")
			return
		}

		var m burgerData
		err = json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			fmt.Println("Error in Decoding the data")
			panic(err)
		}

		fmt.Println("title is :", m.Title)
		fmt.Println("price is :", m.Price)
		fmt.Println("description is :", m.Description)

		//conn,err:= redis.Dial("tcp", "18.144.9.107:6379")
		var conn, err = redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		fmt.Println("Value in code:", m.Code)
		result := conn.Cmd("HMSET", "Burger", m.Code+"Code", m.Code, m.Code+"Title", m.Title, m.Code+"Price", m.Price, m.Code+"Description", m.Description)
		//result:= conn.Cmd("HMSET", "album:3", "testjson", string(body), "testjson1", "helloworld1" )
		//result:= conn.Cmd("HMSET", "album:2", "title", "SPROUTED VEGGIE", "price", 4.95, "description", "housemade vegan veggie,organic mixed greens,red onions")
		fmt.Println("Created New Item", result)
		var response string = "Created as new Burger Recipe!"
		formatter.JSON(w, http.StatusOK, struct {
			Item string `json:"Item,omitempty"`
		}{response})
	}
}

//update burger
func updateItem(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")

		if req.Method == "OPTIONS" {
			//fmt.Println("Ignore Options")
			return
		}

		var m burgerData
		err = json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			fmt.Println("Error in Decoding the data")
			panic(err)
		}

		//conn,err:= redis.Dial("tcp", "18.144.9.107:6379")
		var conn, err = redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		//result:= conn.Cmd("HMSET", "Burger", "title", "SPROUTED VEGGIE", "price", 5.95, "description", "housemade vegan veggie,organic mixed greens,red onions")
		// fmt.Println("Updated Price of Burger",result)
		val := conn.Cmd("HEXISTS", "Burger", m.Code+"Title")
		//var resultset,_ = val.List()
		//var resultlen= len(resultset)

		if val.Err != nil {
			fmt.Println("There is no so key to update", val.Err)
			panic(val.Err)
		} else {
			result := conn.Cmd("HMSET", "Burger", m.Code+"Price", m.Price)
			fmt.Println("After Price Update", result)
			var response string = "Updated Price of Burger"
			formatter.JSON(w, http.StatusOK, struct {
				Item string `json:"Item,omitempty"`
			}{response})
		}
	}
}

//Delete Burger
func deleteItem(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")
		if req.Method == "OPTIONS" {
			return
		}
		var m burgerData
		err = json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			fmt.Println("Error in Decoding the data")
			panic(err)
		}

		//conn,err:= redis.Dial("tcp", "18.144.9.107:6379")
		conn, err = redis.Dial("tcp", "localhost:6379")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		val := conn.Cmd("HEXISTS", "Burger", m.Code+"Title")
		//result,_:=val.List()
		//var length=len(result)
		if val.Err != nil {
			fmt.Println("There is no so key to update", val.Err)
			panic(val.Err)
			//panic(val.Err)
		} else {
			result := conn.Cmd("HDEL", "Burger", m.Code+"Code", m.Code+"Title", m.Code+"Price", m.Code+"Description")
			fmt.Println("After item deleted", result)
			var response string = "This Burger will longer be available in the Catalog"
			formatter.JSON(w, http.StatusOK, struct {
				Item string `json:"Item,omitempty"`
			}{response})
		}
	}
}

func main() {
	Connection()
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	log.Fatal(http.ListenAndServe(":3305", mx))

}
