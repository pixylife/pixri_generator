package pixriLogger

import (
	"github.com/sirupsen/logrus"
	"pixri_generator/pkg/env"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var Log logrus.Logger

func init()  {
	Log  = logrus.Logger{}

	Log.SetFormatter(&logrus.TextFormatter{DisableColors: false})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	if  strings.EqualFold(env.GetLoglevel(),"info"){
		Log.SetLevel(logrus.InfoLevel)
	} else if strings.EqualFold(env.GetLoglevel(),"debug") {
		Log.SetLevel(logrus.DebugLevel)
	} else if strings.EqualFold(env.GetLoglevel(),"warning") {
		Log.SetLevel(logrus.WarnLevel)
	}else if strings.EqualFold(env.GetLoglevel(),"error"){
		Log.SetLevel(logrus.ErrorLevel)
	}

}

func Logger() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return logrus.WithField("file", filename).WithField("function", fn)
}

func IsDebugEnabled() bool  {
	if Log.IsLevelEnabled(logrus.DebugLevel){
		return true
	}
	return false
}