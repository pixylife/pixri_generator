package controller

import ("os"
	"pixri_generator/pixriLogger"
	"text/template"
)


func TemplateFileWriter(data interface{},path string,tmpl *template.Template){
	f, err := os.Create(path)
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






