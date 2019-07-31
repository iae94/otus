package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestEnvCustomer(t *testing.T) {
	type args struct {
		envDir       string
		command []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success", args: args{envDir: "TempDirSuccess", command: []string{"env"}}, wantErr:false},
		{name: "Env dir is not directory", args: args{envDir: "TempDirIsNotDir.txt", command: []string{"env"}}, wantErr:true},
		{name: "Env dir is empty", args: args{envDir: "TempDirEmpty", command: []string{"env"}}, wantErr:true},
		{name: "Command is missed", args: args{envDir: "TempDirEmpty", command: []string{}}, wantErr:true},
		{name: "Ping", args: args{envDir: "TempDirSuccess", command: []string{"ping", "127.0.0.1"}}, wantErr:false},
		{name: "Exit status 1", args: args{envDir: "TempDirSuccess", command: []string{"ping"}}, wantErr:true},
	}
	//Create test dirs
	err := os.Mkdir("TempDirSuccess", 0777)
	if err != nil {
		if _, err := os.Stat("TempDirSuccess"); os.IsNotExist(err) {
			if !os.IsExist(err) {
				t.Errorf("Cannot create TempDirSuccess directory %v", err)
			}
		}
	}
	err = ioutil.WriteFile("TempDirSuccess\\var1.txt", []byte("123"), 0644)
	err = ioutil.WriteFile("TempDirSuccess\\var2.txt", []byte("hello"), 0644)
	defer func() {fmt.Println("Removing TempDirSuccess"); os.RemoveAll("TempDirSuccess")}()

	err = os.Mkdir("TempDirEmpty", 0777)
	if err != nil {
		if _, err := os.Stat("TempDirEmpty"); os.IsNotExist(err) {
			if !os.IsExist(err) {
				t.Errorf("Cannot create TempDirEmpty directory %v", err)
			}
		}
	}
	defer func() {fmt.Println("Removing TempDirEmpty"); os.RemoveAll("TempDirEmpty")}()

	err = ioutil.WriteFile("TempDirIsNotDir.txt", []byte("Some content"), 0644)
	defer func() {fmt.Println("Removing TempDirIsNotDir.txt"); os.Remove("TempDirIsNotDir.txt")}()

	//Exec tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EnvCustomer(tt.args.envDir, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("EnvCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
