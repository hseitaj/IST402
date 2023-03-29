import random

# Generate random key and initialization vector for CBC
cbc_key = bytes(random.randint(0, 255) for i in range(16))
cbc_iv = bytes(random.randint(0, 255) for i in range(16))

# Generate random key for CBF
cbf_key = bytes(random.randint(0, 255) for i in range(16))


def xor_bytes(b1, b2):
    return bytes(x ^ y for x, y in zip(b1, b2))


def pad_message(message):
    padding = 16 - len(message) % 16
    return message + bytes([padding] * padding)


def unpad_message(message):
    padding = message[-1]
    return message[:-padding]


def encrypt_cbc(message, key, iv):
    message = pad_message(message)
    encrypted = b""
    previous_block = iv
    for i in range(0, len(message), 16):
        block = message[i : i + 16]
        block = xor_bytes(block, previous_block)
        previous_block = block
        # Perform AES encryption here
        encrypted += block
    return encrypted


def decrypt_cbc(ciphertext, key, iv):
    decrypted = b""
    previous_block = iv
    for i in range(0, len(ciphertext), 16):
        block = ciphertext[i : i + 16]
        # Perform AES decryption here
        block = xor_bytes(block, previous_block)
        previous_block = ciphertext[i : i + 16]
        decrypted += block
    decrypted = unpad_message(decrypted)
    return decrypted


def encrypt_cbf(message, key):
    encrypted = b""
    for i in range(len(message)):
        encrypted += bytes([(message[i] + key[i % len(key)]) % 256])
    return encrypted


def decrypt_cbf(ciphertext, key):
    decrypted = b""
    for i in range(len(ciphertext)):
        decrypted += bytes([(ciphertext[i] - key[i % len(key)]) % 256])
    return decrypted


plaintext = input("Enter plaintext message: ").encode()

cbc_ciphertext = encrypt_cbc(plaintext, cbc_key, cbc_iv)
print(f"CBC Key: {cbc_key.hex()}")
print(f"CBC IV: {cbc_iv.hex()}")
print(f"CBC Ciphertext: {cbc_ciphertext.hex()}")
decrypted_plaintext = decrypt_cbc(cbc_ciphertext, cbc_key, cbc_iv)
print(f"Decrypted CBC plaintext: {decrypted_plaintext.decode()}")

cbf_ciphertext = encrypt_cbf(plaintext, cbf_key)
print(F"CBF Key: {cbf_key.hex()}")
print(F"CBF Ciphertext: {cbf_ciphertext.hex()}")
decrypted_plaintext = decrypt_cbf(cbf_ciphertext, cbf_key)
print(F"Decrypted CBF plaintext: {decrypted_plaintext.decode()}")
