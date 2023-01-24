# Frigatier

A dead-simple utility that reads Frigate notifications over MQTT and pipes them to a messenger of choice. 

Currently supported messengers:

1. Slack
2. Discord

To run this, you need to either download or compile the binary and then create a config.yml file next to the binary:


```
frigate:
  host: <frigate IP>
  port: <frigate port>
  
mqtt:
  host: <mqtt IP>
  port: <mqtt port>

messengers:
  slack:
    enabled: True
    token: <slack token>
    channel: <channel ID>

  discord:
    enabled: True
    token: <discord bot token>
    channel: <channel or user id>

```

This is a very VERY early release. Your cat might be consumed by black holes.
