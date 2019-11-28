package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const CONFIG_FILE = "server.properties"

type ServerProperties struct {
	World struct {
		Spawning struct {
			Npcs     bool `json:"npcs"`
			Animals  bool `json:"animals"`
			Monsters bool `json:"monsters"`
		} `json:"spawning"`
		LevelName           string `json:"level-name"`
		LevelSeed           string `json:"level-seed"`
		LevelType           string `json:"level-type"`
		GenerateStructures  bool   `json:"generate-structures"`
		SpawnProtection     int    `json:"spawn-protection"`
		MaxWorldSize        int    `json:"max-world-size"`
		ViewDistance        int    `json:"view-distance"`
		BuildHeight         int    `json:"build-height"`
		AllowNether         bool   `json:"allow-nether"`
		EnableCommandBlocks bool   `json:"enable-command-blocks"`
		Gamemode            string `json:"gamemode"`
		Hardcore            bool   `json:"hardcore"`
		Difficulty          string `json:"difficulty"`
		Pvp                 bool   `json:"pvp"`
		AllowFlight         bool   `json:"allow-flight"`
	} `json:"world"`
	Server struct {
		ServerIP                   string `json:"server-ip"`
		ServerPort                 int    `json:"server-port"`
		Motd                       string `json:"motd"`
		NetworkCompressionTreshold int    `json:"network-compression-treshold"`
		UseNativeTransport         bool   `json:"use-native-transport"`
		Whitelist                  bool   `json:"whitelist"`
		EnforceWhitelist           bool   `json:"enforce-whitelist"`
		MaxPlayers                 int    `json:"max-players"`
		MaxTickTime                int    `json:"max-tick-time"`
		OnlineMode                 bool   `json:"online-mode"`
	} `json:"server"`
	Logging struct {
		Verbose bool `json:"verbose"`
	} `json:"logging"`
}

func GetServerSettings() ServerProperties {
	data, err := ioutil.ReadFile(CONFIG_FILE)
	if os.IsNotExist(err) {
		log.Fatal("Could not find server.properties file, make sure its in the same directory")
	}else if err != nil {
		panic(err)
	}
	var serverSettings = ServerProperties{}
	err = json.Unmarshal(data, &serverSettings)
	if err != nil {
		println("Error while parsing json in server.properties")
		panic(err)
	}
	return serverSettings
}