#!/usr/bin/env python3
"""
Script de verificación de calidad y escaneo de secretos para Toby's Minigames.
Ubicación: .agents/skills/verify/scripts/verify.py
"""

import os
import re
import sys
import subprocess
from pathlib import Path

# Directorios y extensiones excluidos del escaneo de secretos
IGNORED_DIRS = {
    ".git",
    ".gemini",
    "node_modules",
    "vendor",
    "bin",
    "obj",
    ".idea",
    ".vscode",
}

IGNORED_FILES = {
    "verify.py",         # Este mismo script contiene los patrones regex
    ".env.example",      # Contiene ejemplos ficticios
}

IGNORED_EXTENSIONS = {
    ".png", ".jpg", ".jpeg", ".gif", ".ico", ".svg",
    ".woff", ".woff2", ".ttf", ".eot",
    ".mp3", ".wav", ".ogg",
    ".zip", ".tar", ".gz",
    ".exe", ".dll", ".so", ".dylib",
}

# Patrones para detectar llaves, credenciales y secretos filtrados
SECRET_PATTERNS = [
    (
        "Clave privada PEM / RSA / SSH",
        re.compile(r"-----BEGIN (?:RSA |EC |DSA |OPENSSH )?PRIVATE KEY-----", re.IGNORECASE)
    ),
    (
        "Secreto JWT hardcodeado",
        re.compile(r"""(?:jwt_secret|jwtSecret|JWT_SECRET)\s*[:=]\s*["']([^"']{8,})["']""")
    ),
    (
        "Contraseña hardcodeada sospechosa",
        re.compile(r"""(?:password|passwd|db_password|DB_PASSWORD)\s*[:=]\s*["']([^"']{6,})["']""", re.IGNORECASE)
    ),
    (
        "API Key o Token de servicio",
        re.compile(r"""(?:api_key|apiKey|API_KEY|auth_token|authToken)\s*[:=]\s*["']([A-Za-z0-9_\-]{20,})["']""")
    ),
    (
        "Cadena de conexión con contraseña explícita",
        re.compile(r"""(?:postgres|postgresql|mysql|redis)://[^:]+:([^@]+)@""")
    )
]

# Valores permitidos que no deben disparar falso positivo en contraseñas
ALLOWED_PLACEHOLDERS = {
    "password", "secret", "change_me", "changeme", "example", "admin", "postgres",
    "your_password", "your_jwt_secret", "test", "demo", "root", "123456"
}

def print_header(title: str):
    print(f"\n{'=' * 60}")
    print(f"  {title}")
    print(f"{'=' * 60}")

def check_tracked_env_files(root_dir: Path) -> list[str]:
    """Verifica si archivos .env reales están bajo control de versiones."""
    issues = []
    try:
        res = subprocess.run(
            ["git", "ls-files", "*.env", "*/*.env"],
            cwd=str(root_dir),
            capture_output=True,
            text=True
        )
        if res.returncode == 0 and res.stdout.strip():
            for line in res.stdout.strip().splitlines():
                if not line.endswith(".example"):
                    issues.append(f"Archivo de entorno sensible rastreado en Git: {line}")
    except Exception:
        pass
    return issues

def scan_file_for_secrets(file_path: Path) -> list[str]:
    """Escanea el contenido de un archivo en busca de patrones de secretos."""
    findings = []
    try:
        with open(file_path, "r", encoding="utf-8", errors="ignore") as f:
            for line_num, line in enumerate(f, start=1):
                clean_line = line.strip()
                if not clean_line or clean_line.startswith(("//", "#", "/*", "*")):
                    # Si la línea es un comentario explicativo, evaluar con menor sensibilidad
                    pass
                for label, pattern in SECRET_PATTERNS:
                    match = pattern.search(line)
                    if match:
                        matched_val = match.group(1) if match.groups() else match.group(0)
                        if matched_val.strip().lower() in ALLOWED_PLACEHOLDERS:
                            continue
                        # Mostrar exclusivamente la línea afectada y la categoría, NUNCA el secreto
                        findings.append(
                            f"Línea {line_num}: detectado patrón [{label}] (contenido sensible oculto)"
                        )
    except Exception as e:
        findings.append(f"No se pudo leer archivo: {e}")
    return findings

def scan_repo_secrets(root_dir: Path) -> bool:
    """Recorre el monorepo y busca posibles secretos filtrados."""
    print_header("1. Escaneo de Secretos y Llaves Filtradas")
    total_findings = 0

    # 1. Chequeo de archivos .env trackeados
    env_issues = check_tracked_env_files(root_dir)
    for issue in env_issues:
        print(f"  [ALERTA] {issue}")
        total_findings += 1

    # 2. Recorrido de archivos del proyecto
    for root, dirs, files in os.walk(root_dir):
        # Excluir directorios ignorados
        dirs[:] = [d for d in dirs if d not in IGNORED_DIRS]

        for file in files:
            if file in IGNORED_FILES:
                continue
            ext = os.path.splitext(file)[1].lower()
            if ext in IGNORED_EXTENSIONS:
                continue

            file_path = Path(root) / file
            findings = scan_file_for_secrets(file_path)
            if findings:
                rel_path = file_path.relative_to(root_dir)
                print(f"\n  [PELIGRO] Posibles secretos en {rel_path}:")
                for item in findings:
                    print(f"    -> {item}")
                total_findings += len(findings)

    if total_findings == 0:
        print("  [OK] No se detectaron secretos ni credenciales en texto plano.")
        return True
    else:
        print(f"\n  [FALLO] Se detectaron {total_findings} posible(s) fuga(s) de credenciales.")
        return False

def check_go_backend(root_dir: Path) -> bool:
    """Ejecuta formateo, análisis estático y pruebas del backend Go si existen."""
    print_header("2. Verificación de Backend Go (server/)")
    server_dir = root_dir / "server"
    go_mod = server_dir / "go.mod"

    if not go_mod.exists():
        print("  [INFO] 'server/go.mod' aún no existe. Saltando ejecución de Go.")
        return True

    success = True

    # 1. go fmt check
    print("  -> Verificando formato con gofmt...")
    try:
        res = subprocess.run(["gofmt", "-l", "."], cwd=str(server_dir), capture_output=True, text=True)
        if res.returncode == 0 and res.stdout.strip():
            print(f"  [AVISO] Los siguientes archivos necesitan 'go fmt':\n{res.stdout}")
            success = False
        else:
            print("  [OK] Formato de código correcto.")
    except FileNotFoundError:
        print("  [AVISO] 'go' no está en el PATH del sistema.")
        return True

    # 2. go vet
    print("  -> Ejecutando análisis estático (go vet)...")
    res = subprocess.run(["go", "vet", "./..."], cwd=str(server_dir), capture_output=True, text=True)
    if res.returncode != 0:
        print(f"  [FALLO] go vet encontró problemas:\n{res.stderr or res.stdout}")
        success = False
    else:
        print("  [OK] go vet completado sin advertencias.")

    # 3. go test
    print("  -> Ejecutando pruebas unitarias (go test)...")
    res = subprocess.run(["go", "test", "-v", "./..."], cwd=str(server_dir), capture_output=True, text=True)
    if res.returncode != 0:
        print(f"  [FALLO] Fallaron las pruebas unitarias:\n{res.stderr or res.stdout}")
        success = False
    else:
        print("  [OK] Todas las pruebas unitarias pasaron.")

    return success

def main():
    # .agents/skills/verify/scripts/verify.py -> parents[4] es la raíz del monorepo
    root_dir = Path(__file__).resolve().parents[4]
    print(f"Iniciando verificación en raíz del monorepo: {root_dir}")

    secrets_ok = scan_repo_secrets(root_dir)
    go_ok = check_go_backend(root_dir)

    print_header("Resumen de Verificación")
    if secrets_ok and go_ok:
        print("  [ÉXITO] Todo limpio y verificado. Listo para commit.")
        sys.exit(0)
    else:
        print("  [ERROR] Se encontraron problemas que deben resolverse antes de continuar.")
        sys.exit(1)

if __name__ == "__main__":
    main()
