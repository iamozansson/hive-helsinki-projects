totalNum=$(find . | wc -l)
result=$((totalNum*5))
printf "\t\vTotal files * 5: $result\v\n"
