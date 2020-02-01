package classpath
import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry stuct{
	absPath string 		// path that store zip document
	func newZipEntry(path string) *ZipEntry {...}
	func (self *ZipEntry) readClass(className string)([]byte,Entry,error) {...}
	func (self *ZipEntry) String() string {...}
}

func newZipEntry(path string) *ZipEntry{
	absDir,err := filepath.Abs(path)
	if err != nil{
		panic(err)	// filepath.Abs transform meet error
	}
	return &ZipEntry{absDir}
}

func (self *ZipEntry) readClass(className string)([]byte,Entry,error){
	r,err := zip.OpenReader(self.absPath)
	if err != nil{
		return nil,nil,err
	}
	defer r.Close()
	for _,f := range r.File { 
		if f.Name == className{	//traver through documents inside zip , until find className indicated document
			rc,err := f.Open()
			if err != nil{
				return nil,nil,err
			}
			defer rc.Close()
			data,err := ioutil.ReadAll(rc)
			if err != nil{
				return nil,nil,err
			}
			return data,self,nil
		}
	}
	return nil,nil,errors.New("class not found: " + className)
}

func (self *DirEntry) String() string{
	return self.absDir
}