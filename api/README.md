# Playground

The `docker-compose up` command will start a swagger-ui. However, because the Account-API does not send the correct 
CORS headers, to use this swagger-ui you will need to disable web security. There are different ways to do this, when 
you have [Chromium](https://github.com/chromium/chromium) installed, you could try the following make command:

```bash
make open-swagger-ui
```

Which will basically perform:

```bash
chromium-browser \
  --disable-web-security \
  --user-data-dir="/tmp/chromium-debug/" \
  'http://localhost:7080/#/Health/get_health' 'http://localhost:7080/#/Health/get_health'

```
