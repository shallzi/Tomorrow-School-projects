#!/bin/bash
curl https://01.tomorrow-school.ai/assets/superhero/all.json | jq ' .[] | select( .id ==  170) | .name'
curl https://01.tomorrow-school.ai/assets/superhero/all.json | jq ' .[] | select( .id ==  170) | .powerstats | .power'
curl https://01.tomorrow-school.ai/assets/superhero/all.json | jq ' .[] | select( .id ==  170) | .appearance | .gender'