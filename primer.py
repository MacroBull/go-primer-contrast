import time

LIM = 9999999

tStart = time.time()

primes = [2, 3]

for n in xrange(5, LIM, 2):
	for m in primes:
		if m*m > n:
			curIsPrime = True
			break
		elif 0 == n % m:
			curIsPrime = False
			break
	if curIsPrime:
		primes.append(n)

print("{}:\t{} \n".format(time.time() - tStart, len(primes)))
