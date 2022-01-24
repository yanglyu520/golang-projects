## golang commandline tool with go

### Description

The advantage of using golang to build the binary and distribute to different system.

This will be a task manager CLI tool that has a database associated and you can add task, list your incomplete task and complete a task.

### Requirement

1. create 3 different commands `add`, `list`, `do`
2. `add`: add a new task to the list
3. `list`: print all the incomplete tasks
4. `do`: complete a task

### Breakdown

1. Interact with a database called BoltDB
2. Build the CLI shell using [cobra](https://github.com/spf13/cobra)
