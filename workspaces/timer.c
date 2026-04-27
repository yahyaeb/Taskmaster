#include <unistd.h>
#include <stdio.h>

int main(void) {
    for (int i = 0; i < 5; i++) {
        printf("Tick %d\n", i);
        sleep(1); // Sleep for 1 second
    }
    return 0;
}

