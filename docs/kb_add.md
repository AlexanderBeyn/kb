## kb add

Add a new task

### Synopsis

Add a new task

You will be prompted for a title if you don't specify it on the command line.
If the title on the command line ends with a single '+' by itself, you will
be prompted for a description for the new task.

```
kb add [flags] [[%%proj] %col] [title] [+]
```

### Examples

```

# Add a new task to the default column:
kb add This is my new task

# Add a new task to a column "backlog", prompting for a description:
kb add %backlog This is a new backlog task +

# Add a new task, prompting for a column and title:
kb add %
```

### Options

```
  -h, --help   help for add
```

### SEE ALSO

* [kb](kb.md)	 - 

