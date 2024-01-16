# PHP DEBUG ADAPTER

```
wget https://github.com/onuragtas/php_debug_adapter/releases/latest/download/php_debug_adapter_Darwin_arm64 -O /usr/local/bin/pda
chmod +x /usr/local/bin/pda
pda
```

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
