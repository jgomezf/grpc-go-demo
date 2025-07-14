# 🧪 gRPC Go Demo

Este proyecto es una **demo educativa** que muestra los 4 tipos de comunicación gRPC usando Go. Ideal para quienes están aprendiendo gRPC desde cero y desean ver casos reales y simples.

---

## 🔧 Tecnologías

- 🔹 [Go](https://golang.org/)
- 🔹 [gRPC](https://grpc.io/)
- 🔹 [Protocol Buffers (proto3)](https://developers.google.com/protocol-buffers)

---

## 🚀 Tipos de Comunicación

| Tipo de RPC             | Método            | Descripción                                 |
|--------------------------|-------------------|---------------------------------------------|
| Unary RPC               | `SayHello`        | Request/Response de 1 a 1                   |
| Server Streaming         | `GreetManyTimes`  | Cliente envía 1 request, servidor responde múltiples veces |
| Client Streaming         | `LongGreet`       | Cliente envía múltiples mensajes, recibe 1 respuesta |
| Bidirectional Streaming  | `GreetEveryone`   | Cliente y servidor se comunican en paralelo |

---

## 📁 Estructura del Proyecto

```
grpc-go-demo/
├── client/           # Cliente gRPC con los 4 tipos de llamadas
│   └── main.go
├── server/           # Servidor gRPC con implementación de los métodos
│   └── main.go
├── proto/            # Definición del servicio en Protobuf
│   └── greet.proto
├── go.mod
└── README.md
```

---

## 🧰 Requisitos

- Go 1.18+
- `protoc` (Protocol Buffers Compiler)
- Plugins:
  - `protoc-gen-go`
  - `protoc-gen-go-grpc`

Instálalos así:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Asegúrate de tener el PATH actualizado
export PATH="$PATH:$(go env GOPATH)/bin"
```

---

## 📦 Instalación y Ejecución

### 1. Clona el repositorio

```bash
git clone https://github.com/tu_usuario/grpc-go-demo.git
cd grpc-go-demo
go mod tidy
```

### 2. Genera el código a partir del `.proto`

```bash
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

---

### 3. Ejecuta el servidor

```bash
go run server/main.go
```

### 4. Ejecuta el cliente

En otra terminal:

```bash
go run client/main.go
```

Verás la ejecución de:

- Unary
- Server Streaming
- Client Streaming
- Bidirectional Streaming

---