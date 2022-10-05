#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "utility.h"

// Utfør fuzzing med address sanitizer på C funksjonen dere lagde
// i øving 4 b)

// Reparer feil dere finner gjennom fuzzing, eller introduser feil
// som oppdages gjennom fuzzing

// Sett opp CI (Continuous Integration) løsning som utfører
// fuzzing med address sanitizer

// Se instruksjoner for å kjøre fuzzing i en terminal (merk at du
// kan begrense antall sekunder libFuzzer skal kjøre med
// argumentet -max_total_time)

// Du kan bruke docker image’et archlinux som utgangspunkt for
// å installere nyere llvm/clang++ med støtte for libFuzzer

// Frivillig: Lag tester, og kjør testene også i CI (fuzzing-example
// inneholder en eksempel test

int main(int argc, const char *argv[])
{
    char* start_value = "Foo<&>Bar";
    printf("Initial: %s\n", start_value);

    char* end_value = swap(start_value);
    printf("  After: %s\n", end_value);
    free(end_value);
}
