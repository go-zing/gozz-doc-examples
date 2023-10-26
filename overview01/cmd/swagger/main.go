package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-zing/gozz-kit/zapi/zswagger"
	"github.com/go-zing/gozz-kit/zdoc"

	"github.com/go-zing/gozz-doc-examples/overview01"
)

func main() {
	swagger := zswagger.Parse(
		overview01.Apis{},
		zswagger.WithDocFunc(zdoc.TypesDoc(overview01.ZZ_types_doc).TypeFieldDoc),
		zswagger.WithBindings(map[string]zswagger.Binding{
			"GET": {
				Path:  "uri",
				Query: "query",
				Body:  false,
			},
			"DELETE": {
				Path: "uri",
				Body: false,
			},
			"*": {
				Path: "uri",
				Body: true,
			},
		}))
	b, _ := json.MarshalIndent(swagger, "", "\t")
	_ = ioutil.WriteFile("swagger.json", b, 0o664)
}
