package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Logfile            string
	SqlDevelop         string
	HostDevelop        string
	PortDevelop        string
	NameDevelop        string
	UserDevelop        string
	PasswordDevelop    string
	GcsBucketNameLocal string
	GcsObjectPathLocal string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("failed to load config.ini: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Logfile:            cfg.Section("fish_go").Key("log_file").String(),
		SqlDevelop:         cfg.Section("db_development").Key("sql_develop").String(),
		HostDevelop:        cfg.Section("db_development").Key("host_develop").String(),
		PortDevelop:        cfg.Section("db_development").Key("port_develop").String(),
		NameDevelop:        cfg.Section("db_development").Key("name_develop").String(),
		UserDevelop:        cfg.Section("db_development").Key("user_develop").String(),
		PasswordDevelop:    cfg.Section("db_development").Key("password_develop").String(),
		GcsBucketNameLocal: cfg.Section("gcp").Key("gcs_bucket_name").String(),
		GcsObjectPathLocal: cfg.Section("gcp").Key("gcs_object_path").String(),
	}
}
