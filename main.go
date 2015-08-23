package main

import (
"fmt"
"runtime"
"os"
"os/user"
"strings"
"io/ioutil"
)

func check(err error) {
	if err != nil {
		fmt.Print("Error: ", err)
	}
}


func main() {
	var hmdir, configPath string
	var configByte []byte
	var lines []string
	
	usr, err := user.Current()
	
	/* handle config file and the .pkgmng folder */
	configPath = usr.HomeDir + "/.pkgmngcfg"
	configByte, err = ioutil.ReadFile(configPath)
	
	/* adjust to config "pkgmng" location */
	if err != nil {
		fmt.Print("Config not found at " + configPath, ". Continue? [y/n]: ")
		tmpstr := ""
		fmt.Scan(&tmpstr)
		if tmpstr != "y" {
			os.Exit(1)
		}
		
		/* make a home folder */
		if runtime.GOOS == "windows" {
			hmdir = usr.HomeDir + "\\.pkgmng"
		} else if (runtime.GOOS == "linux" || runtime.GOOS == "darwin") {
			hmdir = usr.HomeDir + "/.pkgmng"
		}
	
		fmt.Print("Directory to place the PkgMng home folder: ")
		fmt.Scan(&hmdir)	
		err = os.Chdir(hmdir)
		check(err)
		err = os.Mkdir(".pkgmng", 0740)
		check(err)
	} else {	
		lines = strings.Split(string(configByte), "\n")
		tmp := strings.Split(lines[0], " ")
		hmdir = tmp[1]
		fmt.Print("Home directory: " + hmdir, "\n")
	}
	
	/* populate */
}