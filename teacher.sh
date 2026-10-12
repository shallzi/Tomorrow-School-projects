export INTERVIEW_NUM=$(grep -l "Annabel Church" interviews/* | grep -o '[0-9]*')
echo "$INTERVIEW_NUM"
cat "interviews/$INTERVIEW_NUM"
echo "$MAIN_SUSPECT"