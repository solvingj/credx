package credx

const AppDescription string = "A cross-platform CLI for OS keychain and other forms of local credential management."

const LogLevelUsage string = `
	verbosity of credx logging
	possible values: 	( INFO | ERROR | DEBUG )
	default value:	[INFO]
	environment variable:	 
	`

const HomeDirectoryUsage string = `
	directory where credx will look for global configuration and disk-based credentials
	default value:		[~/.credx]
	environment variable:
	`

const BackendUsage string = `
	backend to use for credentials, "os" will auto-detect the native operating system's keychain
	possible values: ( secret-service | keychain | kwallet | wincred | os )
	default value:	[os]
	environment variable:	 
	`

const KeyPathUsage string = `
	override the default path to the keychain, valid for disk-based keychain backends
	supported backends: ( file | pass )
	default value:		[CREDX_HOME_DIR/keys/default]
	environment variable:
	`

const KeychainNameUsage string = `
	name of keychain to use, valid for keychain backends which support multiple keychains
	supported backends: ( keychain )
	default value:		["default"]
	environment variable:
	`

const LibsecretCollectionNameUsage string = `
	collection name, only valid when secret-service backend is used
	default value:		["secret-service"]
	environment variable:
	`

const SecretServiceNameUsage string = `
	secret service name, only valid when secret-service backend is used
	default value:		["secret-service"]
	environment variable:
	`

const KWalletServiceNameUsage string = `
	kwallet service name, only valid when kwallet backend is used
	default value:		["kdewallet"]
	environment variable:
	`

const KWalletAppIdUsage string = `
	app id, only valid when kwallet backend is used
	default value:		["envx"]
	environment variable:
	`

const KWalletFolderUsage string = `
	folder credentials will be set/get from, only valid when kwallet backend is used
	default value:		["envx"]
	environment variable:
	`

const PassCmdUsage string = `
	name of the pass executable
	default value:		["pass"]
	environment variable:
	`

const PassDirUsage string = `
	password-store directory
	default value:		["~/.password-store"]
	environment variable:
	`

const PassPrefixUsage string = `
	string prefix to prepend to the item path stored in pass
	default value:		[""]
	environment variable:
	`

