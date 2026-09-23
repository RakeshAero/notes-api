package models
import (
	"time"
	"fmt"
	"encoding/json"
)

type User struct {
	id   		int 		`json:"id"`
	name 		string 		`json:"username"`
	email 		string 		`json:"email"`
	password 	string 		`json:"-"`
	created_at  time.Time   `json:"created_at"`

}

type Note struct {
	id 			int 		`json:"id"`
	user_id 	int  		`json:"user_id"`
	title 		string 		`json:"title"`
	content 	string  	`json:"content"`
	created_at  time.Time 	`json:"created_at"`
	updated_at  time.Time  	`json:"updated_at"`
}