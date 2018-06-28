package jsonpath

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJsonPath(t *testing.T) {
	testString := `{
		"total_devices": 5,
		"total_threads": 10,
		"shares": {
			"total": 5,
			"accepted": 5,
			"rejected": 0,
			"avg_find_time": 4,
			"tester": "work",
			"tester2": "don't want this",
			"tester3": {
				"hello":"sup",
				"fun":"money",
				"break":9
			}
		}
	}`

	jsonParser := JSONPath{
		MetricName: "jsonpather",
		TagPath:    map[string]string{"total": "$.shares.tester3"},
	}

	metrics, err := jsonParser.Parse([]byte(testString))
	assert.NoError(t, err)
	log.Printf("m[0] name: %v, tags: %v, fields: %v", metrics[0].Name(), metrics[0].Tags(), metrics[0].Fields())
	t.Error()

}
