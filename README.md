# jfvm - JFrog CLI Version Manager

[![CI](https://github.com/bhanurp/jfvm/actions/workflows/release.yml/badge.svg)](https://github.com/bhanurp/jfvm/actions/workflows/release.yml)
[![Latest Release](https://img.shields.io/github/v/release/bhanurp/jfvm)](https://github.com/bhanurp/jfvm/releases)
[![License](https://img.shields.io/github/license/bhanurp/jfvm)](https://github.com/bhanurp/jfvm/blob/main/LICENSE)
[![homebrew installs](https://img.shields.io/badge/homebrew-installs-brightgreen?logo=homebrew)](https://github.com/bhanurp/homebrew-jfvm)

**jfvm** is a lightweight CLI tool that helps you manage multiple versions of the [JFrog CLI](https://jfrog.com/getcli/) on your system. It supports auto-installation, version switching, project-specific defaults, local binary linking, and aliasing — all inspired by tools like `nvm`, `sdkman`, and `volta`.

## 🎥 Demo

https://github.com/user-attachments/assets/6984077c-72ab-4f8c-a11c-671e72870efe

https://github.com/user-attachments/assets/32ce3eb1-4f69-49bd-bdc7-9f95cd9ead34


## 🚀 Why jfvm?

Managing different versions of the JFrog CLI across multiple projects and environments can be challenging. `jfvm` simplifies this by:

- Installing any released version of the `jf` binary
- Allowing you to link locally built versions
- Automatically switching versions based on a `.jfrog-version` file
- Letting you define named aliases (e.g., `prod`, `dev`)
- Providing a smooth `jf` shim for command redirection

No more symlink hacking or hardcoded paths.

---

## 🛠️ Installation

### Via Homebrew (with tap):
```bash
brew tap bhanurp/jfvm
brew install jfvm
```

### Via one-liner:
```bash
brew install https://raw.githubusercontent.com/bhanureddy/homebrew-jfvm/main/jfvm.rb
```

### Or Build From Source:
```bash
git clone https://github.com/bhanurp/jfvm.git
cd jfvm
make install
```

---

## 📦 Commands

### `jfvm install <version>`
Installs the specified version of JFrog CLI (`jf`) from JFrog’s public release server.
```bash
jfvm install 2.74.0
```

### `jfvm use <version or alias>`
Activates the given version or alias. If `.jfrog-version` exists in the current directory, that will be used if no argument is passed.
```bash
jfvm use 2.74.0
jfvm use prod
```

### `jfvm list`
Shows all installed versions and the currently active one.
```bash
jfvm list
```

### `jfvm remove <version>`
Removes a specific version of `jf`.
```bash
jfvm remove 2.72.1
```

### `jfvm clear`
Removes **all** installed versions.
```bash
jfvm clear
```

### `jfvm alias set <name> <version>`
Defines an alias for a specific version.
```bash
jfvm alias set dev 2.74.0
```

### `jfvm link --from <path> --name <name>`
Links a **locally built `jf` binary** to be used via `jfvm`.
```bash
jfvm link --from /Users/bhanu/go/bin/jf --name local-dev
jfvm use local-dev
```

---

## 📁 Project-specific Version

Add a `.jfrog-version` file to your repo:
```bash
echo "2.74.0" > .jfrog-version
```
Then run:
```bash
jfvm use
```

---

## ⚙️ Shell Integration
Add this to your shell profile (`.zshrc`, `.bashrc`, etc.):
```bash
export PATH="$HOME/.jfvm/shim:$PATH"
```
This allows the shimmed `jf` command to delegate to the correct version transparently.

---

## 🧪 Example
```bash
$ jfvm install 2.74.0
$ jfvm alias prod 2.74.0
$ jfvm use prod
$ jf --version
jfrog version 2.74.0
```

---

## 🧼 Uninstall
```bash
rm -rf ~/.jfvm
brew uninstall jfvm  # if installed via Homebrew
```

---

## 📬 Feedback / Contributions
PRs and issues welcome! Open source, MIT licensed.

**GitHub:** https://github.com/bhanurp/jfvm
