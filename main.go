package main

import (
	"fmt"
	"io/ioutil"

	"os"
	"sort"
	"strings"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

type List struct {
	Name []string
}

func init() {
	// Log as JSON instead of the default ASCII formatter, but wrapped with the runtime Formatter.
	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	// Enable line number logging as well
	formatter.Line = true

	// Replace the default Logrus Formatter with our runtime Formatter
	log.SetFormatter(&formatter)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the info severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("start_name_sorter")

	//read the file-name.txt
	randomNamefile := "file-name.txt"

	names, err := ReadFile(randomNamefile)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("main_read_text")
	}

	//Get LastName and assign to map
	mapName := make(map[string]string)
	for _, name := range names {

		i := strings.Split(name, " ")

		countI := len(i)

		//key[lastname, firstname]
		lastFirstName := i[countI-1] + i[0]

		mapName[lastFirstName] = name
	}

	//sort the Lastname
	var keys []string
	for k := range mapName {
		keys = append(keys, k)
	}

	sort.Sort(sort.StringSlice(keys))

	//create a new file file-name.txt
	err = WriteTextFile(randomNamefile, keys, mapName)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("main_write_text")
	}

	log.Info("end_name_sorter")

}

// Open file and parse names
func ReadFile(path string) (list []string, err error) {

	txtFile, err := os.Open(path)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("readfile_open_text")
		return list, err
	}

	byteValue, err := ioutil.ReadAll(txtFile)
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("readfile_read_text")
		return list, err
	}

	names := strings.Replace(string(byteValue), "\n", ",", -1)

	arrNames := strings.Split(names, ",")

	list = arrNames

	return list, nil
}

func WriteTextFile(randomNamefile string, keys []string, mapName map[string]string) error {
	//create a new file file-name.txt
	sortedNameFile, err := os.Create("file-name-sorted.txt")
	if err != nil {
		log.WithFields(log.Fields{"error": err.Error()}).Error("writetextfile_create_text")
	}

	_, err = sortedNameFile.WriteString("name-sorter " + randomNamefile + "\n")
	_, err = sortedNameFile.WriteString("\n")
	for _, k := range keys {
		fmt.Println(mapName[k])
		_, err = sortedNameFile.WriteString(mapName[k])
	}

	return nil
}
