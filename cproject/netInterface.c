#include <stdio.h>
int get_number();

int main() {
    printf("demofunc");
    int res = get_number();
    printf("demofunc output: %d\n", res);
    return 0;
}

int get_number() {
    int res;
    for (int i = 0; i < 10; i++) {
        res += i;
    }
    return res;
}
