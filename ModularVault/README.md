# ModularVault

**ModularVault** is version 2 of the WhisperVault project. It builds upon ContactVault Lite by introducing a modular Go project layout. Code is now organized by responsibility (models, storage, vault logic, etc.) for better scalability and maintainability.

---

## 🔧 Features

- Same functionality as ContactVault Lite:
  - Add, List, Delete contacts
- Structured folder layout using Go best practices
- Separation of concerns:
  - CLI entry point
  - Data models
  - Storage functions
  - Business logic (CRUD)

---

## 📁 File Structure

    ModularVault/
    ├── cmd/
    │   └── main.go
    │
    ├── internal/
    │   ├── models/
    │   │   └── contact.go
    │   │
    │   ├── storage/
    │   │   └── storage.go
    │   │
    │   └── vault/
    │       └── vault.go
    │
    └── contacts.json



---

## 🧪 Usage

### Build:

    go build ./cmd

### Run:

    ./cmd

## 📝 Notes

- This version introduces a **modular structure** using Go’s standard project layout practices.
- Logic is split into:
  - `models`: defines data structures (like `Contact`)
  - `storage`: handles file I/O and JSON encoding/decoding
  - `vault`: contains the core business logic (Add/List/Delete)
  - `cmd/main.go`: CLI interface
- The application still uses **plain JSON** with no encryption.
- File paths are relative; `contacts.json` is created in the root dir.
- Prepares the project for adding encryption and user authentication in the next stage.

---