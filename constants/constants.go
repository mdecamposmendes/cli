package constants

const (
	VshRootPath               = "/usr/local/valet-sh"
	VshBasePath               = VshRootPath + "/valet-sh"
	VshVenvPath               = VshRootPath + "/venv"
	VshEtcPath                = VshRootPath + "/etc"
	VshAnsibleFactsFile       = "/tmp/ansible-facts/local"
	VshServiceFile            = VshEtcPath + "/services.yml"
	VshBundlesFile            = VshEtcPath + "/bundles.yml"
	VshReleaseChannelFile     = "RELEASE_CHANNEL"
	VshReleaseChannelFilePath = VshEtcPath + "/" + VshReleaseChannelFile
)
