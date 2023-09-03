#! /bin/bash

curl https://learn.reboot01.com/assets/superhero/all.json | jq .[$HERO_ID].connections.relatives | sed 's/"//g' 