package settings

import "github.com/joho/godotenv"

func SetupEnv() {
	err := godotenv.Load(ENVPATH)

	if err != nil {
		panic("missing environment variable file")
	}

	/**
	 * Read and Setup
	 */

}
