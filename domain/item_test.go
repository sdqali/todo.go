package domain

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var oldValue string

func CanMarshalToJsonWithUrl(t *testing.T) {
	item := NewItem("test")
	bytes, _ := json.Marshal(item)
	jsonStr := string(bytes)
	if strings.Contains(jsonStr, fmt.Sprintf("/%s", item.Id)) {
		t.Errorf("Expected %s to contain url", jsonStr)
	}
}
