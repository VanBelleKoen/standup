# Standup Application

The Standup application helps you manage your daily standup notes. It allows you to organize your tasks and notes into categories such as "Today" and "Last Workday". This README provides an overview of the application's functionality and user flows.

## Usage

```bash
standup [command] [arguments]
```

### Commands

- `--help`
  - Displays the help message with a list of available commands and their descriptions.

- `--remove [entry]`
  - Removes a specific entry from today's list.
  - **Example:**
    ```bash
    standup --remove "Fix bug #123"
    ```

- `--reset`
  - Resets today's list, clearing all entries.
  - **Example:**
    ```bash
    standup --reset
    ```

- `--note [note content]`
  - Adds a note to the notes section.
  - **Example:**
    ```bash
    standup --note "Remember to review PRs"
    ```

- `--sync-branches`
  - Scans configured workspaces for Git repositories and detects new branches.
  - **Example:**
    ```bash
    standup --sync-branches
    ```

- `[entries...]`
  - Adds one or more entries to today's list.
  - **Example:**
    ```bash
    standup "Complete documentation" "Prepare for meeting"
    ```

## Description

The Standup application is designed to streamline your daily standup process. It categorizes entries into "Today" and "Last Workday" sections, allowing you to focus on what matters most.

### Features

1. **Help Command**
   - Provides a detailed overview of available commands and their usage.

2. **Entry Management**
   - Add, remove, and reset entries in the "Today" section.

3. **Notes**
   - Add notes to keep track of important reminders or information.

4. **Git Integration**
   - Detects new branches in Git repositories within configured workspaces.

5. **Automatic Initialization**
   - Automatically creates a `standup.json` file if it doesn't exist, with default values.

## Configuration

The Standup application uses a configuration file to define the workspaces to scan for Git repositories. This configuration file should be named `.standupconfig` and placed in the user's home directory.

### Example `.standupconfig`

```
workspaces: [
    "~/Documents/Workspace/Personal/project1",
    "~/Documents/Workspace/Personal/project2"
  ]
```

### Configuration Details

- **`workspaces`**: An array of paths to the directories you want to scan for Git repositories. Paths can use `~` to represent the user's home directory.

### How It Works

1. The application reads the `.standupconfig` file at startup.
2. Each workspace path is resolved to its absolute path.
3. The application scans the specified directories for Git repositories.
4. New branches detected in these repositories are logged and displayed.

### Default Behavior

If the `.standupconfig` file is missing or empty, the application will not perform any Git repository scanning. You will need to create and populate the file to enable this functionality.

### Example Workflow

1. Create a `.standupconfig` file in the home directory:
   ```bash
   echo 'workspaces: ["~/Documents/Workspace/Personal/project1"]' > ~/.standupconfig
   ```

2. Run the application with the `--sync-branches` command:
   ```bash
   standup --sync-branches
   ```

3. The application will scan the specified workspace for Git repositories and detect new branches.

## Example Workflow

1. **Add Entries**
   ```bash
   standup "Write unit tests" "Fix login issue"
   ```

2. **View Standup**
   - Run the application without any arguments to view the current standup data.

3. **Remove an Entry**
   ```bash
   standup --remove "Fix login issue"
   ```

4. **Add a Note**
   ```bash
   standup --note "Follow up with the design team"
   ```

5. **Reset Today's List**
   ```bash
   standup --reset
   ```

6. **Sync Git Branches**
   ```bash
   standup --sync-branches
   ```

## License

This project is licensed under the MIT License.