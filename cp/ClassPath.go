package cp

import (
	"os"
	"path/filepath"
	"log"
	"errors"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func (classPath *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := classPath.bootClassPath.readClass(className); err == nil {
		return data, entry, nil
	}

	if data, entry, err := classPath.extClassPath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := classPath.userClassPath.readClass(className); err == nil {
		return data, entry, nil
	}
	return nil, nil, errors.New("class not found :" + className)
}

func Parse(classpathOption string) *ClassPath {
	classPath := ClassPath{}
	classPath.parseBootAndExtClassPath()
	classPath.parseUserClassPath(classpathOption)
	return &classPath
}

func (classPath *ClassPath) parseBootAndExtClassPath() {
	javaHome := getJavaHome()
	jreLibPath := filepath.Join(javaHome, "jre", "lib","*")
	classPath.bootClassPath = newEntry(jreLibPath)
	jreExtLibPath := filepath.Join(javaHome, "jre", "lib", "ext","*")
	classPath.extClassPath = newEntry(jreExtLibPath)
}

func (classPath *ClassPath) parseUserClassPath(classPathOption string) {
	if classPathOption == "" {
		classPathOption = "."
	}
	classPath.userClassPath = newEntry(classPathOption)
}

func getJavaHome() string {
	javaHome := os.Getenv("JAVA_HOME")
	if javaHome == "" {
		panic("JAVA_HOME not set !")
	}
	if path, err := filepath.Abs(javaHome); err == nil {
		log.Println("JAVA_HOME is :" + path)
		return path
	} else {
		panic(err)
	}
}
