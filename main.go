package main

import "fmt"

func main() {
	query := `
# query_metadata_namespace: demo
# query_metadata_task_type: demo2
# odin_service: demo3
`
	fmt.Println(NewQueryMetadataMatcher().MatchQuery(query))
}
