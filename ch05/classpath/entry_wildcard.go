package classpath
import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry{
	baseDir := path[:len(path)-1]  // remove *
	compositeEntry := []Entry{}

	walkFn := func(path string,info os.FileInfo,err error) error{
		if err != nil{
			return err
		}
		// return skipDir to skip child catalogue
		if info.IsDir() && path != baseDir{
			return filepath.SkipDir
		}
		// in wlakFn, choose JAR document according to Suffix a
		if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR"){
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir,walkFn)	//use Walk() function travel baseDir , the second argument is a function
	return compositeEntry 

}