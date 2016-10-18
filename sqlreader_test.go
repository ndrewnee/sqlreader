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
		{"outer/inner/insert", "INSERT INTO some_table(name) VALUES(:name)"},
		{"dir", ""},
	}

	for _, v := range asserts {
		sql := s.Get(v.key)

		if v.expected != sql {
			t.Errorf("Expected %s, got %s\n", v.expected, sql)
		}
	}
}
