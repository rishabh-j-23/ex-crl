# ex-crl

The `ex-crl` project is a command-line tool designed to perform HTTP requests and print the responses in a readable format. Requests are stored in a JSON file, and the tool executes them sequentially in the order defined in the file.

## Prerequisites

Before using `ex-crl`, ensure you have the following installed:

- Go 1.23 or later
- Make
- fzf (a command-line fuzzy finder)
- Neovim 0.11.2 or later

## Installation

To install `ex-crl`, run the following commands:

```bash
git clone https://github.com/rishabh-j-23/ex-crl.git
cd ex-crl
make build
make install
```

## TODO

- [ ] Add support for graphQL requests
- [ ] Add support for windows 

## Usage

### 1. Initialize a New Project

Navigate to your project directory and initialize a new `ex-crl` project:

```bash
cd 
ex-crl init
```

This creates the project structure needed to manage requests and workflows.

### 2. Add Requests

Add HTTP requests to your project with:

```bash
ex-crl add request   
```

- ``: HTTP method such as `GET`, `POST`, `PUT`, `DELETE`, or `PATCH`.
- ``: A unique identifier for the request.
- ``: The URL to which the request will be sent.

### 3. Edit Project Configuration

Edit the project configuration file to set parameters like the active environment, project name, and requests directory:

```bash
ex-crl project --edit
```

This opens the configuration file in your default editor.

### 4. Edit Workflow Configuration

Modify the workflow configuration to specify the order in which requests are executed:

```bash
ex-crl workflow --edit
```

The workflow file schema looks like this:

```json
{
    "workflow": [
        {
            "request-name": "sample-request-name",
            "exec": false
        },
        {
            "request-name": "users_login-rishabh",
            "exec": false
        }
    ]
}
```

- The order of requests in the `"workflow"` array determines their execution order.
- The `"exec"` field controls whether a request is executed (`true`) or skipped (`false`).

### 5. Execute the Workflow

Run the workflow to execute the requests in the specified order:

```bash
ex-crl exec
```

### 6. Change the Active Environment

Switch the active environment in the project configuration:

```bash
ex-crl project --env 
```

This changes the environment used for request execution.

This tool streamlines managing and executing HTTP requests in a project-oriented and environment-aware manner, making it easier to automate and test APIs.

## Logging

`ex-crl` uses structured logging for all operations. By default, logs are written to `/var/log/ex-crl/ex-crl.log` (requires root permissions to create the directory on first use). Logs are rotated daily, compressed after a day, and deleted after a week.

If `/var/log/ex-crl` is not writable, logs will be written to `$HOME/ex-crl/logs/ex-crl.log` instead. You can check logs for troubleshooting and auditing.

## Shell Autocompletion

`ex-crl` supports shell autocompletion for Bash, Zsh, Fish, and PowerShell. To enable autocompletion, run:

```bash
ex-crl completion bash   # for Bash
ex-crl completion zsh    # for Zsh
ex-crl completion fish   # for Fish
ex-crl completion powershell # for PowerShell
```

Follow the output instructions to add completion to your shell profile.

## Running Tests

To run all tests:

```bash
go test ./...
```

Ensure you have all dependencies installed. Tests are located in the `internal/` directory and cover core logic and request construction.

## Contributing

Contributions are welcome! To contribute:

- Fork the repository and create a new branch for your feature or bugfix.
- Write clear, concise commit messages and add tests for new features.
- Run `go test ./...` and ensure all tests pass before submitting a pull request.
- If you find a bug or have a feature request, please open an issue or submit a pull request on the GitHub repository.

## License

`ex-crl` is released under the MIT License. See the [LICENSE](LICENSE) file for details.
