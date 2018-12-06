"""
	A simple example of the usage of decorator in Python
"""
import time

def display_time(func):
	'''show the running time of a function'''
	def wrapper(*args, **kwargs):
		t1 = time.time()
		result = func(*args, **kwargs)
		t2 = time.time()
		print("Running time: {:.6}s".format(t2 - t1))
		return result
	return wrapper

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
	

@display_time
def find_primes(n):
	'''find primes from 2 to n using Sieve of Eratosthenes'''
	result = []
	for i in range(2, n+1):
		if is_prime(i):
			result.append(i)
	return result

print(find_primes(100))