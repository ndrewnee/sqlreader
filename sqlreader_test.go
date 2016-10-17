package sqlreader

import (
	"testing"
)

func TestSqlReader_Get(t *testing.T) {
	s, err := New("sqls")
	if err != nil {
		t.Error(err)
		return
	}

	asserts := []struct {
		key      string
		expected string
	}{
		{"not_found", ""},
		{"select", "SELECT * FROM test_table WHERE id = ?"},
		{"outer/update", "UPDATE test_table SET test = TRUE"},
		{"outer/inner\\insert", "INSERT INTO some_table(name) VALUES(:name)"},
		{"dir", ""},
	}

	f := func(key, expected string) {
		for i := 0; i < 1000; i++ {
			sql := s.Get(key)

			if sql != expected {
				t.Errorf("Expected %s, got %s\n", expected, sql)
			}
		}
	}

	for _, v := range asserts {
		// Concurrency testing
		go f(v.key, v.expected)
		go f(v.key, v.expected)
	}
}
