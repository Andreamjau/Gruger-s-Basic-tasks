package main

// IsPrime checks if a number is a prime number.
// A prime number is a number greater than 1 that
// has no divisors other than 1 and itself.
//
// Example:
//   IsPrime(7) => true
//   IsPrime(10) => false
func IsPrime(n int) bool {
    if n <= 1{
        return false
    }
    for i := 2; i < n; i++{
        if n%i == 0 {
            return false
        }
    }
    return true
}
