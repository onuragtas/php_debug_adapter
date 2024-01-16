# PHP DEBUG ADAPTER

### Example Settings
```
{
    "listen": "0.0.0.0:10000",
    "mappings": [
        {
            "path": "/var/www/html/[path_on_server]",
            "url": "127.0.0.1:9999"
        }
    ]
}
```

Listen 10000 port for xdebug connections and if path mapping is contains /var/www/html/[path_on_server]<br>
send debug packets to 127.0.0.1:9999 ide port
