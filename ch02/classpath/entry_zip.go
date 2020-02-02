// // version 1
// package classpath
// import "archive/zip"
// import "errors"
// import "io/ioutil"
// import "path/filepath"

// type ZipEntry stuct{
// 	absPath string 		// path that store zip document
// 	func newZipEntry(path string) *ZipEntry {...}
// 	func (self *ZipEntry) readClass(className string)([]byte,Entry,error) {...}
// 	func (self *ZipEntry) String() string {...}
// }

// func newZipEntry(path string) *ZipEntry{
// 	absDir,err := filepath.Abs(path)
// 	if err != nil{
// 		panic(err)	// filepath.Abs transform meet error
// 	}
// 	return &ZipEntry{absDir}
// }

// func (self *ZipEntry) readClass(className string)([]byte,Entry,error){
// 	r,err := zip.OpenReader(self.absPath
// 	if err != nil{
// 		return nil,nil,err
// 	}
// 	defer r.Close()
// 	for _,f := range r.File { 
// 		if f.Name == className{	//traver through documents inside zip , until find className indicated document
// 			rc,err := f.Open()
// 			if err != nil{
// 				return nil,nil,err
// 			}
// 			defer rc.Close()
// 			data,err := ioutil.ReadAll(rc)
// 			if err != nil{
// 				return nil,nil,err
// 			}
// 			return data,self,nil
// 		}
// 	}
// 	return nil,nil,errors.New("class not found: " + className)
// }

// func (self *DirEntry) String() string{
// 	return self.absDir
// }


// version 2: everty time use version1::readClass() will open and close zip file, this operation is not effective.
package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string
	zipRC   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath, nil}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if self.zipRC == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := self.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, self, err
}

// todo: close zip
func (self *ZipEntry) openJar() error {
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRC = r
	}
	return err
}

func (self *ZipEntry) findClass(className string) *zip.File {
	for _, f := range self.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (self *ZipEntry) String() string {
	return self.absPath
}
