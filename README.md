# Logger go
Utility for management of logger on your golang app.

# Use Case:
```go
package main

func main(){
    
    logger.Debug("Variable x = %v", x)
    logger.Info("Servidor iniciando...")
    logger.Success("Conexión exitosa")
    logger.Warn("Cache no disponible")
    logger.Error("Error: %v", err)
    logger.Fatal("Error crítico")
    logger.Security("Login exitoso")
}
```

You will have something like this: 

```bash
[2026-03-22 10:30:00] [DEBUG] Variable x = 42
[2026-03-22 10:30:00] [INFO] Servidor iniciando...
[2026-03-22 10:30:00] [SUCCESS] Conexión exitosa
[2026-03-22 10:30:00] [WARN] Cache no disponible
[2026-03-22 10:30:00] [ERROR] Error: base de datos no responde
[2026-03-22 10:30:00] [SECURITY] Login exitoso
```

