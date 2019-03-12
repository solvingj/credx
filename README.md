#### Build Status
| windows | macos | linux | 
|-----------|---------|-------|
|[![Build Status](https://jerrywiltse.visualstudio.com/credx/_apis/build/status/credx?branchName=master&jobName=windows_x64)](https://jerrywiltse.visualstudio.com/credx/_build/latest?definitionId=3&branchName=master)|[![Build Status](https://jerrywiltse.visualstudio.com/credx/_apis/build/status/credx?branchName=master&jobName=macos_x64)](https://jerrywiltse.visualstudio.com/credx/_build/latest?definitionId=3&branchName=master)|[![Build Status](https://jerrywiltse.visualstudio.com/credx/_apis/build/status/credx?branchName=master&jobName=linux_x64)](https://jerrywiltse.visualstudio.com/credx/_build/latest?definitionId=3&branchName=master)

#### Latest Version  
[ ![Download](https://api.bintray.com/packages/solvingj/public-bin/credx/images/download.svg) ](https://bintray.com/solvingj/public-bin/credx/_latestVersion)

#### Project Status  
Pre-pre-alpha: Several core functions not yet implemented

# credx  
`credx` aims to provide a native, cross-platform, cross-language command-line utility for managing credentials which are used during development workflows.  It intends to provide a more secure solution to storing, retrieval, and sharing of credentials on developers' local machines.  `credx` was partially inspired by `pass` (passwordstore.org), but makes the fundamental addition of leveraging the operating system keychain as the core credential store. Also, it does supports `pass` as a storage backend, and adds some conviences to using it. `credx` also aims to be part of a larger toolset and devops tool strategy. 

## How to Use  
The core `credx` api is very simple:

		NAME:
		   credx cred - work with credentials

		USAGE:
		   credx cred command [command options] [arguments...]

		COMMANDS:
			 list    list all credentials
			 set     create a credential or update the value of an existing credential
			 get     retrieve the value of a credential
			 remove  remove a credential (delete)

There is also one global flag which is critical for all operations, and that is the `--backend` flag. `credx` supports the following backends for storing credentials: 

		* wincred (windows)
		* secret-service/libsecret (linux)
		* kwallet (linux)
		* keychain (macos)
		* file (all os) 
		* pass (all os)

By default, `credx` will try to detect and use the native operating-system backend, but if your operating system does not have one, you can choose `file` or `pass`.  Of note, there is a long list of additional flags corresponding to each of the backends. 
 
## Advantage  
Currently, countless integration test suites and devops scripts do something like this to obtain and use credentials: 

		`os.getenv("SOME_PASSWORD")` (python)
		
This implies that the credential is stored as an environment variable in plain-text. This means the developer (or team of developers) read/copied them as plain-text from some shared location, and then had to set the credential as an environment variable on their system (or in a script) manually. 

Using `credx` can provides some level of improvement to the security of this strategy. However, using environment variables is still much more convient.  For most developers, the security advantages of using `credx` are likely to be outweighed by the disadvantages. The roadmap includes adding a secure mechanism for developers to copy credentials to multiple machines they manage, and also to share credentials among team members, all via the CLI. It also includes providing a Json API which can be wrapped using libraries from many different development ecosystems, providing language-native API's.  Use of `credx` is not recommended until some of these features have been added. 

Finally, `credx` has the potential to be leveraged by a number of different ecosystems which have `credential-helpers` (`git`, `docker`, etc).  We will be exploring the needs of those ecosystems to see if it can be a suitable provider. 
 
## Intended Improvements  
Here are the conveniences that `credx` aims to provide: 

	* Provide a cross-platform, native, command-line API to various local credential stores
	* Auto-generate private encryption key, stored in OS keychain
	* Encrypt credentials automatically upon storage
	* Decrypt crecentials automatically upon retrieval
	* Provide json API suitable for various ecosystem tools to integrate with
	* Provide 2+ storage strategies for credentials: 
		* Store credentials directly in OS keychain
		* Store credentials in text files (passwordstore.org)
		* Using other storage tools with API's as backends

## Initial Goals  
- Provide command-line API's for storing and retrieving named credentials
- Provide keychain storage support for all three major operating systems 
- Provide automatic encryption/decryption

## Future Goals  
- Provide credential file support via PGP using passwordstore.org concepts
- Provide support for other credential managers as backends.
- Provide possible integrations with other auth services, such as 2-factor auth
- Provide mechanisms for credential access on servers (for runtime workflows)
- Provide mechansims for sharing team credentials via GIT
- Explore aws-vault project for strategies surrounding temporary passwords and sharing

## Download Instructions  
Precompiled binaries for Windows, Linux, and macOS are hosted on Bintray.com (courtesy of JFrog).  Eventually, we may package them and submit to the various package managers. 

Windows:    
    POSH: `curl -OutFile credx.exe https://dl.bintray.com/solvingj/public-bin/windows_x64/credx.exe`    
    CMD: `powershell -command "curl -OutFile credx.exe https://dl.bintray.com/solvingj/public-bin/windows_x64/credx.exe"`  
	
macOS:   
    `curl -L "https://dl.bintray.com/solvingj/public-bin/macos_x64/credx" -o credx`

Linux:    
    `curl -L "https://dl.bintray.com/solvingj/public-bin/linux_x64/credx" -o credx`
    

## Build Instructions

If you want to contribute to the code, all you need is a recent version of Go (1.10.0+).  With that, you can just run these commands in the root of the repository: 

## Windows: 
    go build -o credx.exe
    
## Linux/macOS
    go build -o credx

To run unit tests, use the following standard command: 

    go test ./...
    
    
## Domain Background
Many development workflows which involve credentials.  This includes integration tests against web services and databases, publishing of artifacts to hosting services, automation testing, and countless other scenarios. Developers generally have a need to manage a number of personal credentials, as well as interact with a number of shared credentials for any number of teams they may be a part of.  Many tools and workflow conventions support reading credentials from environment variables and/or stored in designated text files inside users profiles.  Sadly, many of these still fundamentally require the credentials to be stored in these locations in plain-text. 

There are currently a number of exciting cloud-based credential management tools entering the development ecosystem, including Hashicorp's Vault, as well as "Key Management" facilities provided by most of the PAAS cloud vendors.  There is also keybase.io which takes a very unique identity-oriented approach to authentication in a very general way.  Finally, there are authentication services such as Duo-Security which can add 2-factor auth to a wide variety of workflows. 

Authentication in development workflows is anything but simple.  There is no single tool or strategy which can address all the authenticatin challenges in non-trivial workflows, nor any single collection of tools which can service all workflows for all teams.  One major gap observed across a variety of teams and development ecosystems is the storage, retrieval, and sharing of credentials on developer's local machines. 

## Motivations
First, storing a credential in plain-text feels really bad, yet countless tools in countless ecosystems encourage developers to do this with text files or environment variables.  Furthermore, due to the lack of a generalized and ergonomic solution for development workflows, there is an extensive history of developers storing such plain-text credentials in SCM systems.  The best-practices around storing credentials in SCM systems such as Git are clear and loud:  it is constantly advised against, and has been for a very long time.  Nonetheless, it is a best-practice that is constantly violated on both small and large scales.  In recent years, there have been numerous high-profile data leaks resulting from credentials being stored in Git repositories.  

While storing credentials in Git is a widely-recognized anti-pattern, there is very little guidance on what developers should do instead.  The tools which do exist fail to properly accomodate some of the most common workflows encountered with modern software development, particularly in the OSS world, but they'are also present in private codebases.  Here are the two preliminary workflows which `credx` aims to improve upon: 

	* Store/Retrieve credentials used in integration tests
	* Store/Retrieve credentials used for publishing (web content, artifacts, etc)
	* Store/Retrieve credentials used for infrastructure automation
	* ... all on the local developer machine
	
## Special Thanks  
`credx` was really only made the [keyring](https://github.com/99designs/keyring) library created by brilliant team at 99designs.  `keyring` (and it's dependencies) handles at least 90% of the complexity of working with these credentials, and exposed a uniform Go API (`get`/`set`/`list`/`remove`) over all the various backends.  Thus, `credx` had a very easy time exposing a command-line interface to these Go API's.  `credx` would not have been created if not for `keyring`, it simply would have taken too much time. 