package conf

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var configMap map[string]string

/**
*
 */
func init() {
	configMap = make(map[string]string)
	load()
}

func Get(key string) (value string) {
	value, _ = configMap[key]
	return
}

func GetOrDefault(key, defaultValue string) string {
	value, ok := configMap[key]
	if ok {
		return value
	} else {
		return defaultValue
	}
}

/**
*
 */
func load() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Errorf("Error get current dir: %s", err.Error())
	}
	f := fmt.Sprintf("%s/.env.conf", dir)
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("Error load file: %s\n", err.Error())
	}

	reader := bufio.NewReader(file)
	var key, value string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Errorf("Error read config file: %s", err.Error())
			}
			return
		}
		ss := strings.Split(line, "=")
		key, value = strings.Trim(ss[0], " "), strings.Trim(ss[1], " \n")
		configMap[key] = value
	}
}
