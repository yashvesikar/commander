# Commander

**Commander** is a lightweight CLI tool to organize and run project-specific commands using simple YAML files. It discovers apps in the current directory and executes predefined shell commands from each app’s `commands.yml`.

---

## 📁 Directory Structure

```
apps/
├── caddy/
│   └── commands.yml
├── nginx/
│   └── commands.yml
```

Each subdirectory represents an app and must contain a `commands.yml` file.

---

## 📝 Example `commands.yml`

```yaml
commands:
  start:
    cmd: "docker-compose up -d"
    description: "Start the app"
  stop: "docker-compose down"
  reload:
    cmd: "nginx -s reload"
    description: "Reload config"
```

Supports:

- Simple format: `stop: "docker-compose down"`
- Object format: includes `cmd` and optional `description`

---

## 🚀 Usage

List all apps:

```bash
$ commander
```

List app commands:

```bash
$ commander <app>
```

Run a command:

```bash
$ commander <app> <command>
```

Example:

```bash
$ commander caddy reload
```

---

## ⚙️ Build

```bash
go build -o commander main.go
```

Run `commander` from the root of your `apps/` directory.

---

## 📦 Features

- Auto-discovers apps via `commands.yml`
- Descriptive command listing
- Runs commands in app’s directory context
- Clear error reporting

---

## License

MIT License
