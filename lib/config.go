package lib

import (
	"encoding/json"
	"fmt"
	"os"
)

const filename = "envs.json"

var reservedKeys = [...]string{
	"$powershell_path",
}

type Config map[string]string

func (c Config) Get(key string) string {
	return c[key]
}

//func checkReserved(c Config) {
//	for _, key := range reservedKeys {
//		if _, ok := c[key]; ok {
//			fmt.Println("Key", key, "is reserved")
//			var hook = User32Init()
//			MessageBox(hook, 0, "Key "+key+" is reserved", "Error", MB_OK|MB_ICONSTOP|MB_SYSTEMMODAL)
//			os.Exit(1)
//		}
//	}
//}

func configExists() bool {
	_, err := os.Stat(filename)
	return err == nil
}

func createConfigFile() {
	config := Config{}

	var powershellPath = GetPowershellInstallation().Path
	config["$powershell_path"] = powershellPath

	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("envs.json created successfully")

	var hook = User32Init()
	MessageBox(hook, 0, "envs.json created successfully", "Success",
		MB_OK|
			MB_ICONINFORMATION|
			MB_SYSTEMMODAL,
	)
	os.Exit(0)
}

func GetConfig() Config {
	var config Config

	if !configExists() {
		createConfigFile()
		return config
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	//checkReserved(config)

	return config
}
