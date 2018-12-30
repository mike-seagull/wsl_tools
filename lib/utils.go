package lib

import (
	"os"
	"runtime"
	"path/filepath"
	"strings"
)

var UNIX_C string = os.Getenv("UNIX_C")
var UNIX_HOME string = os.Getenv("UNIX_HOME")
var WINDOWS_C string = os.Getenv("WINDOWS_C")
var WINDOWS_HOME string = os.Getenv("WINDOWS_HOME")

func Home() string {
	// https://gist.github.com/miguelmota/1fb071b6c8920d1a6ed27c957b1a489a
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}
func Convert(file_fullpath string, to_windows bool) string{
    if to_windows {
    	file_fullpath = strings.Replace(file_fullpath,"\\", "/", -1)
        if strings.Contains(file_fullpath, UNIX_HOME) {
            return strings.Replace(file_fullpath, UNIX_HOME, WINDOWS_HOME, 1)
        } else {
            return strings.Replace(file_fullpath, UNIX_C, WINDOWS_C, 1)
        }
    } else {
        if strings.Contains(file_fullpath, WINDOWS_HOME) {
            file_fullpath = strings.Replace(file_fullpath, WINDOWS_HOME, UNIX_HOME, 1)
        } else {
            file_fullpath = strings.Replace(file_fullpath, WINDOWS_C, UNIX_C, 1)
        }
		return strings.Replace(file_fullpath, "/", "\\", -1)
    }
}
func GetFullPath(file string) string {
	if string(file[0]) != "/" {
	    if string(file[0]) == "~" && len(file) == 1 { // "~"
	        file = Home()
	    } else if string(file[0]) == "." && len(file) == 1 { // "."
	    	file, _ = os.Getwd()
	    } else if string(file[0]) + string(file[1]) == "~/" { // "~/$file"
	        file = filepath.Join(Home(), "/", file[2:])
	    } else if string(file[0]) + string(file[1]) == "./" { // "./$file"
	    	cwd, _ := os.Getwd()
	        file = filepath.Join(cwd, file[2:])
	    } else {
	    	cwd, _ := os.Getwd()
	        file = filepath.Join(cwd, file)
	    }
	}
	return file
}