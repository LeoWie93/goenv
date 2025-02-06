package goenv

import (
	"bufio"
	"os"
	"regexp"
)

func LoadIntoEnv(filePaths ...string) (err error) {
	envMap, err := load(filePaths)
	if err != nil {
		return err
	}

	for key, value := range envMap {
		os.Setenv(key, value)
	}

	return nil
}

func GetAsMap(filePaths ...string) (envMap map[string]string, err error) {
	envMap, err = load(filePaths)
	if err != nil {
		return nil, err
	}

	return envMap, nil
}

func load(filePaths []string) (envMap map[string]string, err error) {
	envMap = make(map[string]string)

	if len(filePaths) == 0 {
		filePaths = append(filePaths, ".env")
	}

	for _, filename := range filePaths {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}

		r := regexp.MustCompile(`([A-Z\_]{0,255})=(.{0,255})`)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if matches := r.FindAllStringSubmatch(scanner.Text(), -1); matches != nil {
				envMap[matches[0][1]] = matches[0][2]
			}
		}

		file.Close()
	}

	return envMap, nil
}
