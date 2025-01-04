# Codecrafters HTTP Server in Go

A simple HTTP server built as part of the Codecrafters challenge. This project demonstrates HTTP request handling, routing, file uploads/downloads, and serving dynamic content in Go.

---

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)

---

## Features

- **Routing**:
   - Support for dynamic routes with parameters (e.g., `/echo/:value`).
- **File Handling**:
   - Upload files via `POST`.
   - Download files via `GET`.
- **Custom Handlers**:
   - Retrieve `User-Agent` headers.
   - Echo route parameters.
- **Flexible Directory Management**:
   - Specify a directory for file operations using a CLI flag.
- **Error Handling**:
   - Return proper status codes for missing files, invalid paths, and other errors.

---

## Project Structure

```plaintext
codecrafters-http-server-go/
├── app/
│   └── server.go               # Main entry point for the server
├── pkg/
│   └── server/
│       └── http/
│           ├── request.go       # Handles HTTP request parsing and representation
│           ├── response-writer.go # Handles HTTP responses
│           ├── server.go        # HTTP server implementation
│           ├── utils.go         # Utility functions for file handling, etc.
│           └── utils_test.go    # Unit tests for utilities
├── .codecrafters/              # Codecrafters-specific configurations
├── .gitattributes              # Git attributes configuration