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

## 📁 Installation

To install Commander with a single command:

```bash
curl -sSL https://raw.githubusercontent.com/yashvesikar/commander/main/install.sh | bash
```

This script automatically detects your OS and architecture, downloads the latest release, and installs it to `/usr/local/bin`.

Alternatively, you can manually download a prebuilt binary from the [Releases page](https://github.com/yashvesikar/commander/releases):

```bash
chmod +x commander-linux-amd64
mv commander-linux-amd64 /usr/local/bin/commander
```

Replace `commander-linux-amd64` with the appropriate binary for your system.

---

## 📦 Features

- Auto-discovers apps via `commands.yml`
- Descriptive command listing
- Runs commands in app’s directory context
- Clear error reporting

---

## 🤝 Contributing

This project was primarily to manage my own home server but happy to take contributions.

To build the project locally:

```bash
git clone https://github.com/yashvesikar/commander.git
cd commander
go build -o commander main.go
```

You can now run `./commander` from your working directory. Make sure you have Go 1.22 or later installed.

---

## 🗪 License

MIT License
