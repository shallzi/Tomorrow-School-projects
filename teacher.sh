#!/bin/bash
export INTERVIEW_NUM=$(grep -l "Annabel Church" interviews/* | sed 's/[^0-9]*//g')

echo "$INTERVIEW_NUM"

cat "interviews/interview-$INTERVIEW_NUM"

echo "$MAIN_SUSPECT"