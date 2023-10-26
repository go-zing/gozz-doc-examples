package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/go-zing/gozz-kit/zapi"
	"github.com/go-zing/gozz-kit/zapi/zswagger"
	"github.com/go-zing/gozz-kit/zdoc"

	"github.com/go-zing/gozz-doc-examples/overview01"
)

func main() {
	swagger := zswagger.Parse(
		overview01.Apis{},
		zswagger.WithDocFunc(zdoc.TypesDoc(overview01.ZZ_types_doc).TypeFieldDoc),
		zswagger.WithHttpCast(func(api zapi.Api) zapi.HttpApi {
			sp := strings.SplitN(api.Resource, "|", 2)[:2]
			return zapi.HttpApi{
				Api:    api,
				Method: sp[0],
				Path:   sp[1],
			}
		}), zswagger.WithBindings(map[string]zswagger.Binding{
			"GET": {
				Path:  "uri",
				Query: "query",
				Body:  false,
			},
			"*": {
				Path: "uri",
				Body: true,
			},
		}))
	b, _ := json.MarshalIndent(swagger, "", "\t")
	_ = ioutil.WriteFile("swagger.json", b, 0o664)
}
