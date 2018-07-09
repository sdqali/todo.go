package todo

import (
	"os"
	"testing"
	"todo"
)

const testFilePath string = "/tmp/test.json"

func TestCanAddItems(t *testing.T) {
	store := todo.NewJsonFileStore(testFilePath)
	store.Add(todo.NewItem("test"))
	count := len(store.All())
	if count != 1 {
		t.Errorf("Expected store to have 1 item, but it had %d items.", count)
	}
}

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	os.Remove(testFilePath)
}

func tearDown() {
	os.Remove(testFilePath)
}
