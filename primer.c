#include <stdio.h>
#include <sys/time.h>
#include <unistd.h>

#define LIM 9999999


int primes[LIM + 1];

int main(int argc, char **argv) {
	struct timeval tStart, tEnd;
	
	gettimeofday(&tStart, NULL);
	
	int i, n, m;
	int index;
	
	primes[0] = 2;
	primes[1] = 3;
	index = 2;
	
	for (n = 5;n<LIM;n+=2){
		for (i = 0;m = primes[i], m*m<n;i++){
			if (0 == n % m) break;
			
		}
		if (m*m>n) primes[index++] = n;
	}
	
	gettimeofday(&tEnd, NULL);
	printf("%f:\t%d", tEnd.tv_sec -  tStart.tv_sec + (tEnd.tv_usec -  tStart.tv_usec)*1e-6, index);
	
    return 0;
}
