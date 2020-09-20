package gen

var __init_file_template_head = `
package {{.packageName}}

type FileManager struct {
db map[string][]byte
}

var Files *FileManager = &FileManager{
db: make(map[string][]byte),
}

func (fm *FileManager) Get(key string) ([]byte, bool) {
v, ok := fm.db[key]
return v, ok
}

`
