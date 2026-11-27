function updateTokenPreview() {
  const preview = document.getElementById("token-preview");
  const token = getToken();
  preview.textContent = token ? `${token.slice(0, 28)}...` : "Ninguno";
}

function renderResponse(result) {
  const statusEl = document.getElementById("response-status");
  const endpointEl = document.getElementById("response-endpoint");
  const bodyEl = document.getElementById("response-body");

  statusEl.textContent = result.status === 0 ? "NETWORK_ERR" : `HTTP ${result.status}`;
  statusEl.className = `badge ${result.ok ? "ok" : "err"}`;
  endpointEl.textContent = result.endpoint;
  bodyEl.textContent = JSON.stringify(result.data, null, 2);

  if (result.ok && result.data?.data?.token) {
    setToken(result.data.data.token);
    updateTokenPreview();
  }
}

document.getElementById("btn-clear-token").addEventListener("click", () => {
  setToken("");
  updateTokenPreview();
});

document.getElementById("btn-health").addEventListener("click", async () => {
  const rootUrl = getBaseUrl().replace(/\/api\/v1$/, "");
  const res = await apiRequest("GET", "/health", null, false, rootUrl);
  renderResponse(res);
});

document.getElementById("form-register").addEventListener("submit", async (e) => {
  e.preventDefault();
  const fd = new FormData(e.target);
  const res = await apiRequest("POST", "/auth/register", {
    email: fd.get("email"),
    password: fd.get("password"),
    nickname: fd.get("nickname"),
  });
  renderResponse(res);
});

document.getElementById("form-login").addEventListener("submit", async (e) => {
  e.preventDefault();
  const fd = new FormData(e.target);
  const res = await apiRequest("POST", "/auth/login", {
    email: fd.get("email"),
    password: fd.get("password"),
  });
  renderResponse(res);
});

document.getElementById("form-guest").addEventListener("submit", async (e) => {
  e.preventDefault();
  const nickname = new FormData(e.target).get("nickname")?.toString().trim();
  const payload = nickname ? { nickname } : {};
  const res = await apiRequest("POST", "/auth/guest", payload);
  renderResponse(res);
});

document.getElementById("btn-get-me").addEventListener("click", async () => {
  const res = await apiRequest("GET", "/users/me", null, true);
  renderResponse(res);
});

document.getElementById("form-patch-me").addEventListener("submit", async (e) => {
  e.preventDefault();
  const fd = new FormData(e.target);
  const payload = {};
  for (const field of ["nickname", "bio", "avatar_path"]) {
    const val = fd.get(field)?.toString().trim();
    if (val) payload[field] = val;
  }
  const res = await apiRequest("PATCH", "/users/me", payload, true);
  renderResponse(res);
});

document.getElementById("form-get-subprofile").addEventListener("submit", async (e) => {
  e.preventDefault();
  const slug = encodeURIComponent(new FormData(e.target).get("slug").toString().trim());
  const res = await apiRequest("GET", `/users/me/games/${slug}`, null, true);
  renderResponse(res);
});

document.getElementById("form-list-games").addEventListener("submit", async (e) => {
  e.preventDefault();
  const cat = new FormData(e.target).get("category")?.toString().trim();
  const query = cat ? `?category=${encodeURIComponent(cat)}` : "";
  const res = await apiRequest("GET", `/games${query}`);
  renderResponse(res);
});

document.getElementById("form-get-game").addEventListener("submit", async (e) => {
  e.preventDefault();
  const slug = encodeURIComponent(new FormData(e.target).get("slug").toString().trim());
  const res = await apiRequest("GET", `/games/${slug}`);
  renderResponse(res);
});

updateTokenPreview();
