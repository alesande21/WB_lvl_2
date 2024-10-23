#!/bin/bash

GREEN='\e[32m'
NC='\e[0m'
RED='\e[31m'

arr=(-b)

printf "${GREEN}-----RUNNING TESTS-----${NC}\n"

#sort files/example1.txt >a
#./main.exe -b files/example1.txt >b
#result=$(diff a b)
#failed=0
#c=1

# TEST 1

sort -n files/example.txt >a
./main.exe -n files/example.txt  >b
result=$(diff a b)

if [ $? -eq 0 ]; then
	printf " TEST #$c ${GREEN}PASSED${NC}\n"
else
	printf " TEST #$c ${RED}FAILED${NC}\n"
	printf "$result"
	((failed++))
fi
((c++))

#for (( i = 0; i < 1; i++ )) # для одного флага
#do
#    cat ${arr[i]} test_1.txt >a
#    ./s21_cat ${arr[i]} test_1.txt >b
#    result=$(diff a b)
#    if [ $? -eq 0 ]; then
#	    printf " TEST #$c ${GREEN}PASSED${NC}\n"
#    else
#	    printf " TEST #$c ${RED}FAILED${NC}\n"
#	    printf "$result"
#	    ((failed++))
#    fi
#((c++))
#done










printf "\n ${GREEN}-----DONE[$((c - failed))/$((c))]-----${NC}\n"

#rm a b
