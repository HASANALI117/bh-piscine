#! /bin/bash

curl https://learn.reboot01.com/assets/superhero/all.json | jq ' .[] | select(.id==170) | .name , .powerstats.power , .appearance.gender'