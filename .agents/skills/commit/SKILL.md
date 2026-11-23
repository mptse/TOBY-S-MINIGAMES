---
name: commit
description: Use when the user requests committing pending changes.
---

1. Inspect working tree and changes (`git status`).
2. **Execute mandatory quality verification (Verify skill):**
   - Run repo verification script:
     ```powershell
     python .agents/skills/verify/scripts/verify.py
     ```
   - Confirm no leaked secrets (private keys, hardcoded JWTs, plaintext passwords, tracked `.env` files).
   - Ensure code is formatted (`gofmt`), passes static analysis (`go vet`), and tests pass (`go test`).
   - **Blocker:** If verification fails, resolve issues before proceeding. Do not propose or execute commits on failure.
3. Check previous commits to maintain language and style (`git log -n 5`).
4. Check `git config --local user.email` and `git config --local user.name`.
5. If unset or first commit, confirm desired email and author name with the user.
6. Group changes logically.
7. Follow commit sequence hierarchy:
   DB → models → config → services → client/UI.
8. Do not mix unrelated changes in a single commit.
9. Propose each commit:
   - Included files
   - Reason for change
   - Proposed message
10. Await explicit user approval before executing the commit.
11. Use **Conventional Commits**:
    `feat:`, `fix:`, `refactor:`, `test:`, `docs:`, `chore:`, etc.
12. Maintain commit history language (Spanish by default unless requested).
13. Execute the commit.
14. Verify result with `git status` and `git log -n 1`.
