const TOKEN_KEY = "tobys_mock_jwt";

function getBaseUrl() {
  const input = document.getElementById("api-base-url");
  return (input?.value || "http://localhost:8080/api/v1").replace(/\/+$/, "");
}

function getToken() {
  return sessionStorage.getItem(TOKEN_KEY) || "";
}

function setToken(token) {
  if (token) {
    sessionStorage.setItem(TOKEN_KEY, token);
  } else {
    sessionStorage.removeItem(TOKEN_KEY);
  }
}

async function apiRequest(method, path, body = null, useAuth = false, overrideBaseUrl = null) {
  const baseUrl = overrideBaseUrl || getBaseUrl();
  const headers = { "Content-Type": "application/json" };

  if (useAuth) {
    const token = getToken();
    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }
  }

  const options = { method, headers };
  if (body !== null && method !== "GET") {
    options.body = JSON.stringify(body);
  }

  try {
    const res = await fetch(`${baseUrl}${path}`, options);
    const json = await res.json().catch(() => ({ raw: "No JSON body" }));
    return { status: res.status, ok: res.ok, data: json, endpoint: `${method} ${baseUrl}${path}` };
  } catch (err) {
    return {
      status: 0,
      ok: false,
      data: { error: "Network error", detail: err.message },
      endpoint: `${method} ${baseUrl}${path}`,
    };
  }
}
