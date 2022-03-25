# resistor-counts

![alt text](https://github.com/arbaregni/resistor-counts/blob/main/problem-statement.png?raw=true)

For example, we know `R^2 = {2, 1, 1/2}` because each element can be constructed with two or less resistors.

- `1 + 1 = 2`
- `1 || 1 = 1/2`

For a more involved example, we have `3/4` in `R^4` because

- `1 || (1 + 1 + 1) = 1 || 3 = 3/4`

This is a program to brute-force calculate `R^n`.
