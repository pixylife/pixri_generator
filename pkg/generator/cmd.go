package generator

import ("os"
        "pixri_generator/pixriLogger")

func GenerateDir(path string){
	if _, err := os.Stat(path); os.IsNotExist(err) {
		pixriLogger.Log.Info("Creating "+path)
		e := os.Mkdir(path, os.ModePerm)
		if e != nil {
			pixriLogger.Log.Error("Failed to create "+path,e)
		}
	}
}