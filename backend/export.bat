@echo off
setx DB_USERNAME root /m
setx DB_PASSWORD password /m
setx DB_DATABASE database_name /m
setx DATABASE_HOST 127.0.0.1 /m
setx KAFKA_EVENTHUB_ENDPOINT "cbq-hackathon.servicebus.windows.net:9093" /m
setx KAFKA_EVENTHUB_CONNECTION_STRING "Endpoint=sb://cbq-hackathon.servicebus.windows.net/;SharedAccessKeyName=n;SharedAccessKey=p3fH0pzw46YajywaIyAaWRK+HGqMBLgBV+AEhNWlq+4=;EntityPath=team8" /m
