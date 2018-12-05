def is_prime(n):
	'''check a number is a prime'''
	if n == 0:
		return False
	if n == 1:
		return False
	if n == 2:
		return True
		
	for i in range(2, n-1):
		if n % i == 0:
			return False
	return True
	
	
def find_primes(n):
	'''find primes from 2 to n using Sieve of Eratosthenes'''
	primes = [True] * (n+1)
	primes[0] = False
	primes[1] = False
	
	result = []
	for i in range(2, n+1):
		if primes[i]:
			result.append(i)
			for j in range(2*i, n+1, i):
				primes[j] == False
	
	return result

print(find_primes(10000))