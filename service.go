package main

import "path/filepath"

func createService(ent entity) error {
	data := struct {
		PackageName string
	}{
		PackageName: ent.packageName,
	}
	return save(filepath.Join(ent.getDir(), "service.go"), "templates/service.tmpl", data)
}
