---
name: verify
description: Run repository verification suite (secret scanning with Python, gofmt, go vet, go test, and .env.example parity) before commits or closing tasks.
---

# Verify Skill - Quality Verification Suite

Use this skill to ensure code meets quality standards, contains no leaked credentials, builds cleanly, and passes automated tests before committing changes.

## 1. Automated Verification (Python Script)

Run the cross-platform verification script:

```powershell
python .agents/skills/verify/scripts/verify.py
```

### Checks performed:
1. **Secret & Key Scanning:**
   - Detects private keys (`PEM`, `RSA`, `OPENSSH`).
   - Detects hardcoded JWT secrets.
   - Detects plaintext passwords and database connection strings.
   - Verifies no sensitive `.env` files are tracked in Git.
2. **Go Backend Validation:**
   - Checks code formatting with `gofmt`.
   - Runs static analysis with `go vet`.
   - Runs automated test suite with `go test`.

---

## 2. Manual Verification (Fallback)

1. **Secrets & Environment:**
   - Verify no hardcoded secrets or tracked `.env` files.
   - Ensure new required environment variables are documented in `server/.env.example`.

2. **Formatting & Style (Go):**
   - Run in `server/`:
     ```powershell
     cd server; go fmt ./...
     ```

3. **Static Analysis (Linting):**
   - Run in `server/`:
     ```powershell
     cd server; go vet ./...
     ```

4. **Automated Tests:**
   - Run tests with race condition detector:
     ```powershell
     cd server; go test -v -race ./...
     ```

5. **Report:**
   - Summarize verification findings before proceeding with the commit.
