# sqlitecat
cat for sqlite database

```
$ sqlitecat <filename>
```

If there's more than 1 table in the database
```
$ sqlitecat <filename> <tablename>
```

With config envvars
```
$ SEP=", " QUERY="size != 0" sqlitecat all.db
```

## Envvars

Set envvar `QUERY` for "where" expressions. `QUERY="size > 100" ./sqlitecat` will only print records mathcing query "SELECT * FROM TABLENAME WHERE size > 100".

Set envvar `SEP` if you want row values separated by sth else than `" | "`. 
