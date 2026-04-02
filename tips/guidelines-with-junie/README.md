# Demo Walkthrough

### Go App Project Guidelines using Junie

![demo](./tip.gif)

You can customize project guidelines and tell Junie how to build the application. This includes the coding standards, best practices, and recommendations you want Junie to follow.

To create project guidelines:

- Click the `+` button.
- Select *Create project guidelines*.

### Project Guidelines

This is a placeholder of the project guidelines for Junie.
Replace this text with any project-level instructions for Junie, e.g.:

- What is the project structure

```sh
This is the recommended project structure. For more details, you can also check out: https://github.com/golang-standards/project-layout
myapp/
â”œâ”€â”€ cmd/
â”‚ â””â”€â”€ myapp/ # Main application entry point
â”‚ â””â”€â”€ main.go
â”œâ”€â”€ internal/ # Private application and library code
â”‚ â””â”€â”€ auth/ # Example: internal authentication logic
â”‚ â””â”€â”€ db/ # Example: internal DB access code
â”œâ”€â”€ pkg/ # Public library code (can be used by other projects)
â”‚ â””â”€â”€ utils/ # Example: reusable utility functions
â”œâ”€â”€ api/ # OpenAPI/Swagger specs, Protobuf, gRPC definitions
â”œâ”€â”€ configs/ # Configuration files (JSON, YAML, TOML, etc.)
â”œâ”€â”€ deployments/ # Docker, Kubernetes, Helm, etc.
â”œâ”€â”€ scripts/ # Bash, CLI tools for setup, migration, CI, etc.
â”œâ”€â”€ web/ # Web frontend (React, static files, templates, etc.)
â”œâ”€â”€ test/ # Additional test data and utilities
â”œâ”€â”€ go.mod
â”œâ”€â”€ go.sum
â””â”€â”€ README.md
```

- Whether Junie should run tests to check the correctness of the proposed solution

`Yes, it should include unit tests to verify the functionality end-to-end. Additionally, if the solution involves performance or latency considerations, make sure to include benchmark tests in Go as well.`

- How does Junie run tests (once it requires any non-standard approach)
- Whether Junie should build the project before submitting the result

`Yes, the project should be built, and the build artifacts should be placed in a build folder at the project root. This folder should include builds for Windows, Linux, and macOS platforms.`

- Any code-style related instructions

As an option you can ask Junie to create these guidelines for you.

`Junie, after creating the project, ensure all dependencies are listed in the go.mod file. Then, run go mod download to fetch them. Before starting the application, run go mod tidy to clean up any unused dependencies.`

- `If you are building CRUD application, then use postgres as the default database choice.`
- `Make sure to always include a Dockerfile under deployments.`
