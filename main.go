package main

import (
	"context"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	// rego object can be prepared or evaluated
	r := rego.New(
		rego.Query("x = data.test.allow"),
		rego.Load([]string{"./test.rego"}, nil))

	//b create a prepared query that can be evaluated
	ctx := context.Background()
	query, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// load input document
	input := `{
		"servers": [
			{"id": "app", "protocols": ["https", "ssh"], "ports": ["p1", "p2", "p3"]},
			{"id": "db", "protocols": ["mysql"], "ports": ["p3"]},
			{"id": "cache", "protocols": ["memcache"], "ports": ["p3"]},
			{"id": "ci", "protocols": ["http"], "ports": ["p1", "p2"]},
			{"id": "busybox", "protocols": ["telnet"], "ports": ["p1"]}
		],
		"networks": [
			{"id": "net1", "public": false},
			{"id": "net2", "public": false},
			{"id": "net3", "public": true},
			{"id": "net4", "public": true}
		],
		"ports": [
			{"id": "p1", "network": "net1"},
			{"id": "p2", "network": "net3"},
			{"id": "p3", "network": "net2"}
		]
	}`

	// execute the prepared query
	rs, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs)
}
