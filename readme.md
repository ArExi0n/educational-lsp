# educationalsp

> **A Language Server built to teach you how Language Server Protocol actually works.**

It is designed to answer a different question:

> **What is my editor actually doing when I use an LSP?**

Instead of hiding LSP behind a framework, `educationalsp` keeps the implementation explicit and readable so you can follow the communication between an editor and a language server.

It was tested with **Neovim** and is primarily intended as a learning project.

---

## Why?

If you've used Neovim, VS Code, Helix, Zed, or another modern editor, you've probably interacted with language servers without thinking about what happens underneath.

You write:

```text
fn main() {
    println!("hello");
}
```

Your editor somehow knows:

- where the errors are
- what symbols exist
- what a function means
- what completions to suggest
- what documentation to display
- where a definition is located

But how?

The answer is often **LSP**.

`educationalsp` exists to make that process understandable.

---

## What is LSP?

**Language Server Protocol (LSP)** is a protocol that allows an editor and a language server to communicate using a standardized interface.

Conceptually:

```text
┌──────────────┐
│    Editor    │
│   Neovim     │
└──────┬───────┘
       │
       │ JSON-RPC / LSP
       │
       ▼
┌──────────────┐
│ Language     │
│ Server       │
│ educationalsp│
└──────┬───────┘
       │
       ▼
   Language
   Intelligence
```

The editor is the **LSP client**.

`educationalsp` is the **LSP server**.

They communicate using **JSON-RPC messages** over standard input/output.

---

## The Goal

The goal of this project is not to build the best language server.

The goal is to make the following concepts easy to understand:

```text
Editor
  ↓
LSP Client
  ↓
JSON-RPC
  ↓
LSP Server
  ↓
Request Handler
  ↓
Response / Notification
  ↓
Editor
```

Every layer should be visible.

---

## Features

The server intentionally implements only a small set of functionality.

Depending on the current implementation, these may include:

- [x] LSP initialization
- [x] JSON-RPC communication
- [x] `initialize`
- [x] `initialized`
- [x] document synchronization
- [x] basic hover support
- [x] basic diagnostics
- [x] basic completion
- [x] document symbols
- [x] verbose protocol logging
- [x] Neovim compatibility

More features may be added, but complexity is deliberately kept low.

> **Educationalsp values readability over completeness.**

---

# See LSP in Action

Run the server with tracing enabled:

```bash
educationalsp --trace
```

Then open a file in Neovim.

You can observe messages similar to:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "rootUri": "file:///project"
  }
}
```

The server responds:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "capabilities": {}
  }
}
```

This is the important part of the project.

You can **see what your editor is actually sending to the language server**.

---

# Architecture

The project intentionally keeps the architecture straightforward:

```text
                    ┌─────────────┐
                    │   Neovim    │
                    │ LSP Client  │
                    └──────┬──────┘
                           │
                           │ stdin/stdout
                           │
                     JSON-RPC / LSP
                           │
                           ▼
                 ┌───────────────────┐
                 │     Transport     │
                 └─────────┬─────────┘
                           │
                           ▼
                 ┌───────────────────┐
                 │   JSON-RPC Layer  │
                 └─────────┬─────────┘
                           │
                           ▼
                 ┌───────────────────┐
                 │  Request Router   │
                 └─────────┬─────────┘
                           │
             ┌─────────────┼─────────────┐
             ▼             ▼             ▼
        Initialize       Hover      Diagnostics
             │             │             │
             └─────────────┼─────────────┘
                           ▼
                    LSP Response
```

The intention is that you can start at `main.go` and follow a request all the way through the server.

---

# Project Structure

```text
educationalsp/
├── cmd/
│   └── educationalsp/
│       └── main.go
│
├── internal/
│   ├── jsonrpc/
│   │   ├── message.go
│   │   ├── reader.go
│   │   └── writer.go
│   │
│   ├── lsp/
│   │   ├── server.go
│   │   ├── initialize.go
│   │   ├── hover.go
│   │   ├── completion.go
│   │   └── diagnostics.go
│   │
│   └── document/
│       └── document.go
│
├── docs/
│   ├── architecture.md
│   ├── jsonrpc.md
│   └── lsp.md
│
├── go.mod
├── go.sum
└── README.md
```

The exact structure may change as the project evolves.

---

# Running

## Requirements

- Go 1.XX+
- Neovim
- A basic understanding of JSON is helpful

---

## Build

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/educationalsp.git
cd educationalsp
```

Build it:

```bash
go build ./...
```

Or install it:

```bash
go install ./cmd/educationalsp
```

---

# Using with Neovim

Add the server to your Neovim LSP configuration.

For example:

```lua
vim.lsp.config("educationalsp", {
    cmd = { "educationalsp" },
    filetypes = { "educational" },
    root_markers = { ".git" },
})

vim.lsp.enable("educationalsp")
```

The exact configuration may change depending on your Neovim version and project setup.

---

# Understanding the Protocol

One of the best ways to learn LSP is to watch the messages.

For example, when Neovim starts the server, it sends:

```text
initialize
```

The server responds with its capabilities.

Then Neovim sends:

```text
initialized
```

When a document is opened:

```text
textDocument/didOpen
```

When the document changes:

```text
textDocument/didChange
```

When you request hover information:

```text
textDocument/hover
```

When you request completion:

```text
textDocument/completion
```

The important realization is:

> **LSP features are ultimately protocol messages.**

The editor doesn't magically "know" what a language server is doing.

It communicates with it.

---

# JSON-RPC

LSP uses **JSON-RPC** as its communication mechanism.

A simplified request looks like:

```json
{
  "jsonrpc": "2.0",
  "id": 42,
  "method": "textDocument/hover",
  "params": {
    "textDocument": {
      "uri": "file:///example.txt"
    },
    "position": {
      "line": 4,
      "character": 10
    }
  }
}
```

The server processes the request and returns:

```json
{
  "jsonrpc": "2.0",
  "id": 42,
  "result": {
    "contents": {
      "kind": "markdown",
      "value": "Hello from educationalsp!"
    }
  }
}
```

That's the core idea.

Everything else builds on top of this communication model.

---

# What You Should Learn From This Project

After reading through `educationalsp`, you should have a much clearer understanding of:

### 1. What an LSP server is

A process that communicates with an LSP client using the LSP protocol.

### 2. What an LSP client is

Usually your editor.

For example:

```text
Neovim
   │
   ▼
LSP Client
   │
   ▼
educationalsp
```

### 3. How processes communicate

The server normally communicates through:

```text
stdin  → requests from editor
stdout → responses / notifications
```

### 4. What JSON-RPC does

JSON-RPC provides the request/response mechanism underneath LSP.

### 5. What LSP actually defines

LSP defines standardized messages and capabilities for things like:

```text
completion
hover
definition
references
diagnostics
symbols
formatting
rename
code actions
```

### 6. How your editor knows what the server supports

During initialization, the server advertises its capabilities.

---

# Design Philosophy

`educationalsp` follows a few rules.

### Small over clever

Code should be understandable before it is optimized.

### Explicit over abstract

If an abstraction makes the protocol harder to understand, don't use it.

### Protocol over magic

The actual LSP messages should be easy to inspect.

### Learning over features

A smaller implementation that teaches you something is more valuable than a huge implementation that hides everything.

---

# Why Go?

Go is a particularly good fit for this project because it provides:

- simple concurrency primitives
- straightforward I/O
- excellent JSON support
- fast compilation
- a small runtime
- relatively little language complexity

Most importantly, Go lets the implementation stay close to the concepts being taught.

The project is intentionally **not** trying to demonstrate every advanced feature of Go.

---

# Who Is This For?

`educationalsp` is useful if you:

- use Neovim and want to understand LSP
- are learning how language servers work
- want to implement an LSP server from scratch
- want to understand JSON-RPC
- are interested in editor tooling
- want to learn how modern developer tools communicate
- want a small Go networking/protocol project to study

You don't need to understand compilers to start.

You don't need to understand LSP beforehand.

You can simply follow the messages.

---

# Learning Path

A good way to explore the project is:

```text
1. main.go
      ↓
2. stdin/stdout transport
      ↓
3. JSON-RPC
      ↓
4. initialize
      ↓
5. document synchronization
      ↓
6. request routing
      ↓
7. LSP features
```

Start with `initialize`.

Then follow a single request.

Don't try to understand the entire codebase at once.

---

# What This Project Is NOT

`educationalsp` is not intended to replace:

- rust-analyzer
- gopls
- clangd
- pyright
- typescript-language-server

It is also not intended to be a production-grade language server.

Those projects solve a different problem.

`educationalsp` exists to make the underlying machinery easier to understand.

---

# Roadmap

- [ ] Complete JSON-RPC implementation
- [ ] Better protocol tracing
- [ ] Interactive request inspector
- [ ] Document synchronization examples
- [ ] Hover example
- [ ] Completion example
- [ ] Diagnostics example
- [ ] Go-to-definition example
- [ ] Document symbols
- [ ] Rename example
- [ ] Code actions
- [ ] Detailed LSP walkthroughs
- [ ] Neovim demo configuration
- [ ] Tests for every protocol interaction
- [ ] "Build an LSP from scratch" tutorial

---

# Contributing

Contributions are welcome.

However, this project has a slightly unusual priority:

> **A contribution that makes the code easier to understand is often more valuable than one that adds a feature.**

If you want to contribute, consider:

- improving documentation
- simplifying an abstraction
- adding protocol examples
- improving logging
- adding tests
- explaining an LSP concept
- adding a small, isolated LSP feature

Before adding a large dependency, ask whether the underlying concept would become harder to understand.

---

# License

MIT License.

See [`LICENSE`](LICENSE) for details.

---

# The One-Sentence Version

**`educationalsp` is an intentionally simple Go implementation of LSP designed to let you see and understand what happens between your editor and a language server.**

---

## Start Here

If you're completely new to LSP, don't start by reading the entire specification.

Start the server.

Open Neovim.

Enable tracing.

Make a change to a file.

Watch the messages.

Then ask:

> **"Why did Neovim send this?"**

That question is what this project is built to answer.
