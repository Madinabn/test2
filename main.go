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
}

//Book contains information about book
type Book struct {
	IDBook   int
	Title    string
	Author   string
	Students []Student
	Headings []Heading
}

//Heading contains information about heading
type Heading struct {
	Description string
	Price       string
}

var books = []Book{
	{
		IDBook: 1,
		Title:  "Winnie the Pooh",
		Author: "A.A.Milne",
		Headings: []Heading{
			{
				Description: "Hello",
				Price:       "299",
			},
		},
		Students: []Student{
			{
				ID:      5,
				Name:    "Berk",
				Year:    2,
				Faculty: "PMI",
				Date:    "28.02.2021",
			},
		},
	},
	{
		IDBook: 2,
		Title:  "Harry Potter",
		Author: "J.K.Rowling",
		Headings: []Heading{
			{
				Description: "Hello Madi",
				Price:       "299",
			},
		},
		Students: []Student{
			{
				ID:      5,
				Name:    "Berk",
				Year:    2,
				Faculty: "PMI",
				Date:    "28.02.2021",
			},
		},
	},

	{
		IDBook: 3,
		Title:  "Aiport",
		Author: "A.A.Hailey",
		Headings: []Heading{
			{
				Description: "Hello Zhas",
				Price:       "299",
			},
		},
	},

	{
		IDBook: 4,
		Title:  "Jeeves and Woosters stories",
		Author: "P.G.Wodehouse",
		Headings: []Heading{
			{
				Description: "Hello",
				Price:       "299",
			},
		},
		Students: []Student{
			{
				ID:      5,
				Name:    "Berk",
				Year:    2,
				Faculty: "PMI",
				Date:    "28.02.2021",
			},
		},
	},
	{
		IDBook: 5,
		Title:  "The Adventures Of Sherlock Holmes",
		Author: "A.C.Doyle",
		Headings: []Heading{
			{
				Description: "Hello Saikal",
				Price:       "299",
			},
		},
		Students: []Student{
			{
				ID:      5,
				Name:    "Berk",
				Year:    2,
				Faculty: "PMI",
				Date:    "28.02.2021",
			},
		},
	},
	{
		IDBook: 6,
		Title:  "Jane Eyre",
		Author: "C.Bronte",
		Headings: []Heading{
			{
				Description: "Hello everyone",
				Price:       "299",
			},
		},
		Students: []Student{
			{
				ID:      5,
				Name:    "Berk",
				Year:    2,
				Faculty: "PMI",
				Date:    "28.02.2021",
			},
		},
	},
	{
		IDBook: 7,
		Title:  "Bridget Jones' Diary",
		Author: "H.Fielding",
		Headings: []Heading{
			{
				Description: "Hello hi",
				Price:       "299",
			},
		},
		Students: []Student{
			{
				ID:      5,
				Name:    "Berk",
				Year:    2,
				Faculty: "PMI",
				Date:    "28.02.2021",
			},
		},
	},
	{
		IDBook: 8,
		Title:  "To Kill Mockingbird",
		Author: "H.Lee",
		Headings: []Heading{
			{
				Description: "Hello me",
				Price:       "299",
			},
		},
		Students: []Student{
			{
				ID:      5,
				Name:    "Berk",
				Year:    2,
				Faculty: "PMI",
				Date:    "28.02.2021",
			},
		},
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
			"students": &graphql.Field{
				Type: graphql.NewList(studentType),
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
			"id": &graphql.Field{
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
		},
	},
)

var headingType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Headings",
		Fields: graphql.Fields{
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
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

	fmt.Println("Server is running on port 8087")
	http.ListenAndServe(":8087", nil)

}
