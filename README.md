# ex-crl

The `ex-crl` project is a command line tool that lets you perform a http request and print the response in a readable format. The requests are stored in a json file and the tool will perform the requests in the order they are specified in the file.

## Prerequisites

Before using the `ex-crl` tool, make sure you have the following installed:

- Go 1.23 or later
- Make
- fzf
- neovim 0.11.2 or later

## Installation

To install the `ex-crl` tool, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/rishabh-j-23/ex-crl.git
```

2. Change to the project directory:

```bash
cd ex-crl
```

3. Build the tool:

```bash
make build
```

4. Install the tool:

```bash
make install
```

## Usage

To use the `ex-crl` tool, follow these steps:

1. Create a new project:

```bash
cd <your-project-directory>
ex-crl init
```

This will create a new project directory with the following structure:


2. Add requests to the project:

```bash
ex-crl add request <http-method> <request-name> <endpoint>
```

This will add a new request to the project. The `http-method` can be `GET`, `POST`, `PUT`, `DELETE`, or `PATCH`. The `request-name` is a unique identifier for the request, and the `endpoint` is the URL of the request.

3. Edit the project configuration:

```bash
ex-crl project --edit
```

This will open the project configuration file in your default editor. You can edit the file to specify the active environment, the project name, and the requests directory.

4. Edit the workflow configuration:

```bash
ex-crl workflow --edit
```

This will open the workflow configuration file in your default editor. You can edit the file to specify the order of requests to be executed.

5. Execute the workflow:

```bash
ex-crl exec
```

This will execute the requests in the order specified in the workflow configuration file.

