package classpath
import(
	"os"
	"strings"
)


// different JVM use different ways to search class, oracle's JVM search class by class path. class path can be divided into 3 parts:
// part 1: bootstrap classpath, default correspond to jre/lib catalogue
// part 2: extension classpath, default correspond to jre/lib/ext catalogue
// part 3: user classpath, including self implement class and third part class, default correspond to ".", namely current catalogue.

const pathListSeparator = string(os.PathListSeparator) 	// const pathListSeparator used for store separator char

type Entry interface{
	readClass(className string)([]byte,Entry,error) 	// search and load class document, argument is class document's relative path
														// if read java.lang.Object class , should pass java/lang/Object.class argument
	String() string 	// like java toString() function, return variable string format
}

// recording different arguments to create different Entry instance 
// there are 4 types Entry interface implementation : DirEntry , ZipEntry , CompositeEntry and WildcardEntry
func newEntry(path string) Entry{
	if strings.Contains(path,pathListSeparator){
		return newCompositeEntry(path)	
	}
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	if strings,HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") || 
			strings.HasSuffix(".zip") || strings.HasSuffix(path,".ZIP"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}