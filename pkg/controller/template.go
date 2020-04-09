package controller

import ("os"
	"path/filepath"
	"pixri_generator/pixriLogger"
	"text/template"
)


func TemplateFileWriter(data interface{},path string,tmpl *template.Template){
	f, err := os.Create(filepath.FromSlash(path))
	if err != nil {
		pixriLogger.Log.Error("Failed to create "+path,err)
		return
	}
	err = tmpl.Execute(f, data)
	if err != nil {
		pixriLogger.Log.Error("Failed to execute template ",err)
		return
	}
	f.Close()
}

func TemplateFileWriterByName(data interface{},path string,tmpl *template.Template,name string){
	f, err := os.Create(filepath.FromSlash(path))
	if err != nil {
		pixriLogger.Log.Error("Failed to create "+path,err)
		return
	}
	err = tmpl.ExecuteTemplate(f,name,data)
	if err != nil {
		pixriLogger.Log.Error("Failed to execute template ",err)
		return
	}
	f.Close()
}






