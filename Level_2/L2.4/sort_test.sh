#!/bin/bash

GREEN='\e[32m'
NC='\e[0m'
RED='\e[31m'

arr=(-n -r -u)

printf "${GREEN}-----RUNNING TESTS-----${NC}\n"

#sort files/example1.txt >a
#./main.exe -b files/example1.txt >b
#result=$(diff a b)
#failed=0
c=1

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

for (( i = 0; i < 3; i++ )) # для одного флага files/example.txt
do
    sort ${arr[i]} files/example.txt >a
    ./main.exe ${arr[i]} files/example.txt >b
    result=$(diff a b)
    if [ $? -eq 0 ]; then
	    printf " TEST #$c ${GREEN}PASSED${NC}\n"
    else
	    printf " TEST #$c ${RED}FAILED${NC}\n"
	    printf "$result"
	    ((failed++))
    fi
((c++))
done

for (( i = 0; i < 3; i++ )) # для одного флага files/example1.txt
do
    sort ${arr[i]} files/example1.txt >a
    ./main.exe ${arr[i]} files/example1.txt >b
    result=$(diff a b)
    if [ $? -eq 0 ]; then
	    printf " TEST #$c ${GREEN}PASSED${NC}\n"
    else
	    printf " TEST #$c ${RED}FAILED${NC}\n"
	    printf "$result"
	    ((failed++))
    fi
((c++))
done

printf "\n ${GREEN}-----DONE[$((c - failed))/$((c))]-----${NC}\n"

#rm a b
