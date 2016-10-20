package sqlreader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"errors"
)

var (
	ErrSqlsNotFound = errors.New("Required sqls not found")
)

// New is just a constructor of SqlReader
// E.g. sqlReader, err := New("path")
// @param path is a path with sql files inside
// Reads all sql files once, then you can use method sqlReader.Get to get sql
func New(path string) (sqlReader *SqlReader, err error) {
	sqlReader = &SqlReader{
		path:      path,
		filePaths: make([]string, 0),
		sqlFiles:  make(map[string]string),
	}

	err = sqlReader.read()
	return
}

// SqlReader is a store of sql files saved in memory as map
type SqlReader struct {
	path      string
	filePaths []string

	sqlFiles map[string]string
}

func (s *SqlReader) Check(required ...string) (notFoundSqls []string, err error) {
	for _, sql := range required {
		_, ok := s.sqlFiles[sql]
		if !ok {
			notFoundSqls = append(notFoundSqls, sql)
		}
	}

	if len(notFoundSqls) > 0 {
		err = ErrSqlsNotFound
	}

	return
}

// Gets sql string by key. Key is a path of file without root directory and extension
// E.g. if file saved as "/path/some_dir/insert.sql"
// Then use sqlReader.Get("some_dir/insert") or sqlReader.Get("some_dir\insert") to get this sql
// You can use "/", "\" as separators, no matter
func (s *SqlReader) Get(key string) (sql string) {
	// Ignoring os specific path separators
	key = filepath.FromSlash(key)

	sql = s.sqlFiles[key]
	return
}

// Recursively reads directory and add each sql file to map
func (s *SqlReader) read() (err error) {
	err = filepath.Walk(s.path, s.findSqlFiles)
	if err != nil {
		return
	}

	var fileBytes []byte
	for _, path := range s.filePaths {
		fileBytes, err = ioutil.ReadFile(path)
		if err != nil {
			return
		}

		// Clearing path, e.g. ../../sqls/object/insert.sql -> /sqls/object/insert.sql
		cleanPath := filepath.Clean(path)
		// Removing root path and separator, e.g. /sqls/object/insert.sql -> object/insert.sql
		begin := len(s.path) + 1
		// Removing extension, e.g. object/insert.sql -> object/insert
		end := len(cleanPath) - len(filepath.Ext(cleanPath))
		// We will use cleaned names as keys
		key := cleanPath[begin:end]

		s.sqlFiles[key] = string(fileBytes)
	}

	return
}

// walkerFunc that finds all sql files in directory and add it to slice
func (s *SqlReader) findSqlFiles(path string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	matched, err := filepath.Match("*.sql", info.Name())
	if err != nil {
		return err
	}

	if matched {
		s.filePaths = append(s.filePaths, path)
	}

	return nil
}
