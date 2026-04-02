# Awesome GoLand

[![Awesome](https://awesome.re/badge.svg)](https://github.com/sindresorhus/awesome#readme)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![GoLand](https://img.shields.io/badge/IDE-GoLand-000000?style=flat&logo=jetbrains)](https://www.jetbrains.com/go/)

> A curated guide to the best tips and tricks for mastering GoLand.

![background](./misc/images/background.png)

## Why Awesome GoLand?

Whether you're new to GoLand or a seasoned user, this repository helps you unlock the IDE's full potential:

- 🐛 **Debug Like a Pro** — Master goroutine labeling, smart step-into, and core dumps
- 🧪 **Automate Testing** — Generate and run tests with minimal effort
- 🔄 **JSON ↔ Struct** — Convert JSON to Go structs in seconds
- 🐳 **Cloud-Native Ready** — Docker and Kubernetes workflows out of the box
- ⚡ **Performance Insights** — CPU, memory, and block profiling made simple
- 🤖 **AI-Powered Coding** — Explore Junie and MCP for smarter development
- 🔍 **Structural Refactoring** — Pattern-based search and replace across your codebase

> 💡 Each module includes hands-on demos — learn by doing, not just reading.


---

## Table of Contents
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Featured Modules](#featured-modules)
  - [Core Development](#core-development)
  - [API & Microservices](#api--microservices)
  - [Docker & Kubernetes](#docker--kubernetes)
  - [Profiling & Performance](#profiling--performance)
  - [AI & Extensibility](#ai--extensibility)
- [Tips & Tricks](#tips--tricks)
- [How to Use](#how-to-use)
- [Contributing](#contributing)


---

## Prerequisites

Before you begin, ensure you have the following installed:
- [Go 1.22+](https://go.dev/dl/)
- [GoLand IDE](https://www.jetbrains.com/go/download/) (Latest version recommended)
- [Docker](https://www.docker.com/get-started) (Required for Docker and Kubernetes modules)
- [Kubernetes / kubectl](https://kubernetes.io/docs/tasks/tools/) (Required for Kubernetes modules)

## Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/mukulmantosh/awesome_goland.git
   cd awesome_goland
   ```
2. **Open in GoLand:**
   - Launch GoLand.
   - Click `Open` and select the `awesome_goland` directory.
3. **Download dependencies:**
   - GoLand will automatically detect `go.mod` files and prompt you to download dependencies.
   - Alternatively, run: `go mod download` in the root (though many modules have their own `go.mod`).

## Featured Modules

Each module is designed to demonstrate specific GoLand capabilities. Look for `DEMO.md` or `README.md` in each directory for detailed walkthroughs.

### Core Development

- **[Debugging](./features/debugging)**: Master the debugger with [Smart Step Into](./features/debugging/smart_step), [Goroutine Labeling](./features/debugging/label_goroutines), and [Core Dumps](./features/debugging/dump_goroutines).

- **[Testing](./features/testing/automate)**: Streamline your testing workflow with automated test generation and execution. [View Demo](./features/testing/automate/DEMO.md).
- **[JSON to Struct](./features/json-to-struct)**: Effortlessly convert JSON data into Go structs.
- **[Resource Leak Detection](./features/resource_leak_detection)**: Detect and fix resource leaks early in the development cycle. [View Demo](./features/resource_leak_detection/DEMO.md).
- **[Structural Search & Replace](./features/structural_search)**: Powerfully refactor code using pattern-based search. [Watch Video](./features/structural_search/structural_search.mp4).

### API & Microservices
- **[gRPC CRUD](./features/go-grpc-crud)**: A complete gRPC service implementation with Create, Read, Update, and Delete operations. [Read More](./features/go-grpc-crud/README.md).
- **[Endpoint Discovery](./features/smarter_endpoint_discovery)**: Use the Endpoints tool window to discover and test your REST APIs. [View Demo](./features/smarter_endpoint_discovery/DEMO.md).

### Docker & Kubernetes

#### 🐳 Docker
- **[Docker Run Targets](./features/docker/targets/basic_rest_sqlite)**: Learn how to run and debug Go applications inside Docker containers. [View Demo](./features/docker/targets/basic_rest_sqlite/DEMO.md).

#### ☸️ Kubernetes
- **[Kubernetes Integration](./features/kubernetes)**: Orchestrate your Go services with Kubernetes integration. [View Walkthrough](./features/kubernetes/DEMO.md).
- **[Debugging Kubernetes Pods with Telepresence](https://www.youtube.com/watch?v=3khaEf7bmWI)**: Learn how to debug pods directly in the cluster.
- **[Deploying Go Apps with Kubernetes](https://blog.jetbrains.com/go/2024/11/20/deploying-go-apps-with-kubernetes)**: Step-by-step guide from the official JetBrains blog.

### Profiling & Performance
- **[Profilers](./features/profilers)**: Analyze CPU, Memory, and Block profiles to optimize your Go code performance.

### AI & Extensibility
- **[Junie](./features/junie)**: Explore AI-assisted development workflows. [View Demo](./features/junie/DEMO.md).
- **[MCP (Model Context Protocol)](./features/mcp)**: Enhancing LLM interactions within the IDE. [View Demo](./features/mcp/DEMO.md).
- **[Junie Guidelines for Go Apps](https://github.com/JetBrains/junie-guidelines/tree/main/guidelines/go)**: A GitHub repository that provides a collection of technology-specific guidelines to improve Junie code generation.


## Tips & Tricks

Discover quick and powerful ways to boost your productivity in GoLand. Here is the complete list of tips available in this repository:

<details>
<summary><b>Click to expand the full list of 100+ tips</b></summary>

- [Add Field to Struct](./tips/add-field-to-struct)
- [Add Line After Before](./tips/add-line-after-before)
- [Ask Mode Junie](./tips/ask-mode-junie)
- [Assert Completion in Testify](./tips/assert-completion-in-testify)
- [Attach Process](./tips/attach-process)
- [Auto Run Tests](./tips/auto-run-tests)
- [Back and Forth](./tips/back-and-forth)
- [Basic Code Completion](./tips/basic-code-completion)
- [Better Go Version Handling](./tips/better-go-version-handling)
- [Better Time Format Completion](./tips/better-time-format-completion)
- [Better View of Attached Projects](./tips/better-view-of-attached-projects)
- [Boilerplate with Junie](./tips/boilerplate-with-junie)
- [Builtin HTML Preview](./tips/builtin-html-preview)
- [Call Hierarchy](./tips/call-hierarchy)
- [Change Font Size Using Mouse Wheel and From Keyboard](./tips/change-font-size-using-mouse-wheel-and-from-keyboard)
- [Change Struct Tag Style](./tips/change-struct-tag-style)
- [Check APIs From The IDE](./tips/check-apis-from-the-ide)
- [Check Vulnerabilities in Go Mod](./tips/check-vulnerabilities-in-go-mod)
- [Cleanup Unused IDE Directories](./tips/cleanup-unused-ide-directories)
- [Code Completion in Evaluate Expression and Watches](./tips/code-completion-in-evaluate-expression-and-watches)
- [Code Completion in Language Injections](./tips/code-completion-in-language-injections)
- [Code Completion in Run Configurations](./tips/code-completion-in-run-configurations)
- [Code Coverage for Applications](./tips/code-coverage-for-applications)
- [Code Folding Options](./tips/code-folding-options)
- [Code Vision in Editor](./tips/code-vision-in-editor)
- [Compare With Clipboard](./tips/compare-with-clipboard)
- [Complete Current Statement](./tips/complete-current-statement)
- [Completion for Non Imported Go Modules](./tips/completion-for-non-imported-go-modules)
- [Completion With The Tab Key](./tips/completion-with-the-tab-key)
- [Containers Device Cgroup Rules Values Inspection](./tips/containers-device_cgroup_rules-values-inspection)
- [Containers Pause and Unpause](./tips/containers-pause-and-unpause)
- [Containers Port Mapping Inspection](./tips/containers-port-mapping-inspection)
- [Containers Restart a Running Container](./tips/containers-restart-a-running-container)
- [Containers Shm Size Inspection](./tips/containers-shm-size-inspection)
- [Convert Between Different String Types](./tips/convert-between-different-string-types)
- [Convert Empty Interfaces to Any](./tips/convert-empty-interfaces-to-any)
- [Convert Integers to Strings](./tips/convert-integers-to-strings)
- [Create Function in Any Package](./tips/create-function-in-any-package)
- [Create Getters and Setters](./tips/create-getters-and-setters)
- [Create Inspections With Regular Expressions](./tips/create-inspections-with-regular-expressions)
- [Create Missing Struct Type With All Fields](./tips/create-missing-struct-type-with-all-fields)
- [Create Parameter Refactoring](./tips/create-parameter-refactoring)
- [Custom Printf Like Functions](./tips/custom-printf-like-functions)
- [Custom Structure Tags](./tips/custom-structure-tags)
- [Cyclic Expand Word](./tips/cyclic-expand-word)
- [Delete Type Parameter With Empty Parameter List](./tips/delete-type-parameter-with-empty-parameter-list)
- [Disable Tabs](./tips/disable-tabs)
- [Download Kubernetes Pod Log](./tips/download-kubernetes-pod-log)
- [Drag and Drop Editor Tabs](./tips/drag-and-drop-editor-tabs)
- [Dump Goroutines](./tips/dump-goroutines)
- [Evaluate Expression](./tips/evaluate-expression)
- [Exclude From Imports and Completion](./tips/exclude-from-imports-and-completion)
- [Expanding Shrinking Selection](./tips/expanding-shrinking-selection)
- [Extract Type Refactoring](./tips/extract-type-refactoring)
- [Fill Paragraph for Go Comments](./tips/fill-paragraph-for-go-comments)
- [FLCC With GoLand](./tips/flcc-with-goland)
- [Generate a Test for an Element](./tips/generate-a-test-for-an-element)
- [Generate a Test for Generic Function](./tips/generate-a-test-for-generic-function)
- [Generate Go Work](./tips/generate-go-work)
- [Generate Imports While Typing](./tips/generate-imports-while-typing)
- [Get Container and Tag Completion for Docker Files](./tips/get-container-and-tag-completion-for-docker-files)
- [Go 1.17 Convert Slice to Array Pointer](./tips/go-1-17-convert-slice-to-array-pointer)
- [Go 1.16 Retract Directive](./tips/go-116-retract-directive)
- [Go Embed Support](./tips/go-embed-support)
- [Go to From a Test](./tips/go-to-from-a-test)
- [Goroutines Profiler Labels](./tips/goroutines_profiler_labels)
- [Guidelines With Junie](./tips/guidelines-with-junie)
- [Handle Go Errors With Postfix Completion](./tips/handle-go-errors-with-postfix-completion)
- [Hide All Tool Windows](./tips/hide-all-tool-windows)
- [Implement an Interface](./tips/implement-an-interface)
- [Implement Unexported Interfaces](./tips/implement-unexported-interfaces)
- [Inline Watches in Debugger](./tips/inline-watches-in-debugger)
- [Inspect Context Cancel Func Usage](./tips/inspect-context-cancel-func-usage)
- [Inspect Detect Incorrect Usage of Println Like Function](./tips/inspect-detect-incorrect-usage-of-println-like-function)
- [Install and Import](./tips/install-and-import)
- [Integrated Go Playground](./tips/integrated-go-playground)
- [JSON to Go Struct Type](./tips/json-to-go-struct-type)
- [Kubernetes Configure Custom Namespaces Manually](./tips/kubernetes-configure-custom-namespaces-manually)
- [Kubernetes Convert Resources Format](./tips/kubernetes-convert-resources-format)
- [Kubernetes Delete Resource From Run Button](./tips/kubernetes-delete-resource-from-run-button)
- [Live Template for For Loop in Bench Function](./tips/live-template-for-for-loop-in-bench-function)
- [Live Template for Go Bench Function](./tips/live-template-for-go-bench-function)
- [Live Template for Go Test Function](./tips/live-template-for-go-test-function)
- [Load GoMod Changes Manually](./tips/load-gomod-changes-manually)
- [Mermaid JS Support in Markdown](./tips/mermaid-js-support-in-markdown)
- [Method Like Completion for Functions](./tips/method-like-completion-for-functions)
- [Move Block](./tips/move-block)
- [Multiple Selections](./tips/multiple-selections)
- [Navigate Between Opened Files Using the Switcher](./tips/navigate-between-opened-files-using-the-switcher)
- [Navigate Test in Table Test](./tips/navigate-test-in-table-test)
- [Navigate to File](./tips/navigate-to-file)
- [Navigate to Symbol](./tips/navigate-to-symbol)
- [New Scratch File From Selection](./tips/new-scratch-file-from-selection)
- [Open Console in Kubernetes](./tips/open-console-in-kubernetes)
- [Open File in Split Editor](./tips/open-file-in-split-editor)
- [Open Shell in Kubernetes](./tips/open-shell-in-kubernetes)
- [Optimize Imports](./tips/optimize-imports)
- [Parameter Info](./tips/parameter-info)
- [Partial Match Completion](./tips/partial-match-completion)
- [Paste From History](./tips/paste-from-history)
- [Postfix Completion](./tips/postfix-completion)
- [Preview File Contents](./tips/preview-file-contents)
- [Put Project Under VCS](./tips/put-project-under-vcs)
- [Quick Docs](./tips/quick-docs)
- [Quick Documentation](./tips/quick-documentation)
- [Quick Documentation Go Doc Comments](./tips/quick-documentation-go-doc-comments)
- [Quickly Handle Go Error](./tips/quickly-handle-go-error)
- [Recent Locations](./tips/recent-locations)
- [Reduce Clutter](./tips/reduce-clutter)
- [Remove Scratch File When Closing](./tips/remove-scratch-file-when-closing)
- [Rename Constants That Use Names of Builtin Constants](./tips/rename-constants-that-use-names-of-builtin-constants)
- [Rename File](./tips/rename-file)
- [Rename Generic Receivers](./tips/rename-generic-receivers)
- [Rename Go Module Name](./tips/rename-go-module-name)
- [Rename Symbol](./tips/rename-symbol)
- [Rerun Testify Subtest](./tips/rerun-testify-subtest)
- [Run Actions on Save](./tips/run-actions-on-save)
- [Run Anything](./tips/run-anything)
- [Run Gofmt After Builtin Formatter](./tips/run-gofmt-after-builtin-formatter)
- [Run Individual Testify Test Suites](./tips/run-individual-testify-test-suites)
- [Run Targets Docker](./tips/run-targets-docker)
- [Run Tests Fuzzing](./tips/run-tests-fuzzing)
- [Run Tests With Coverage](./tips/run-tests-with-coverage)
- [Run to Cursor](./tips/run-to-cursor)
- [Select All Occurrences in a File](./tips/select-all-occurrences-in-a-file)
- [Select In](./tips/select-in)
- [Show Usages](./tips/show-usages)
- [Simple Math in Search Everywhere](./tips/simple-math-in-search-everywhere)
- [Smart Code Completion](./tips/smart-code-completion)
- [Speed Typing](./tips/speed-typing)
- [Split Screen Without Tabs](./tips/split-screen-without-tabs)
- [Step Over](./tips/step-over)
- [Stop Remote Process](./tips/stop-remote-process)
- [Structure Popup](./tips/structure-popup)
- [Structure Tags](./tips/structure-tags)
- [Switch to the Editor](./tips/switch-to-the-editor)
- [Sync IDE Theme With OS Theme](./tips/sync-ide-theme-with-os-theme)
- [Tailwind CSS Support](./tips/tailwind-css-support)
- [Terminal Cursor Shape](./tips/terminal-cursor-shape)
- [Terminal Select Shell](./tips/terminal-select-shell)
- [Test Bench Fatal Calls](./tips/test-bench-fatal-calls)
- [Test Name Completion in Testify](./tips/test-name-completion-in-testify)
- [Text Search in Local History](./tips/text-search-in-local-history)
- [Type Hierarchy](./tips/type-hierarchy)
- [Unused Dependency in Go Mod Inspection](./tips/unused-dependency-in-go-mod-inspection)
- [Variable Shadowing Display](./tips/variable-shadowing-display)
- [VCS Cleanup Code Before Commit](./tips/vcs-cleanup-code-before-commit)
- [VCS Run Inspections Before Commit](./tips/vcs-run-inspections-before-commit)
- [VCS Run Tests Before Commit](./tips/vcs-run-tests-before-commit)
- [VCS Sign Commit With GPG](./tips/vcs-sign-commit-with-gpg)
- [Vue Support](./tips/vue-support)
- [Web Arrow Function Live Template](./tips/web-arrow-function-live-template)
- [Web Completion for Parameter Types Based on Function Calls](./tips/web-completion-for-parameter-types-based-on-function-calls)
- [Web React Completion for Classnames Clsx Libraries](./tips/web-react-completion-for-classnames-clsx-libraries)
- [Web React Refactor Names in UseState Hooks](./tips/web-react-refactor-names-in-usestate-hooks)
- [Web Support for TypeScript in JSDoc](./tips/web-support-for-typescript-in-jsdoc)
- [Workspaces Group Use Directives](./tips/workspaces-group-use-directives)

</details>

> 💡 **Tip**: Click the summary above to see the full collection of over 140+ GoLand productivity tips!

## How to Use

1. **Explore Modules**: Navigate to any directory in `features/` to see specific examples.
2. **Follow Demos**: Most modules include a `DEMO.md` file with step-by-step instructions.
3. **Run Code**: Use GoLand's gutter icons (run/debug) to execute the `main.go` or `*_test.go` files.


## Blogs

- **[Moving Your Codebase to Go 1.26 With GoLand Syntax Updates](https://blog.jetbrains.com/go/2026/02/13/moving-your-codebase-to-go-1-26-with-goland-syntax-updates/)**
- **[Best Practices for Secure Error Handling in Go](https://blog.jetbrains.com/go/2026/03/02/secure-go-error-handling-best-practices/)**
- **[Write Modern Go Code With Junie and Claude Code](https://blog.jetbrains.com/go/2026/02/20/write-modern-go-code-with-junie-and-claude-code/)**
- **[“GoLand Can Do That?” Ten Secret Superpowers You Might Not Know](https://blog.jetbrains.com/go/2025/09/08/goland-can-do-that-ten-secret-superpowers-you-might-not-know/)**
- **[Preventing Resource Leaks in Go: How GoLand Helps You Write Safer Code](https://blog.jetbrains.com/go/2025/12/09/preventing-resource-leaks-in-go-how-goland-helps-you-write-safer-code/)**
- **[Interprocedural Analysis: Catch nil Dereferences Before They Crash Your Code](https://blog.jetbrains.com/go/2025/07/28/interprocedural-analysis-catch-nil-dereferences-before-they-crash-your-code/)**
- **[Data Flow Analysis for Go](https://blog.jetbrains.com/go/2024/03/26/data-flow-analysis-for-go/)**
- **[Profiling Go Code with GoLand](https://blog.jetbrains.com/go/2023/02/02/profiling-go-code-with-goland/)**
- **[Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)**
- **[Build a Blog With Go Templates](https://blog.jetbrains.com/go/2022/11/08/build-a-blog-with-go-templates/)**


## Repos

- **[Modern Go Guidelines – Best practices that help coding agents write clean, efficient, and idiomatic Go code](https://github.com/JetBrains/go-modern-guidelines)**

## YouTube Videos & Playlist

- **[GoLand Playlist](https://www.youtube.com/playlist?list=PLQ176FUIyIUaXlyMXddbZ7TOdW2Yns4yV)**
- **[Spot the Nil Dereference: How to avoid a billion-dollar mistake](https://www.youtube.com/watch?v=lNnSRakOlvI)**
- **[GopherCon 2025: Nil Today, Outage Tomorrow - Mukul Mantosh](https://www.youtube.com/watch?v=dEnn45M2Los)**
- **[GopherCon 2023: The Blueprints to Building Your Own Badass Community - Benjamin Bryant](https://www.youtube.com/watch?v=sxflWkIvRUU)**

## Courses

- **[Mastering Go with GoLand - byteSizeGo](https://www.bytesizego.com/courses/mastering-go-with-goland)**


## Mastering Shortcuts

- [Reference Sheet](https://resources.jetbrains.com/storage/products/goland/docs/GoLand_ReferenceCard.pdf)


## Contributing

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Community

Join the conversation and connect with other Go developers:

-  [Gophers Slack](https://gophers.slack.com/) — Official Slack community for Go developers. Join the `#goland` channel for GoLand discussions.
-  [Follow GoLand on X](https://x.com/GoLandIDE) — New features: `#NewInGoLand` | Tips and tricks: `#GoLandTip`

Found this helpful? Consider giving it a ⭐ it means a lot!


## Resources

- [Official GoLand Documentation](https://www.jetbrains.com/help/go/)
- [GoLand Blog](https://blog.jetbrains.com/go/)

## License 

MIT


