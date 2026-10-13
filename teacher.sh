#!/bin/bash

# Шаг 1: Находим файл по фразе и выделяем чистый номер через sed (отрезая всё до дефиса)
export INTERVIEW_NUM=$(grep -l "Ms. Church" interviews/* | sed 's/[^0-9]//g')

# Шаг 2: Выводим созданную переменную окружения
echo "$INTERVIEW_NUM"

# Шаг 3: Выводим содержимое файла интервью, собирая правильный путь
cat "interviews/interview-$INTERVIEW_NUM"

# Шаг 4: Выводим содержимое переменной MAIN_SUSPECT
echo "$MAIN_SUSPECT"