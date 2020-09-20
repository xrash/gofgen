package gen

var __init_file_template_head = `
package {{.packageName}}

type FileSystem struct {
db map[string][]byte
}

var FS *FileSystem = &FileSystem{
db: make(map[string][]byte),
}

func (fs *FileSystem) Get(key string) ([]byte, bool) {
v, ok := fs.db[key]
return v, ok
}

`
