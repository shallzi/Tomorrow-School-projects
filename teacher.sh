#!/bin/bash

# Шаг 1: выделяем только число из найденного файла интервью
export INTERVIEW_NUM=$(grep -l "Annabel Church" interviews/* | grep -o '[0-9]*')

# Шаг 2: выводим значение переменной окружения
echo "$INTERVIEW_NUM"

# Шаг 3: выводим содержимое интервью
cat interviews/interview-$INTERVIEW_NUM

# Шаг 4: выводим содержимое MAIN_SUSPECT
echo "$MAIN_SUSPECT"