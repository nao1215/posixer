[![Build](https://github.com/nao1215/posixer/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/posixer/actions/workflows/build.yml)
[![UnitTest](https://github.com/nao1215/posixer/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/posixer/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/nao1215/posixer/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/posixer/actions/workflows/reviewdog.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nao1215/posixer)](https://goreportcard.com/report/github.com/nao1215/posixer)
![GitHub](https://img.shields.io/github/license/nao1215/posixer)  
# posixer - Provides information about POSIX commands
**posixer** checks if the POSIX command is installed in your system.

# How to install
### Step.1 Install golang
posixer command only supports installation with `$ go install`. If you does not have the golang development environment installed on your system, please install golang from the [golang official website](https://go.dev/doc/install).

### Step2. Install posixer
```
$ go install github.com/nao1215/posixer@latest
```
# How to use
### Checks if the POSIX commands is already installed
```
$ posixer check
+------------+----------------+----------------+---------------------+
|    NAME    |      TYPE      | IN YOUR SYSTEM |        PATH         |
+------------+----------------+----------------+---------------------+
| alias      | shell built-in | installed      | /usr/bin/alias      |
| bg         | shell built-in | installed      | /usr/bin/bg         |
| cd         | shell built-in | installed      | /usr/bin/cd         |
| command    | shell built-in | installed      | /usr/bin/command    |
   :
   :
| wait       | shell built-in | installed      | /usr/bin/wait       |
| ar         | required       | installed      | /usr/bin/ar         |
   :
   :
| xargs      | required       | installed      | /usr/bin/xargs      |
   :
   :
| val        | optional       | not installed  | -                   |
| vi         | optional       | installed      | /usr/bin/vi         |
| what       | optional       | installed      | /usr/bin/what       |
| yacc       | optional       | installed      | /usr/bin/yacc       |
| zcat       | optional       | installed      | /usr/bin/zcat       |
+------------+----------------+----------------+---------------------+
``` 

### List up POSIX commands
```
$ posixer list
+------------+----------------+---------------------------------------------------------------------------------+
|    NAME    |      TYPE      |                                   DESCRIPTION                                   |
+------------+----------------+---------------------------------------------------------------------------------+
| alias      | shell built-in | define or display aliases                                                       |
| bg         | shell built-in | run jobs in the background                                                      |
| cd         | shell built-in | change the working directory                                                    |
| command    | shell built-in | execute a simple command                                                        |
| echo       | shell built-in | write arguments to standard output                                              |
| false      | shell built-in | return false value                                                              |
| fc         | shell built-in | process the command history list                                                |
| fg         | shell built-in | run jobs in the foreground                                                      |
| getopts    | shell built-in | parse utility options                                                           |
   :
   :
| unget      | optional       | undo a previous get of an SCCS file                                             |
| unlink     | optional       | call the unlink function                                                        |
| uucp       | optional       | system-to-system copy                                                           |
| uustat     | optional       | uucp status enquiry and job control                                             |
| uux        | optional       | remote command execution                                                        |
| val        | optional       | validate SCCS files                                                             |
| vi         | optional       | screen-oriented (visual) display editor                                         |
| what       | optional       | identify SCCS files                                                             |
| yacc       | optional       | yet another compiler compiler                                                   |
| zcat       | optional       | expand and concatenate data                                                     |
+------------+----------------+---------------------------------------------------------------------------------+
```

# Contributing
First off, thanks for taking the time to contribute! ❤️
See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.  

# Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/posixer/issues)

# LICENSE
The posixer project is licensed under the terms of [MIT LICENSE](./LICENSE).
