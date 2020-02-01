package classpath
import "io/ioutil"
import "path/filepath"

type DirEntry stuct{
	absDir string 		// path that store absolute catalogue
	func newDirEntry(path string) *DirEntry {...}
	func (self *DirEntry) readClass(className string)([]byte,Entry,error) {...}
	func (self *DirEntry) String() string {...}
}

// default use newClass as class construct , like newDirEntry create struct instance
func newDirEntry(path string) *DirEntry{
	absDir,err := filepath.Abs(path)
	if err != nil{
		panic(err)	// filepath.Abs transform meet error
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string)([]byte,Entry,error){
	fileName := filepath.Join(self.absDir,className)
	data,err := ioutil.ReadFile(filename)
	return data,self,err
}

func (self *DirEntry) String() string{
	return self.absDir
}