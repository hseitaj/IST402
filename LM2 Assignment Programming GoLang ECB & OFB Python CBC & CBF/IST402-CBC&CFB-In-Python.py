"""
    - IST 402, CBC and CBF Python Solution.
    
    - Please installl pycryptodome before running the application:
      pip install pycryptodome

    - This program demonstrates how to use the AES algorithm 
      to encrypt text using two modes: CBC and CBF. 
      The program prompts the user for plaintext input and 
      generates a random encryption key and initialization 
      vectors for both CBC and CFB modes.

    Inputs:
    - The user is prompted to enter plaintext when the program is run.

    Limitations:
    - The program only encrypts ASCII text. Non-ASCII characters 
      will need to be converted to bytes before being encrypted.

    - The program does not decrypt the ciphertext. It only 
      demonstrates how to encrypt plaintext using AES.

    - The program uses a fixed block size of 16 bytes, 
      the plaintext input must be padded to a multiple 
      of 16 bytes in order to be encrypted using CBC mode.

    - The program does not handle errors related to 
      incorrect input type or length.
"""


from Crypto.Cipher import AES
from Crypto.Util.Padding import pad
from Crypto.Random import get_random_bytes

# Generate a random encryption key
encryptionKey = get_random_bytes(16)

# Initialization vector for Cipher Block Chaining (CBC) mode
cbcIv = get_random_bytes(16)

# Initialization vector for Cipher Feedback (CFB) mode
cfbIv = get_random_bytes(16)

# Create AES cipher objects for CBC and CFB modes
cbcCipher = AES.new(encryptionKey, AES.MODE_CBC, cbcIv)
cfbCipher = AES.new(encryptionKey, AES.MODE_CFB, cfbIv, segment_size=8)

# Get plaintext from user
plainText = input("Enter plaintext: ")

# Pad the plaintext to a multiple of 16 bytes for CBC mode
paddedPlainText = pad(plainText.encode('utf-8'), AES.block_size)

# Encrypt the plaintext using CBC mode
cbcEncryptedText = cbcCipher.encrypt(paddedPlainText)

# Encrypt the plaintext using CFB mode
cfbEncryptedText = cfbCipher.encrypt(plainText.encode('utf-8'))

# Display details of encryption process
print("Encryption Key:", encryptionKey.hex())
print("IV (CBC):", cbcIv.hex())
print("IV (CFB):", cfbIv.hex())
print("Plaintext:", plainText)
print("Padded plaintext (CBC):", paddedPlainText)
print("Encrypted text (CBC):", cbcEncryptedText.hex())
print("Encrypted text (CFB):", cfbEncryptedText.hex())
