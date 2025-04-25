# ContactVault Lite

**ContactVault Lite** is the first version of the WhisperVault project — a basic CLI-based contact manager written in Go. It allows you to **Add**, **List**, and **Delete** contacts stored in a local `contacts.json` file.

---

## 🔧 Features

- Add new contacts (Name + Phone)
- View existing contacts
- Delete contact by ID
- JSON-based local storage

---

## 📁 File Structure

    ContactVault-Lite/
    ├── main.go
    └── contacts.json


---

### Build:

    go build

### Run:

    ./ContactVault-Lite

## 📝 Notes

- This version is a **prototype**, built for simplicity and clarity.
- Contact data is stored **unencrypted** in `contacts.json`.
- Each contact is assigned a **unique ID** incrementally.
- No third-party packages are used — only Go standard library.
- All logic is written in a **single `main.go` file** for minimalism.
- Error handling is basic — future versions will improve this.

---