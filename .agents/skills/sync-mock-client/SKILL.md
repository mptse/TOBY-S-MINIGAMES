---
name: sync-mock-client
description: Sync backend endpoint contracts with client/mock-client by updating native Fetch API calls and HTML test panels.
---

# Sync Mock-Client Skill

Use this skill when a backend endpoint in `server/` is created or modified and needs to be reflected in `client/mock-client/` for manual testing.

## Synchronization Steps

1. **Inspect Endpoint Contract:**
   - Review handler in `server/internal/modules/<module>/handler/`.
   - Identify: HTTP method (`GET`, `POST`, etc.), route path, auth requirements (`Authorization: Bearer <token>`), request DTO payload, and expected response schema (`success`, `data`, `error`).

2. **Update JavaScript Client (`client/mock-client/js/api.js`):**
   - Add/update fetch function:
     - Set required headers (`Content-Type: application/json`).
     - Attach JWT token if route is protected.
     - Return HTTP status code and parsed JSON response.

3. **Update HTML Interface (`client/mock-client/index.html`):**
   - Add/update card section for the endpoint under the corresponding module.
   - Add input fields (`<input>`, `<textarea>`, `<select>`) for request parameters.
   - Add submit button with event listener.
   - Provide output container (`<pre><code>`) displaying HTTP status and formatted JSON.

4. **Validate:**
   - Test success flow (`200 OK` / `201 Created`).
   - Test error flow (`400 Bad Request` / `401 Unauthorized` / `404 Not Found`).
   - Confirm response renders clearly with no browser console errors.
