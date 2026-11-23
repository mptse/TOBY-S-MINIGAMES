# Regla de Estándares de Seguridad (Security Standards)

Esta regla aplica a todo el desarrollo del backend (`server/`) y manejo de datos sensibles en **TOBY'S MINI GAMES**.

---

## 1. Hashing de Contraseñas
- **Uso estricto de bcrypt:** Ninguna contraseña de usuario debe almacenarse en texto plano.
- **Coste de hashing:** Utilizar un coste mínimo de `bcrypt.DefaultCost` (10) o superior (12 recomendado para producción).
- **Prohibido reversibilidad:** Nunca crear mecanismos que intenten desencriptar o retornar la contraseña del usuario.

## 2. Prevención de Inyección SQL
- **Consultas Parametrizadas Obligatorias:** Todas las operaciones de base de datos a través de GORM o SQL crudo deben utilizar placeholders de parámetros (`?` o `$1`, `$2`).
- **Prohibida la concatenación:** Jamás concatenar variables de entrada de usuario dentro de strings SQL (`fmt.Sprintf("SELECT * FROM users WHERE name = '%s'", name)` está estrictamente prohibido).

## 3. Autenticación y Tokens JWT
- **Firma robusta:** Los tokens deben firmarse utilizando un algoritmo seguro (ej. `HMAC-SHA256`).
- **Secreto por entorno:** La clave secreta de firmado debe obtenerse de la variable `JWT_SECRET` en `.env`. Prohibido hardcodear cadenas por defecto.
- **Validación completa:** Toda verificación de token debe comprobar:
  1. Integridad de la firma.
  2. Fecha de expiración (`exp`).
  3. Presencia de claims requeridos (`user_id`).
- **Expiración obligatoria:** Ningún token de acceso puede emitirse sin una fecha de expiración finita y razonable.

## 4. Gestión de Secretos y Configuración
- **Prohibido versionar credenciales:** Ningún archivo `.env`, contraseña, clave privada o credencial de base de datos debe ser incluido en Git.
- **Plantilla .env.example:** Toda variable requerida debe figurar en `.env.example` con valores de ejemplo o vacíos, nunca valores reales de producción.

## 5. Validación y Sanitización de Entradas
- Todos los payloads de entrada deben validarse en la capa `handler/` antes de llegar a la capa de aplicación.
- Validar formato de correos electrónicos, longitudes mínimas y máximas de contraseñas y tipos de datos numéricos para evitar desbordamientos o datos maliciosos.
