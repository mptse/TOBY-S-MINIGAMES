# Tobi's Minigames - Especificaciones de Cliente y Mock-Client

Documento de especificaciones para las interfaces que interactúan con el backend de **TOBY'S MINI GAMES** dentro del directorio `client/`.

---

## 1. Contexto y Estrategia del Área `client/`

El directorio `client/` albergará las interfaces de usuario que consumen los servicios del servidor. Para mantener un desarrollo ordenado y progresivo, el área se divide en dos etapas:

```text
client/
├── UI.md            # Este documento de especificaciones
├── mock-client/     # Cliente de prueba temporal (fase actual)
└── web/             # Frontend definitivo (fase futura)
```

---

## 2. Mock-Client (`client/mock-client/`)

### 2.1. Propósito
El `mock-client` es un cliente temporal y ligero desarrollado exclusivamente para probar y validar los endpoints reales del backend durante su construcción, sin bloquear el avance a la espera del diseño visual final.

### 2.2. Tecnologías del Mock-Client
- **HTML5:** Estructura básica de formularios y paneles de respuesta.
- **CSS3:** Estilos mínimos y utilitarios para legibilidad.
- **JavaScript Vanilla (ES6+):** Lógica de interacción y llamadas HTTP mediante la **Fetch API** nativa.
- **Sin dependencias pesadas:** No utiliza frameworks (React, Vue, Angular) ni empaquetadores complejos.

### 2.3. Estructura Prevista
```text
client/mock-client/
├── index.html       # Interfaz con formularios para cada endpoint
├── css/             # Estilos básicos
└── js/              # Módulos JS con llamadas fetch a la API
```

### 2.4. Funcionalidades a Validar con el Mock-Client
1. **Autenticación:**
   - Registro de usuario.
   - Login y recepción de token JWT (almacenado temporalmente en memoria o `sessionStorage`).
2. **Perfil de Usuario:**
   - Consulta de datos personales y subperfiles.
3. **Simulación de Minijuegos:**
   - Obtención del catálogo de juegos disponibles.
   - Envío de puntuación simulada (`POST /api/v1/scores`) con duración y metadatos.
4. **Rankings:**
   - Consulta de tablas de clasificación generales y por juego.
5. **Economía:**
   - Visualización de saldo de monedas y catálogo de recompensas.
   - Simulación de canje de artículos.

### 2.5. Reglas del Mock-Client
- Mantenerlo sincronizado con los contratos de endpoints documentados en [BACKEND.md](../server/BACKEND.md).
- Visualizar siempre el código de estado HTTP recibido (200, 201, 400, 401, 500) y el cuerpo JSON de respuesta para facilitar la depuración.

---

## 3. Frontend Definitivo (`client/web/`)

El subdirectorio `client/web/` está **reservado para el frontend final**.

- Se desarrollará conjuntamente una vez que la compañera de equipo complete la definición de lineamientos visuales, wireframes y maquetación de la interfaz.
- Consumirá los mismos endpoints REST probados previamente con el `mock-client`.
- **Regla estricta:** No inicializar frameworks ni crear código en `client/web/` hasta que los diseños y la arquitectura de frontend estén formalmente acordados.
