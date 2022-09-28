#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* swap(char* input) {
    size_t current_len = 40;
    size_t new_len = 0;

    char* output = (char*)malloc(current_len);

    while(*input) {
        char* add = input;
        size_t add_len = 1;
        switch (*add)
        {
            case '&':
                add = "&amp;";
                add_len = 5;
                break;
            case '<':
                add = "&lt;";
                add_len = 4;
                break;
            case '>':
                add = "&gt;";
                add_len = 4;
                break;
        }
        if (new_len + add_len >= current_len) {
            current_len += 10;
            output = (char*)realloc(output, current_len);
        }
        memcpy(&output[new_len], add, add_len);
        new_len += add_len;

        input++;
    }
    return output;
}

int main(int argc, const char *argv[])
{
    char* start_value = "Foo<&>Bar";
    printf("Initial: %s\n", start_value);

    char* end_value = swap(start_value);
    printf("  After: %s\n", end_value);
    free(end_value);
}
