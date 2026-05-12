#include <cstdio>
#include <iostream>
using namespace std;

int demofunc();

int main() {
    int res = demofunc();
    printf("demofunc output %d\n", res);
    return 0;
}

int demofunc() {
    int res;
    for (int i = 0; i < 10; i++) {
        res += i;
    }
    return res;
}
