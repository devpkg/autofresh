# Autofresh

        ___         __        ______               __
       /   | __  __/ /_____  / ____/_______  _____/ /_
      / /| |/ / / / __/ __ \/ /_  / ___/ _ \/ ___/ __ \
     / ___ / /_/ / /_/ /_/ / __/ / /  /  __(__  ) / / /
    /_/  |_\__,_/\__/\____/_/   /_/   \___/____/_/ /_/


Autofresh is a simple live-reload development server that rebuilds your program
every time a file is saved, added, or deleted. Gone are the days where you need
to manually recompile a program to test it. Install Autofresh, run it in your
project directory, and start coding away.

### Installation

You can checkout the releases page and download the latest version of autofresh.
Install the binary compiled for your operating system.

If you have a Go runtime installed, you can `go get` this repository and
automatically install it. Make sure you put the GOBIN in your PATH.

```bash 
go get github.com/TerrenceHo/autofresh 
```

### Dependencies

##### Watchman

Watchman is hard dependency, because autofresh uses watchman to recursively
watch your project directory, rather use a custom file watcher.  Install
watchman [here](https://facebook.github.io/watchman/docs/install.html).

### Configuration

Autofresh takes in three ways of configuration, flags, environment variables, a
configuration file autofresh-config.{json,yml/yaml,toml}, and finally the
defaults, with the each item taking precedence in that order (i.e. if both flags
and environment variables set watchman's path, the flag's path will be taken).
The easiest way to configure Autofresh is to use a autofresh-config file. A
sample autofresh-config file is shown below.

```json 
autofresh-config.json here 
```

### Live reload vs hot reload

While partially inspired by React Native's hot reloading development server,
Autofresh does not hot reload, only live reload. It is intended to work for most
languages/scripts, and simply rebuilds and restarts the program.

However, Autofresh works with all languages, as long as you provide a build
command and a run command.

