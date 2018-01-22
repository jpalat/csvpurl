CSVPurl
-------

Simple command line tool to sample csv files.  It matches the header with the contents of a particular line.  The index starts with 0 as the first line AFTER the header.

## Arguments
 - filename
 - line number

## Example

```bash
$ cat inventory.csv
item,cost,qty
ball, 3.00, 23
bat, 16.00, 4
glove, 32.00, 5
helmet, 19.00, 15
```

```bash
$ ./csvpurl inventory.csv 3
item : glove
cost :  32.00
qty :  5
```
