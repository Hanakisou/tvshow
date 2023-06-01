
package api

import (
        "github.com/YanG-1989/list/blob/main/golang/douyu.go"
	"github.com/YanG-1989/list/blob/main/golang/bilibili.go"
	"github.com/YanG-1989/list/blob/main/golang/douyin.go"
	"github.com/YanG-1989/list/blob/main/golang/huya.go"
	"github.com/YanG-1989/list/blob/main/golang/youtube.go"
	"net/http"
)

func Tool(w http.ResponseWriter, r *http.Request) {
	method := r.PostFormValue("method")
	var (
		err  error
		data string
	)
	input := r.PostFormValue("input")
	switch method {
	case "sql2ent":
		data, err = sql2ent.Parse(input)
	case "sql2es":
		data, _, err = elasticsql.Convert(input)
	case "xml2json":
		b, e := xml2json.Convert(bytes.NewBufferString(input),
			xml2json.WithTypeConverter(xml2json.Float))
		data = b.String()
		err = e
	case "yaml2go":
		yaml := yaml2go.New()
		data, err = yaml.Convert([]byte(input))
	case "sql2gorm":
		res, e := parser.ParseSqlFormat(input,
			parser.WithGormType(),
			parser.WithJsonTag(),
		)
		data = string(res)
		err = e
	case "sql2gozero":
		cache := r.PostFormValue("cache")
		data, err = SqlToGoZero(input, cache)
	}

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		response.WriteError(w, err)
		return
	}
	response.WriteSuccess(w, data)
}

func SqlToGoZero(ddl, cache string) (string, error) {
	os.Setenv("HOME", "./")
	g, err := gen.NewGenerator("model", &config.Config{NamingFormat: config.DefaultFormat})
	if err != nil {
		return "", err
	}

	isCache := false
	if cache == "1" {
		isCache = true
	}

	res, err := g.GenFromDDContent([]byte(ddl), isCache, "")
	if err != nil {
		return "", err
	}

	for _, v := range res {
		return v, nil
	}
	return "", errors.New("empty")
}
