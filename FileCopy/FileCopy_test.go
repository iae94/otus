package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileCopy(t *testing.T) {
	type args struct {
		from   string
		to     string
		offset int64
		limit  uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success", args: args{from: "temp_in.txt", to: "temp_out.txt", offset: 4, limit: 10}, wantErr:false},
		{name: "Source is not exist", args: args{from: "temp_in1.txt", to: "temp_out.txt", offset: 4, limit: 10}, wantErr:true},
		{name: "Negative offset", args: args{from: "temp_in.txt", to: "temp_out.txt", offset: -4, limit: 10}, wantErr:true},
		{name: "Offset > len(file)", args: args{from: "temp_in.txt", to: "temp_out.txt", offset: 64, limit: 10}, wantErr:true},
	}
	for _, tt := range tests {
		// Create test_in file
		testInFile, err := os.Create("temp_in.txt")
		if err != nil{
			t.Errorf("FileCopy() unable to create temp_in file: %v", err)
		}
		err = ioutil.WriteFile("temp_in.txt", []byte("Hello world!"), 0644)
		if err != nil {
			panic(err)
		}

		// Exec test
		t.Run(tt.name, func(t *testing.T) {
			if err := FileCopy(tt.args.from, tt.args.to, tt.args.offset, tt.args.limit); (err != nil) != tt.wantErr {
				t.Errorf("FileCopy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		//Close and remove test_in file
		testInFile.Close()
		err = os.Remove("temp_in.txt")

		if err != nil {
			t.Errorf("FileCopy() unable to delete temp_in file: %v", err)
		}

		//Close and remove test_out file (if exist)
		testOutFile, err := os.OpenFile(from, os.O_RDONLY, 0644)
		if err != nil {
			if os.IsExist(err) {
				testOutFile.Close()
				err = os.Remove("temp_out.txt")
				if err != nil {
					t.Errorf("FileCopy() unable to delete temp_out file: %v", err)
				}
			}
		}
	}
}
