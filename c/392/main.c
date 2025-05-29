#include <stdbool.h>
#include <string.h>
#include <stdio.h>

bool isSubsequence(char* s, char* t) {
    int i = 0;
    int j = 0;
    int sLen = strlen(s);
    int tLen = strlen(t);
    while (i < sLen && j < tLen) {
        if (s[i] == t[j]) {
            i++;
        }
        j++;
    }
    return i >= sLen;
}

int main() {
    char* s = "abc";
    char* t = "ahbgdc";
    printf("%d\n", isSubsequence(s, t));
    s = "axc";
    t = "ahbgdc";
    printf("%d\n", isSubsequence(s, t));
}
