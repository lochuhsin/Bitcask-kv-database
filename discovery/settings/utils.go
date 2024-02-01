package settings

func SetupConfig(envPaths ...string) {
	Config = NewConfiguration(
		envPaths,
		SetClusterMemberCount(),
		SetHttpPort(),
	)
}
