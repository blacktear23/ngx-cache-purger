# ngx-cache-purger

ngx-cache-purge is a tool for purge Nginx cache files.

## Usage

```
Usage of ./ngx-cache-purger:
  -k string
        Key Prefix
  -p string
        Cache path
```

Example:

```
./ngx-cache-purger -p /data/cache -k "/path/something"
```

ngx-cache-purger will scan nginx cache path and iterate all cached files then get the KEY for this file and then match the prefix. If prefix matched the file will be deleted.
