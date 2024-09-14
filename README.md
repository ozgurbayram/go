## Project Structure

<pre>
<code>
project_name/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── user/
│   │   ├── domain/
│   │   ├── application/
│   │   └── adapters/
│   │       ├── inbound/
│   │       │   └── http/
│   │       │       ├── user_controller.go
│   │       │       ├── user_router.go
│   │       │       └── dto.go
│   │       └── outbound/
│   │           └── persistence/
│   │               └── postgres/
│   │                   └── user_repository.go
</code>
</pre>
