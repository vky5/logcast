// To Load env to the entire project

package utils

import "github.com/joho/godotenv"

func LoadEnv(fileName string) error {
	err := godotenv.Load(fileName)
	return FailedOnError(err, "[Env]", "Failed to laod the env file")
}
