package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sqweek/dialog"
	"gopkg.in/ini.v1"
)
func GetPath() (path_src, path_dest string) {
	src, err := dialog.Directory().Title("Select the source folder").Browse()

	if err != nil {
		fmt.Println("error while selecting source folder")
		os.Exit(1)
	}

	dest, err := dialog.Directory().Title("Select the destination folder").Browse()

	if err != nil {
		fmt.Println("error while selecting destination folder")
		os.Exit(1)
	}

	return src, dest
}

func WritePathINI(src_path string, dest_path string){

	cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

	saved_src := cfg.Section("paths").Key("src").String()
	saved_dest := cfg.Section("paths").Key("dest").String()

	cfg.Section("paths").Key("src").SetValue(saved_src + src_path)
	cfg.Section("paths").Key("dest").SetValue(saved_dest + dest_path)
	cfg.SaveTo("config.ini")

	fmt.Println("Saved succefully")
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
			return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
			return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
			return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
			return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// func CopyDir(source, destination string) error {
//     var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
//         var relPath string = strings.Replace(path, source, "", 1)
//         if relPath == "" {
//             return nil
//         }
//         if info.IsDir() {
// 			if _, err := os.Stat(filepath.Join(destination, relPath)); os.IsNotExist(err) {
// 				// fmt.Println("backing up " + filepath.Join(destination, relPath))
// 				return os.MkdirAll(filepath.Join(destination, relPath), 0755)
// 				// return os.MkdirAll(filepath.Join(destination, relPath), os.ModePerm)
// 			}else{
// 				fmt.Println(relPath + " already exists")
// 				return err
// 			}
//             // return os.Mkdir(filepath.Join(destination, relPath), 0777)
//         } else {
//             var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
//             if err1 != nil {
//                 return err1
//             }
// 			if _, err2 := os.Stat(filepath.Join(destination, relPath)); os.IsNotExist(err) {
// 				// fmt.Println("backing up " + filepath.Join(destination, relPath))
// 				return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
// 			}else{
// 				fmt.Println(relPath + " already exists")
// 				return err2
// 			}
            
//         }
//     })
//     return err
// }


func CopyDir(source, destination string) error {
    var err error = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
        var relPath string = strings.Replace(path, source, "", 1)
        if relPath == "" {
            return nil
        }
        if info.IsDir() {
            return os.MkdirAll(filepath.Join(destination, relPath), 0755)
        } else {
            var data, err1 = ioutil.ReadFile(filepath.Join(source, relPath))
            if err1 != nil {
                return err1
            }

			if _, err2 := os.Stat(filepath.Join(destination, relPath)); os.IsNotExist(err2) {
				fmt.Println("copying " + filepath.Join(destination, relPath))
				return ioutil.WriteFile(filepath.Join(destination, relPath), data, 0777)
			}else{
				fmt.Println(filepath.Join(destination, relPath) + " already exists")
				return err2
			}
			
			
            
        }
    })
    return err
}

var clear map[string]func()
func Clear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    }
}

func FilePathWalkDir(root string) ([]string, error) {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}