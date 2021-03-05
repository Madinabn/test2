package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Student contains information about one student
type Student struct {
	ID      int
	Name    string
	Year    int
	Faculty string
	Date    string
	Books   []int
}

//Book contains information about book
type Book struct {
	IDBook   int
	Title    string
	Author   string
	Student  Student
	Headings []Heading
}

//Heading contains information about heading
type Heading struct {
	Description string
}

var students = []Student{
	{
		ID:      1,
		Name:    "Mario",
		Year:    3,
		Faculty: "Fit",
		Date:    "21.02.2021",
	},
	{
		ID:      2,
		Name:    "Beka",
		Year:    4,
		Faculty: "PMI",
		Date:    "12.02.2021",
	},
	{
		ID:      3,
		Name:    "Jasmin",
		Year:    2,
		Faculty: "ItU",
		Date:    "20.02.2021",
	},
	{
		ID:      4,
		Name:    "Saiko",
		Year:    1,
		Faculty: "Bi",
		Date:    "01.03.2021",
	},
	{
		ID:      5,
		Name:    "Berk",
		Year:    2,
		Faculty: "PMI",
		Date:    "28.02.2021",
	},
}
var headings = []Heading{
	{
		Description: "something1",
	},
	{
		Description: "something2",
	},
	{
		Description: "something3",
	},
	{
		Description: "something4",
	},
	{
		Description: "something5",
	},
	{
		Description: "something6",
	},
	{
		Description: "something7",
	},
	{
		Description: "something8",
	},
}
var books = []Book{
	{
		IDBook: 1,
		Title:  "Winnie the Pooh",
		Author: "A.A.Milne",
	},
	{
		IDBook: 2,
		Title:  "Harry Potter",
		Author: "J.K.Rowling",
	},
	{
		IDBook: 3,
		Title:  "Aiport",
		Author: "A.A.Hailey",
	},
	{
		IDBook: 4,
		Title:  "Jeeves and Woosters stories",
		Author: "P.G.Wodehouse",
	},
	{
		IDBook: 5,
		Title:  "The Adventures Of Sherlock Holmes",
		Author: "A.C.Doyle",
	},
	{
		IDBook: 6,
		Title:  "Jane Eyre",
		Author: "C.Bronte",
	},
	{
		IDBook: 7,
		Title:  "Bridget Jones' Diary",
		Author: "H.Fielding",
	},
	{
		IDBook: 8,
		Title:  "To Kill Mockingbird",
		Author: "H.Lee",
	},
}
var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"idbook": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: graphql.String,
			},
			"student": &graphql.Field{
				Type: studentType,
			},
			"headings": &graphql.Field{
				Type: graphql.NewList(headingType),
			},
		},
	},
)
var studentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"idst": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"faculty": &graphql.Field{
				Type: graphql.String,
			},
			"date": &graphql.Field{
				Type: graphql.String,
			},
			"books": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var headingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Heading",
		Fields: graphql.Fields{
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			"Book": &graphql.Field{
				Type:        bookType,
				Description: "Get Book by IDBook",
				Args: graphql.FieldConfigArgument{
					"idbook": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// Find book
						for _, book := range books {
							if int(book.IDBook) == id {
								return book, nil
							}
						}
					}
					return nil, nil
				},
			},

			"list": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "Get Full book list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return books, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	})

	fmt.Println("Server is running on port 8091")
	http.ListenAndServe(":8091", nil)

}
