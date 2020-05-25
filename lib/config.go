package lib

import (
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"path"
	"strings"
)

func ReadConfig() error {
	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("home directory not found: %w", err)
	}

	viper.SetConfigPermissions(0600)
	viper.SetConfigName(".kb")
	viper.AddConfigPath(home)
	viper.SetEnvPrefix("KB")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	_ = viper.BindEnv("server.user")
	_ = viper.BindEnv("server.password")
	_ = viper.BindEnv("server.url")

	err = viper.ReadInConfig()
	if err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			err = viper.WriteConfigAs(path.Join(home, ".kb.yaml"))
			if err != nil {
				return err
			}
			err = viper.ReadInConfig()
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("couldn't read config file %s: %w\n", viper.ConfigFileUsed(), err)
		}
	}

	return nil
}

func WriteConfig() error {
	return nil
}
