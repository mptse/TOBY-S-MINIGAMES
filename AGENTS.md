# Tobi's Minigames - Guía General del Monorepo

Documento principal de entrada al proyecto formativo **TOBY'S MINI GAMES** (SENA). Toda persona o agente de IA que colabore en este repositorio debe leer y acatar estas instrucciones antes de iniciar cualquier tarea.

---

## 1. ¿Qué es Toby's Minigames?

**TOBY’S MINI GAMES** es una plataforma web de entretenimiento interactivo diseñada para ofrecer una experiencia de juego directa, rápida y ligera desde cualquier navegador, sin necesidad de descargas ni instalaciones.

Permite a los usuarios registrarse, personalizar sus perfiles, almacenar su progreso en subperfiles por juego, competir en tablas de clasificación globales y periódicas, canjear recompensas en una tienda virtual y disfrutar de minijuegos web independientes.

---

## 2. Estructura del Monorepo

El repositorio está organizado en tres áreas claramente delimitadas:

```text
tobys-mini-games/
│
├── AGENTS.md                # Guía general del monorepo y punto de entrada para agentes
├── README.md                # Resumen del proyecto y catálogo de áreas
├── .gitignore               # Exclusiones de Git para todo el monorepo
├── docker-compose.yml       # Orquestación de servicios en desarrollo (Go app + PostgreSQL)
│
├── server/                  # ÁREA BACKEND: Servidor en Go, Gin, GORM y PostgreSQL
│   ├── BACKEND.md           # Especificaciones técnicas y arquitectónicas del backend
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── tests/
│   ├── Dockerfile
│   └── go.mod
│
├── client/                  # ÁREA CLIENTE: Interfaces que consumen el backend
│   ├── UI.md                # Especificaciones de mock-client y frontend definitivo
│   ├── mock-client/         # Cliente temporal (HTML, CSS, JS vanilla) para probar la API
│   └── web/                 # Reservado para el frontend definitivo (futuro)
│
└── games/                   # ÁREA MINIJUEGOS: Aplicaciones web jugables independientes
    ├── GAMES.md             # Especificaciones de desarrollo e integración de minijuegos
    ├── trivia/
    └── memory/
```

> [!NOTE]
> La estructura anterior representa la organización objetivo. Los directorios y archivos de código se crearán progresivamente a medida que se desarrolle cada componente, sin generar carpetas vacías innecesarias.

---

## 3. Sistema de Documentación (Regla Fundamental)

Las instrucciones del proyecto están distribuidas en cuatro documentos con responsabilidades bien delimitadas:

| Documento | Área | Cuándo consultarlo |
| :--- | :--- | :--- |
| [AGENTS.md](AGENTS.md) | **Monorepo Global** | Siempre, como punto de inicio y marco de referencia general. |
| [server/BACKEND.md](server/BACKEND.md) | **Backend (`server/`)** | Obligatorio al trabajar en el servidor, base de datos, APIs o Docker backend. |
| [client/UI.md](client/UI.md) | **Cliente (`client/`)** | Obligatorio al trabajar en `mock-client/` o en el frontend `web/`. |
| [games/GAMES.md](games/GAMES.md) | **Minijuegos (`games/`)** | Obligatorio al desarrollar o integrar aplicaciones en `games/`. |

> [!IMPORTANT]
> **Reglas estrictas de consulta y límites de área:**
> - Si trabajas en `server/`, consulta obligatoriamente [server/BACKEND.md](server/BACKEND.md).
> - Si trabajas en `client/`, consulta obligatoriamente [client/UI.md](client/UI.md).
> - Si trabajas en `games/`, consulta obligatoriamente [games/GAMES.md](games/GAMES.md).
> - Si una tarea involucra varias áreas (e.g. nuevo endpoint y su prueba en `mock-client`), debes consultar todos los documentos correspondientes.
> - **Las reglas de este documento ([AGENTS.md](AGENTS.md)) aplican a todo el repositorio.**
> - **No confundir áreas:** `server/internal/modules/games/` administra funcionalidades y metadatos desde el backend; `games/` contiene las aplicaciones jugables independientes.

---

## 4. Principios Generales de Desarrollo

1. **Evitar archivos y funciones gigantes:** Mantener código modular, legible y desacoplado.
2. **Principio de Responsabilidad Única (SRP):** Cada módulo, capa o script debe tener un único motivo de cambio.
3. **Respetar fronteras de capas:** No violar las dependencias arquitectónicas definidas en cada área.
4. **Cero asunciones de dependencias:** Verificar siempre `go.mod` o `package.json` antes de importar librerías. No agregar dependencias sin justificación documentada.
5. **Cero asunciones de entorno:** Toda nueva variable de configuración requerida debe documentarse inmediatamente en el `.env.example` respectivo.
6. **No sobreingeniería (YAGNI):** Diseñar soluciones simples y extensibles sin introducir abstracciones innecesarias.
7. **No alterar decisiones arquitectónicas sin acuerdo:** Cualquier cambio técnico de base debe fundamentarse previamente.
8. **Pruebas automatizadas:** Toda nueva funcionalidad debe respaldarse con pruebas unitarias o de integración.
9. **No intervenir áreas ajenas sin justificación:** Si una tarea es exclusiva de `server/`, no tocar `client/` ni `games/`, y viceversa.
10. **Commits estructurados:** Utilizar la skill en [.agents/skills/commit/SKILL.md](.agents/skills/commit/SKILL.md) siguiendo la convención de *Conventional Commits*.

---

## 5. Flujo de Trabajo para Implementar Funcionalidades

1. **Comprensión y Consulta:** Revisar el requerimiento y consultar el documento específico (`server/BACKEND.md`, `client/UI.md` o `games/GAMES.md`).
2. **Definición de Contrato:** Si involucra comunicación cliente-servidor o entre juegos y plataforma, definir la estructura de datos (request/response/eventos) antes de codificar la lógica interna.
3. **Implementación por Capas:** Desarrollar respetando el orden y las dependencias arquitectónicas del área.
4. **Validación:** Ejecutar pruebas automatizadas y comprobar el comportamiento con el `mock-client` o scripts de prueba.
5. **Documentación de Lecciones:** Si surgió un bloqueo o error técnico real, registrar la solución en la tabla de lecciones aprendidas.
6. **Commit Estructurado:** Versionar los cambios con la skill de commits.

---

## 6. Lecciones Aprendidas

> [!NOTE]
> Registrar únicamente problemas reales encontrados durante el desarrollo. Cada entrada debe incluir causa raíz y regla preventiva concreta. No agregar recomendaciones genéricas ni hipótesis.

| Fecha | Módulo / Área | Problema detectado | Causa raíz / Solución | Regla o lección para el futuro |
| :--- | :--- | :--- | :--- | :--- |
| *YYYY-MM-DD* | *Ej: server/auth* | *Descripción breve del error o bloqueo* | *Qué causó la falla y cómo se solucionó* | *Qué hacer o evitar en próximas implementaciones* |
