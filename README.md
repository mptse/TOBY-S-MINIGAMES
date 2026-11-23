# TOBY'S MINI GAMES

Plataforma web de entretenimiento interactivo diseñada para ofrecer una experiencia de juego directa, rápida y ligera desde cualquier navegador, sin descargas ni instalaciones.

El proyecto permite a los jugadores registrarse, personalizar perfiles, guardar su progreso en subperfiles por juego, competir en tablas de clasificación globales y periódicas, y canjear recompensas en una tienda virtual de puntos.

---

## Estructura del Monorepo

El repositorio está organizado en tres áreas claramente delimitadas para permitir un desarrollo progresivo y desacoplado:

```text
tobys-mini-games/
│
├── AGENTS.md                # Guía principal para agentes y desarrolladores
├── README.md                # Este documento general
├── .gitignore               # Reglas de exclusión de Git
├── docker-compose.yml       # Orquestación de servicios en desarrollo
│
├── server/                  # Backend: API REST en Go, Gin, GORM y PostgreSQL
├── client/                  # Cliente: Mock-client de pruebas y futuro frontend web
└── games/                   # Minijuegos: Aplicaciones web jugables independientes
```

---

## Contenido de las Áreas y Tecnologías

### 1. `server/` — Backend
Servidor central y API REST de la plataforma.
- **Tecnologías:** Go 1.24+, Gin Framework, GORM, PostgreSQL, Docker, JWT, bcrypt y migraciones SQL versionadas. (Redis previsto como integración futura justificada).
- **Arquitectura:** Monolito modular con Arquitectura Hexagonal Lite (`domain`, `application`, `infrastructure`, `handler`).

### 2. `client/` — Cliente de la Plataforma
Interfaces que interactúan con el backend.
- **`mock-client/`:** Cliente de prueba temporal en HTML5, CSS3 y JavaScript vanilla (Fetch API) para validar los endpoints del backend mientras se definen los diseños finales.
- **`web/`:** Reservado para el frontend definitivo de la plataforma que se desarrollará en una fase posterior.

### 3. `games/` — Minijuegos
Aplicaciones web interactivas que corren de forma autónoma en el navegador.
- **Tecnologías previstas:** HTML5 Canvas/DOM, CSS3, JavaScript/TypeScript y Phaser 3 (solo cuando la complejidad gráfica/física lo amerite).
- **Integración:** Se contempla un contrato común denominado `PlatformSDK` para estandarizar el ciclo de vida y el reporte de puntuaciones al backend.

---

## Documentación del Proyecto

Cada área cuenta con su propio documento de especificaciones técnicas:

| Documento | Enfoque |
| :--- | :--- |
| [AGENTS.md](AGENTS.md) | **Guía Principal:** Reglas globales, principios de desarrollo y flujo de trabajo. |
| [server/BACKEND.md](server/BACKEND.md) | **Backend:** Arquitectura Hexagonal Lite, módulos funcionales y base de datos. |
| [client/UI.md](client/UI.md) | **Cliente:** Funcionamiento del `mock-client` y directrices de UI. |
| [games/GAMES.md](games/GAMES.md) | **Minijuegos:** Estándares de desarrollo de minijuegos y `PlatformSDK`. |
