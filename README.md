# One-time Pad

Let ⊕ represent the operation XOR.

The one-time pad takes a plaintext P and a random key K (same length as P), that together produce a ciphertext C:

**C = P ⊕ K**

The decryption is identical to the encryption: 

**P = C ⊕ K**

We can verify **C ⊕ K = P ⊕ K ⊕ K = P** because XORing anything with itself results in an all-zero bit string.

## Trivia
### Unique key
Each key should be used only once. If a same key K is used to encrypt P1 and P2, we could compute the following:

**C1 ⊕ C2 = (P1 ⊕ K) ⊕ (P2 ⊕ K) = P1 ⊕ P2 ⊕ K ⊕ K = P1 ⊕ P2**

Then one would know the XOR difference between P1 and P2, and if one of them is known, the other can easily be recovered.
