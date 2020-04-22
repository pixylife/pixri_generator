package env

////////////////////////Path/////////////////////////////////
const Root  = Lib+Src
const Src  = "/src/"
const Lib  =  "/lib"
const MODEL_PATH = "model/"
const API_PATH = "api/"

////////////////////////Suffix///////////////////
const API_SUFFIX = API_Class+ DartExtension
const MODEL_SUFFIX = ""+ DartExtension
const FormViewSuffix  = ""+DartExtension

////////////////File Extensions////////////////////
const DartExtension = ".dart"

//////////////////Data types/////////////////////
const String  = "String"
const Integer  = "Integer"

/////////////Classes////////////////////////////////
const API_Class  = "_api_service"
const UI_Form  = "_from"
const List_View  = "_listView"

//////////////////////API////////////////////////////
const ENTITY_API  ="ENTITY_API"
func GetSupportedCruds() []string {
	return []string{"Create","Delete","Update","Find","GetList"}
}

