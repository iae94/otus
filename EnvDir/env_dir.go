package main
//Эта утилита позволяет запускать программы получая переменные окружения из определенной директории.
//Пример использования:
//```
//go-envdir /path/to/env/dir some_prog
//```
//Если в директории /path/to/env/dir содержатся файлы
//* `A_ENV` с содержимым `123`
//* `B_VAR` с содержимым `another_val`
//То программа `some_prog` должать быть запущена с переменными окружения `A_ENV=123 B_VAR=another_val`

import (
	"errors"
	"fmt"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

var envDir string
var command []string

func init() {
	flag.StringVarP(&envDir, "envdir", "e", "", "Directory with env files")
	flag.StringSliceVarP(&command, "command", "c", []string{}, "Command name")
}

func EnvCustomer(envDir string, command []string) (err error){
	var cmdEnv []string
	//Check command was specified
	if len(command) == 0{
		err = errors.New("command (-c flag) is missed!")
		log.Println(err)
		return err
	}

	//If envdir was specified
	if envDir != "" {
		//Check envDir is exist and is directory
		dir, err := os.Stat(envDir)
		if os.IsNotExist(err) {
			err = errors.New(fmt.Sprintf("Directory %v is not exist!", envDir))
			log.Println(err)
			return err
		}
		if !dir.IsDir() {
			err = errors.New(fmt.Sprintf("%v is not direcrtory!", envDir))
			log.Println(err)
			return err
		}
		//Read all files
		files, err := ioutil.ReadDir(envDir)
		if err != nil {
			log.Println(err)
			return err
		}
		if len(files) == 0 {
			err = errors.New(fmt.Sprintf("Directory %v is empty!", envDir))
			log.Println(err)
			return err
		}
		for _, f := range files {
			filePath := path.Join(envDir, f.Name())
			fileInfo, err := os.Stat(filePath)
			if err != nil {
				log.Println(err)
				continue
			}
			//Check file not Dir
			if !fileInfo.IsDir() {
				envValue, err := ioutil.ReadFile(filePath)
				if err != nil {
					log.Println(err)
					continue
				} else {
					//Read file content and append to cmdEnv
					value := strings.TrimSuffix(fileInfo.Name(), filepath.Ext(fileInfo.Name())) + "=" + string(envValue)
					cmdEnv = append(cmdEnv, value)
				}
			}
		}
	}
	//Set stdout and stderr
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//Set new env variables if cmdEnv is not empty
	if len(cmdEnv) > 0{
		cmd.Env = cmdEnv
	}
	
	//Run command
	err = cmd.Start()
	err = cmd.Wait()
	fmt.Printf("Command: %v exit with error: %v\n", command, err)
	return err
}
