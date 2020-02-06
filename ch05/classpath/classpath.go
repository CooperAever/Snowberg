package classpath
import "os"
import "path/filepath"

type Classpath struct{
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

// Parse() function use -Xjre option parse bootClasspath and extClasspath , and use -classpath/-cp option to parse userClasspath
func Parse(jreOption,cpOption string) *Classpath{
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	jreDir := getJreDir(jreOption)	//get jreDir
	jreLibPath := filepath.Join(jreDir,"lib","*") 	// like jre/lib/*
	self.bootClasspath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir,"lib","ext","*")	// like jre/lib/ext/*
	self.extClasspath = newWildcardEntry(jreExtPath)
}


// if user not provide -classpath/-cp option, use current folder as userclasspath 
func (self *Classpath) parseUserClasspath(cpOption string){
	if cpOption == ""{
		cpOption="."
	}
	self.userClasspath = newEntry(cpOption)
}


// first use input -Xjre option as jre catalogue.If not given, then try use current catalogue search for jre catalogue.
// If jre not exist in current catalogue, try JAVA_HOME environment.
func getJreDir(jreoption string) string{
	if jreoption != "" && exists(jreoption){
		return jreoption
	}
	if(exists("./jre")){
		return "./jre"
	}
	if jh:= os.Getenv("JAVA_HOME");jh != ""{
		return filepath.Join(jh,"jre")
	}
	panic("Can not find jre folder")
}

//exists() used for judge if there exist this folder
func exists(path string) bool{
	// golang use os.Stat() return bool value to judge whether this folder exist
	// if return value == nil , means folder exist
	// if return value use as os.IsNotExist()'s argument and return true , means folder not exist
	if _,err := os.Stat(path); err != nil{
		if os.IsNotExist(err){
			return false
		}
	}
	return true
}

// ReadClass() search class document in bootpath、extpath、userpath successively
func (self *Classpath) ReadClass(className string) ([]byte,Entry,error){
	className = className + ".class"
	if data,entry,err := self.bootClasspath.readClass(className);err == nil{
		return data,entry,err
	}
	if data,entry,err := self.extClasspath.readClass(className);err == nil{
		return data,entry,err
	}
	return self.userClasspath.readClass(className)
}

// String() return userClasspath's string format
func (self *Classpath) String() string{
	return self.userClasspath.String()
}


