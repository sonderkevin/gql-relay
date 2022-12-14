package gqlrelay_test

import (
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
	gqlrelay "github.com/sonderkevin/gql-relay"
)

func TestConnection_TestFetching_CorrectlyFetchesUsers(t *testing.T) {
	query := `
        query UsersQuery {
          node (id: "1") {
			user {
				id,
				name
			}
          }
        }
      `
	expected := &graphql.Result{
		Data: map[string]interface{}{
			"user": map[string]interface{}{
				"id":   "1",
				"name": "Kevin",
			},
		},
	}

	result := graphql.Do(graphql.Params{
		Schema:        gqlrelay.Schema,
		RequestString: query,
	})

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, graphql result diff: %v", testutil.Diff(expected, result))
	}
}
