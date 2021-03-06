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
}

//Heading contains information about heading

var books = []Book{
	{
		IDBook: 1,
		Title:  "Winnie the Pooh",
		Author: "A.A.Milne",

		Students: []Student{
			{
				ID:      326482,
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

		Students: []Student{
			{
				ID:      293874,
				Name:    "Beka",
				Year:    4,
				Faculty: "ITU",
				Date:    "20.02.2021",
			},
		},
	},

	{
		IDBook: 3,
		Title:  "Aiport",
		Author: "A.A.Hailey",
		Students: []Student{
			{
				ID:      239084,
				Name:    "Kelly",
				Year:    1,
				Faculty: "PI",
				Date:    "10.02.2021",
			},
		},
	},

	{
		IDBook: 4,
		Title:  "Jeeves and Woosters stories",
		Author: "P.G.Wodehouse",
		Students: []Student{
			{
				ID:      438750,
				Name:    "Olga",
				Year:    5,
				Faculty: "UTS",
				Date:    "29.01.2021",
			},
		},
	},
	{
		IDBook: 5,
		Title:  "The Adventures Of Sherlock Holmes",
		Author: "A.C.Doyle",
		Students: []Student{
			{
				ID:      34875,
				Name:    "Zhasmin",
				Year:    3,
				Faculty: "IB",
				Date:    "01.03.2021",
			},
		},
	},
	{
		IDBook: 6,
		Title:  "Jane Eyre",
		Author: "C.Bronte",
		Students: []Student{
			{
				ID:      748003,
				Name:    "Bikka",
				Year:    1,
				Faculty: "PMI",
				Date:    "03.03.2021",
			},
		},
	},
	{
		IDBook: 7,
		Title:  "Bridget Jones' Diary",
		Author: "H.Fielding",
		Students: []Student{
			{
				ID:      38743,
				Name:    "Saiko",
				Year:    4,
				Faculty: "BI",
				Date:    "14.02.2021",
			},
		},
	},
	{
		IDBook: 8,
		Title:  "To Kill Mockingbird",
		Author: "H.Lee",
		Students: []Student{
			{
				ID:      37433,
				Name:    "Jonas",
				Year:    5,
				Faculty: "PI",
				Date:    "15.02.2021",
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

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//http://localhost:8092/book?query={book(idbook:2){title,author,students{id,name}}}
			"book": &graphql.Field{
				Type:        bookType,
				Description: "Get Book by IDBook",
				Args: graphql.FieldConfigArgument{
					"idbook": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["idbook"].(int)
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
			//http://localhost:8092/book?query={list{idbook,title,author,students{id,name,faculty,date,year}}}
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

	fmt.Println("Server is running on port 8092")
	http.ListenAndServe(":8092", nil)

}
