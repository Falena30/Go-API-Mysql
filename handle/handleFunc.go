package handle

import (
	"Go-API-SQL/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mux"
	"net/http"
	"strconv"
)

type Post struct {
	ID     int    `json:"id"`
	Tittle string `json:"tittle"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//buat variabel untuk menampung struct
	var posts []Post
	//buat variabel untuk menampung conect
	conn, err := db.Conn()
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	//buat query dan tampung
	selectData, err := conn.Query("SELECT * FROM Post")
	if err != nil {
		panic(err.Error())
	}

	//lalukan perlungan untuk menambhkan hasil dari select data kedalam posts

	for selectData.Next() {
		//buat variabel untuk menampung select data
		var post Post
		err := selectData.Scan(&post.ID, &post.Tittle)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	//lakukan json encoder
	json.NewEncoder(w).Encode(posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.Conn()
	if err != nil {
		panic(err.Error())
	}
	defer dbConn.Close()

	//inserting data which does not return any rows
	result, err := dbConn.Prepare("INSERT INTO Post(tittle) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}

	// getting data from the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyValue := make(map[string]string)
	json.Unmarshal(body, &keyValue)
	tittle := keyValue["tittle"]
	//end
	_, err = result.Exec(tittle)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "New Post Created")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	query, _ := strconv.Atoi(vars["id"])
	fmt.Println(query)

	dbconn, err := db.Conn()
	if err != nil {
		panic(err.Error())
	}
	defer dbconn.Close()

	result, err := dbconn.Query("SELECT * FROM Post WHERE ID = ?", query)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)

	//buat varibel penmapung post
	var post Post
	for result.Next() {
		err := result.Scan(&post.ID, &post.Tittle)
		if err != nil {
			panic(err.Error())
		}
	}
	//encode ke json
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	dbconn, err := db.Conn()
	if err != nil {
		panic(err.Error())
	}
	defer dbconn.Close()

	result, err := dbconn.Prepare("UPDATE Post SET Tittle = ? WHERE ID = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newTittle := keyVal["tittle"]

	_, err = result.Exec(newTittle, param["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "post dengan ID %s telah berahasil dirubah", param["id"])
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	dbconn, err := db.Conn()
	if err != nil {
		panic(err.Error())
	}
	defer dbconn.Close()

	result, err := dbconn.Prepare("DELETE FROM Post WHERE ID = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = result.Exec(param["id"])
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "Post dengan %s telah berhasil dihapus", param["id"])

}
