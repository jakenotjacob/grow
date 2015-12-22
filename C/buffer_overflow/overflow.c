#include <stdlib.h>
#include <string.h>

void overflow(char *str){
    char buffer[20];
    strcpy(buffer, str);
}

int main(){
    char big_fucker[128];

    int i;
    //Overwrite SFP, ret, and *str
    for(i = 0; i < 128; i++){
        big_fucker[i] = 'A';
    }

//EIP?! -- unknown
//Stack is fucked; all heap freed, though
overflow(big_fucker);
exit(0);
}

//valgrind --tool=memcheck --leak-check=yes --show-reachable=yes --num-callers=20 --track-fds=yes ./test


