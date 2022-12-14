package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	gqlrelay "github.com/sonderkevin/gql-relay"
)

func main() {
	query := `
		query {
			node{
				users {
					id
					name
				}
			}
		}
	`
	params := graphql.Params{Schema: gqlrelay.Schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute query. Errors: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
