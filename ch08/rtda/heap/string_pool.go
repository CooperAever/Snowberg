package heap
import "unicode/utf16"

// to save memory, JVM store a string pool
var internedStrings = map[string] *Object{} 	//key is go string, value is java string

// according to go string, return java string instance
func JString(loader *ClassLoader,goStr string) *Object{
	if internedStr, ok := internedStrings[goStr];ok{
		return internedStr
	}
	chars := stringToUtf16(goStr) 	// transform UTF8 to java array UTF16
	jChars := &Object{loader.LoadClass("[C"),chars} 	
	jStr := loader.LoadClass("java/lang/String").NewObject() //create a new instance 
	jStr.SetRefVar("value","[C",jChars) 	// set value
	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf8 -> utf16
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}

