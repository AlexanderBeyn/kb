## kb move

Move a task to a new column

### Synopsis

Move a task to a new column

Tasks will be moved from the default column, if a source column is not
specified. If a target column is not specified, you will be prompted to
select one.

```
kb move [flags] [%%proj] [^%source_col] [%target_col] [/search]
```

### Examples

```

# Find a task matching "awesome" in the default column and move it to "done":
kb move %done /awesome

# Display all tasks in "backlog" and select one to move to "today":
kb move ^%backlog %today
```

### Options

```
  -h, --help   help for move
```

### SEE ALSO

* [kb](kb.md)	 - 

