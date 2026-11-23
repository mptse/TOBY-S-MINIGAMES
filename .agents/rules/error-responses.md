# Regla de Formato Unificado de Respuestas y Errores HTTP

Esta regla aplica a todos los controladores y endpoints expuestos en el backend de **TOBY'S MINI GAMES** (`server/internal/modules/*/handler/`).

---

## 1. Estructura Estándar de Respuesta JSON

Todas las respuestas emitidas por la API deben apegarse estrictamente a la convención provista por el paquete compartido `server/internal/shared/response/`.

### 1.1. Respuestas Exitosas (`2xx`)
```json
{
  "success": true,
  "data": { ... },
  "message": "Descripción concisa de la acción realizada"
}
```
- **`200 OK`:** Permite devolver el formato JSON estándar con `"data": null` cuando una operación exitosa no devuelve información (por ejemplo, confirmaciones de acción, cierre de sesión o eliminaciones con acuse en JSON).
- **`204 No Content`:** Debe devolver una respuesta completamente vacía, sin cuerpo JSON ni contenido.

### 1.2. Respuestas de Error (`4xx` / `5xx`)
```json
{
  "success": false,
  "error": {
    "code": "CODIGO_SEMANTICO_DE_ERROR",
    "message": "Mensaje legible y descriptivo para el cliente",
    "details": null
  }
}
```
- **`code`:** Cadena en mayúsculas tipo `SNAKE_CASE` que identifique el error de forma programática (ej. `INVALID_CREDENTIALS`, `USER_ALREADY_EXISTS`, `INSUFFICIENT_FUNDS`, `RESOURCE_NOT_FOUND`, `VALIDATION_ERROR`).
- **`message`:** Mensaje claro en español que explique el motivo del error.
- **`details`:**
  - Cuando existan errores de validación (`VALIDATION_ERROR`), debe contener un objeto mapa clave-valor donde cada clave sea el campo afectado y el valor describa el problema específico:
    ```json
    {
      "success": false,
      "error": {
        "code": "VALIDATION_ERROR",
        "message": "Los datos enviados no son válidos",
        "details": {
          "email": "El correo electrónico no tiene un formato válido",
          "password": "La contraseña es obligatoria"
        }
      }
    }
    ```
  - Cuando no haya información adicional que desglosar, `details` debe ser `null`.

---

## 2. Mapeo de Códigos de Estado HTTP

Los controladores deben retornar el código de estado HTTP adecuado según la naturaleza de la respuesta:

| Código HTTP | Cuándo utilizarlo | Ejemplos / Convención |
| :--- | :--- | :--- |
| **`200 OK`** | Consultas (`GET`) o actualizaciones (`PUT`, `PATCH`) exitosas. Admite `data: null` si no hay entidad. | Lectura de perfil, actualización de usuario. |
| **`201 Created`** | Creación exitosa de un recurso (`POST`). Retorna la entidad creada en `data`. | Registro completado, nuevo puntaje guardado. |
| **`204 No Content`** | Operación exitosa sin cuerpo de retorno. Respuesta totalmente vacía. | Eliminación sin cuerpo de confirmación. |
| **`400 Bad Request`** | JSON malformado, tipos de datos incompatibles o errores generales de petición. | Error al parsear JSON entrante. |
| **`401 Unauthorized`** | Petición a ruta protegida sin token, con token expirado o firma inválida. | Token JWT ausente o inválido. |
| **`403 Forbidden`** | Usuario autenticado pero sin permisos para ejecutar la acción. | Rol insuficiente para una acción. |
| **`404 Not Found`** | El recurso solicitado no existe en la base de datos. | Juego, usuario o ítem no encontrado. |
| **`409 Conflict`** | Conflictos con el estado actual del recurso o restricciones de unicidad. | `USER_ALREADY_EXISTS`, email/username ya en uso. |
| **`422 Unprocessable Entity`** | Solicitud sintácticamente válida, pero rechazada por una regla de negocio. | `INSUFFICIENT_FUNDS`, saldo insuficiente, partida ya finalizada. |
| **`500 Internal Server Error`** | Error inesperado o falla no recuperable de infraestructura/DB. | Falla de conexión a base de datos. |

---

## 3. Manejo de Pánico y Trazas Internas

- **Prohibido `panic()`:** Queda terminantemente prohibido utilizar `panic()` para controlar errores esperados o el flujo normal del negocio. Todos los errores deben propagarse de forma explícita mediante el tipo `error` de Go hacia los controladores.
- **Middleware de recuperación:** Gin debe contar obligatoriamente con middleware de recuperación (`gin.Recovery()` o middleware de recovery propio) para interceptar cualquier pánico imprevisto en tiempo de ejecución y evitar la caída del proceso.
- **Registro y confidencialidad:** Ante errores inesperados (`500 Internal Server Error`), el servidor debe registrar internamente el error y su stack trace en los logs. Jamás se debe exponer información sensible, credenciales o trazas de código al cliente en el payload JSON.

---

## 4. Centralización de Respuestas (`server/internal/shared/response/`)

El paquete **`server/internal/shared/response/`** debe centralizar todos los formatos y mecanismos de construcción de respuestas de la API.

- **Regla estricta:** Los handlers no deben construir manualmente estructuras JSON ad-hoc para cada endpoint. Deben apoyarse exclusivamente en las funciones provistas por este paquete compartido.
- **Funciones base sugeridas:**
  - `Success(c *gin.Context, httpStatus int, data any, message string)`
  - `Created(c *gin.Context, data any, message string)`
  - `Error(c *gin.Context, httpStatus int, code string, message string)`
  - `ValidationError(c *gin.Context, errors map[string]string)`
