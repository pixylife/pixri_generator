package env

////////////////////////Path/////////////////////////////////
const MODEL_PATH = "/src/model/"
const API_PATH = "/src/api/"

////////////////////////Suffix///////////////////
const API_SUFFIX = "_api_service"+ DartExtension
const MODEL_SUFFIX = ""+ DartExtension
const FormViewSuffix  = ""+DartExtension

////////////////File Extensions////////////////////
const DartExtension = ".dart"

//////////////////Data types/////////////////////
const String  = "String"
const Integer  = "Integer"


//////////////////////API////////////////////////////
const ENTITY_API  ="ENTITY_API"
func GetSupportedCruds() []string {
	return []string{"Create","Delete","Update","Find","GetList"}
}

