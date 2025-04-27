# Standup

Standup is a command-line tool designed to help you manage your daily standup notes efficiently. It allows you to track what you worked on yesterday, what you're working on today, and any additional notes.

## Features
- Record daily standup updates.
- View previous standup notes.
- Add and manage notes for better productivity.

---

## Installation

Download the binary and add to path

## Downloadable Binaries

You can download pre-built binaries for macOS and Windows from the [Releases](https://github.com/VanBelleKoen/standup/releases) page.

### Running the Binary

#### macOS
1. Download the `standup-macos` binary from the Releases page.
2. Open a terminal and navigate to the directory where the binary is located.
3. Make the binary executable:
   ```bash
   chmod +x standup-macos
   ```
4. Run the binary:
   ```bash
   ./standup-macos
   ```

#### Windows
1. Download the `standup-windows.exe` binary from the Releases page.
2. Open a Command Prompt or PowerShell and navigate to the directory where the binary is located.
3. Run the binary:
   ```cmd
   standup-windows.exe
   ```

## Manual Installation

### Prerequisites
- **Go**: Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/).

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/VanBelleKoen/standup.git
   cd standup
   ```

2. Build the application:
   ```bash
   go build -o standup ./src
   ```

3. (Optional) Install the application globally:
   ```bash
   mv standup /usr/local/bin/
   ```

---

## Usage

### Running the Application
To run the application, use the following command:
```bash
./standup
```
If installed globally:
```bash
standup
```

### Commands
1. **Add Today's Updates**:
   - Add what you're working on today:
     ```bash
     standup pr-123
     ```

2. **Add Notes**:
   - Add a note:
     ```bash
     standup --note "Remember to review PR #123"
     ```

3. **View Standup Notes**:
   - View the standup:
     ```bash
     standup
     ```

4. **Help**:
   - Display help information:
     ```bash
     standup help
     ```

---

## Contributing
Feel free to fork the repository and submit pull requests. Contributions are welcome!

---

## License
This project is licensed under the MIT License.