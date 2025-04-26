# 🔐 EncryptedVault

A CLI-based encrypted contact manager using AES-256 encryption.

---

## 📚 Overview

- Contacts (name + phone) stored in a secure encrypted vault.
- Encryption using **AES-256-CTR** with password-derived key (SHA-256).
- Data is only decrypted in memory, never saved in plaintext.
- Master password is required to access or modify contacts.

---

## 🗂 Project Structure

Stage-3-EncryptedVault/
├── cmd/
│   └── main.go
├── internal/
│   ├── models/
│   │   └── contact.go
│   ├── storage/
│   │   └── storage.go        # uses crypto to read/write
│   ├── crypto/
│   │   └── crypto.go         # AES-256 encryption/decryption
│   └── vault/
│       └── vault.go
├── go.mod
└── encrypted_contacts.vault  # encrypted JSON blob

---

## 🔑 How it Works

1. **On Startup**:  
   - Asks user for **master password**.

2. **Actions**:
   - **List Contacts**: Decrypts and shows contacts.
   - **Add Contact**: Adds a new contact and re-encrypts.
   - **Delete Contact**: Deletes a contact by ID and re-encrypts.

3. **Security**:
   - Password is **hashed via SHA-256** to produce a 32-byte AES key.
   - Vault file contains only **AES-256 encrypted data**.
   - Decryption happens only during runtime.

---

## 🚀 Run Instructions

    go mod init EncryptedVault
    go run cmd/main.go

---

## Important Notes
   - Same password must be used every time you open the vault.

   - If password is incorrect ➔ Decryption produces random garbage ➔ Contacts cannot be loaded.

   - If encrypted_contacts.vault file is missing, a new one will be created when saving.  

## 📈 Current Level

| Level | Feature                              |
|-------|--------------------------------------|
| 1     | CLI JSON Contact Manager             |
| 2     | Modular Project Structure            |
| 3     | AES-256 Encrypted Vault with Password|

---