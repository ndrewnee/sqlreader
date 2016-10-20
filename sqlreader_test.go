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

func TestSqlReader_Check(t *testing.T) {
	asserts := []struct {
		files  []string
		err    error
		length int
	}{
		{
			[]string{"select", "non_exists", "not_found"},
			ErrSqlsNotFound,
			2,
		},
		{
			[]string{"select", "outer/update"},
			nil,
			0,
		},
	}

	for _, v := range asserts {
		s, err := New("sqls")
		if err != nil {
			t.Error(err)
			continue
		}

		notFound, err := s.Check(v.files...)
		if v.err != err {
			t.Errorf("Expected %v but got %v\n", v.err, err)
			continue
		}

		if v.length != len(notFound) {
			t.Errorf("Expected length: %v but got %v\n", v.length, len(notFound))
		}
	}
}
