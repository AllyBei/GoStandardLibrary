package main

import (
	"log"
	"os"

	//"file"
	"fmt"
	"time"
)

func main() {
	//open test.txt
	file, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	// Count bytes and print content of test.txt
	data := make([]byte, 500)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	//change the permision of test.txt
	err = os.Chmod("test", 0770)
	if err != nil {
		log.Fatal(err)
	}

	//change file owner with UID and GID, -1 means not change
	err = os.Chown("test", -1, -1)
	if err != nil {
		log.Fatal(err)
	}

	//change file's access time and modify time
	aTime := time.Date(1970, time.January, 2, 4, 5, 6, 0, time.UTC)
	mTime := time.Date(1970, time.January, 1, 3, 4, 5, 0, time.UTC)
	err = os.Chtimes("test", aTime, mTime)
	if err != nil {
		log.Fatal(err)
	}

	//read env var
	envVar := os.Environ()
	for _, v := range envVar {

		fmt.Println(v)

		//fail to input env var into file
		/*
			bs := []byte(v)
			bs = append(bs, []byte("\n")...)
			_, err = file.Write(bs)
			if err != nil {
				log.Fatal(err)
			}
		*/
	}

	//print the path of exec programe
	exec, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Exec:", exec)

	//exit program with code
	//os.Exit(100)

	//Expand replaces ${var} or $var in the string based on the mapping function
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	fmt.Println(os.Expand("Good ${DAY_PART}, ${NAME}!", mapper))

	//set 2 env var
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	//test if env var added sucessfully
	for _, v := range os.Environ() {
		switch v {
		case "NAME=gopher":
			fmt.Println(v)
		case "BURROW=/usr/gopher":
			fmt.Println(v)
		}

	}

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

	//get Group ID
	groupID := os.Getegid()
	fmt.Println("Group ID:", groupID)

	//get env with key name. if not exit, return ""
	goPath := os.Getenv("GOPATH")
	sysPath := os.Getenv("PATH")
	fmt.Println("GOPATH:", goPath)
	fmt.Println("PATH:", sysPath)

	//get effective UID / GID / Caller Groups / Page Size / pid / parent pid / UID  of caller
	euid := os.Geteuid()
	gid := os.Getgid()
	callerGroup, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	}
	pageSize := os.Getpagesize()
	pid := os.Getpid()
	parentPid := os.Getppid()
	uid := os.Getuid()
	fmt.Printf("effective UID: %d\nGID: %d\ncallerGroup: %v\npageSize: %d\npid: %d\nparentPid: %d\nUID: %d\n", euid, gid, callerGroup, pageSize, pid, parentPid, uid)

	//get root path
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("root path:", rootPath)

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hostname:", hostname)

	filenames := []string{"a-nonexistent-file", "test", "os.go", "plan"}
	for _, filename := range filenames {
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println(filename, "file does not exist")
		} else {
			fmt.Println(filename, "file exists")
		}
	}

}
