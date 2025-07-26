# ex-crl

`ex-crl` is a flexible, scriptable command-line tool for managing and executing HTTP requests in a project-oriented, environment-aware way. It helps automate, test, and organize API workflows with ease.

---

## Minimum Requirements

- **Go**: 1.23 or later
- **Make**: for build/install convenience
- **A text editor**: The `$EDITOR` environment variable should be set (e.g., `vim`, `nvim`, `nano`, `code`, etc.)
- **fzf**: required for fuzzy selection in interactive commands
- **Neovim**: required as the default editor fallback

> **Note:**
> - The `$EDITOR` environment variable is used to open configuration and request files for editing. If not set, `ex-crl` will use `nvim` (Neovim) as the default fallback editor. Please ensure Neovim is installed and accessible in your PATH for the best experience.

---

## Features

### 🚀 Project Initialization
**Command:** `ex-crl init [project-name] [environment] [base-api-url]`  
Quickly bootstrap a new project directory with all necessary configuration files and folders.  
- Sets up project structure for requests, environments, and workflows.
- Lets you specify a project name, environment, and base API URL.

---

### 📦 Request Management
**Command:** `ex-crl add request`  
Add, edit, and organize HTTP requests for your project.
- Supports all HTTP methods (`GET`, `POST`, `PUT`, `DELETE`, `PATCH`).
- Each request is stored as a JSON file with a unique name and URL.
- Supports custom headers, cookies, and environment variables.

---

### 🔄 Workflow Automation
**Command:** `ex-crl workflow --edit`  
Define and manage workflows to automate sequences of requests.
- Specify the order of requests in a workflow JSON file.
- Control which requests are executed or skipped using the `exec` flag.
- Easily edit workflows to match your API testing or automation needs.

---

### ⚙️ Project Configuration
**Command:** `ex-crl project --edit`  
Manage project-level settings and environment configurations.
- Edit the active environment, project name, and requests directory.
- Switch environments with `ex-crl project --env`.

---

### 📝 Interactive CLI
- User-friendly command-line interface with clear help and command completion.
- All commands support `--help` for detailed usage instructions.

---

### 🧪 Testing and Validation
- Built-in test suite for robust request and workflow validation.
- Run all tests with `go test ./...` to ensure reliability.

---

### 🛠️ Utilities and Logging
- Utilities for file management, path handling, and more.
- Structured logging for all operations, with automatic log rotation and fallback to user directory if system logs are unavailable.

---

### 🖥️ Shell Autocompletion
**Command:**  
```bash
ex-crl completion bash        # for Bash
ex-crl completion zsh         # for Zsh
ex-crl completion fish        # for Fish
ex-crl completion powershell  # for PowerShell
```
- Enables tab-completion for commands and flags in your shell.
- Follow the output instructions to add completion to your shell profile.

---

## Installation

```bash
git clone https://github.com/rishabh-j-23/ex-crl.git
cd ex-crl
make build
make install
```

---

## Usage

### Example: From Project Init to Workflow Execution

1. **Initialize a New Project**

   ```bash
   ex-crl init myproject dev https://api.example.com
   ```
   Creates a new project named `myproject` with a `dev` environment and sets the base API URL.

2. **Add HTTP Requests**

   ```bash
   ex-crl add request
   ```
   Follow the prompts to add a new request (e.g., `get-users`). Repeat to add more requests (e.g., `create-user`).

3. **Edit the Workflow**

   ```bash
   ex-crl workflow --edit
   ```
   This opens the workflow configuration in your editor (default: `nvim` or your `$EDITOR`).
   Example workflow file:
   ```json
   {
       "workflow": [
           { "request-name": "get-users", "exec": true },
           { "request-name": "create-user", "exec": true }
       ]
   }
   ```
   Arrange the requests in the order you want them executed and set `"exec": true` for each.

4. **Execute the Workflow**

   ```bash
   ex-crl exec
   ```
   Runs the workflow, executing each request in order. Responses are printed to the terminal.

5. **(Optional) Switch Environments**

   ```bash
   ex-crl project --env
   ```
   Change the active environment if you have multiple (e.g., `dev`, `staging`, `prod`).

---

### 1. Initialize a New Project

```bash
ex-crl init myproject dev https://api.example.com
```
Creates the project structure needed to manage requests and workflows.

---

### 2. Add Requests

```bash
ex-crl add request
```
Guides you through adding a new HTTP request to your project.

---

### 3. Edit Project Configuration

```bash
ex-crl project --edit
```
Opens the configuration file in your default editor for customization.

---

### 4. Edit Workflow Configuration

```bash
ex-crl workflow --edit
```
Modify the workflow file to specify the order and execution of requests.

---

### 5. Execute the Workflow

```bash
ex-crl exec
```
Runs the workflow, executing requests in the specified order.

---

### 6. Change the Active Environment

```bash
ex-crl project --env
```
Switches the environment used for request execution.

---

## Logging

- Logs are written to `/var/log/ex-crl/ex-crl.log` (system-wide) or `$HOME/ex-crl/logs/ex-crl.log` (user fallback).
- Logs are rotated daily, compressed after a day, and deleted after a week.

---

## Running Tests

```bash
go test ./...
```
Tests are located in the `internal/` directory and cover core logic and request construction.

---

## Contributing

Contributions are welcome!  
If you find a bug or want to suggest a new feature, please open an issue or pull request.

---

## License

`ex-crl` is released under the MIT License. See the [LICENSE](LICENSE) file for details.
