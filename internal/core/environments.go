package core

import (
	"MovieAPI/pkg/utils"
	"errors"
)

const (
	OS_Environment_Var_Name = "ENV"
)

var (
	EnvironmentVar = map[string]string{
		"dev": "movieapi-dev",
	}
)

func GetEnvironmentVariables() string {
	return utils.EnvString(OS_Environment_Var_Name, "dev")
}

func SetEnvironment(env string) error {

	if len(EnvironmentVar[env]) == 0 {
		return errors.New("Environment value is not valid. example: use =>dev")
	}

	if err := utils.SetEnv(OS_Environment_Var_Name, env); err != nil {

		return errors.New("Setup error:" + err.Error())
	}
	return nil
}
