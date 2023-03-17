# LM1: Assignment Brute Force Ceasar Cipher

import string

# performs encryption by shifting the text provided
def encrypt(text, key):

    encrypted_text = list(range(len(text)))
    alphabet = string.ascii_lowercase  # 'abcdefghijklmnopqrstuvwxyz'
    first_half = alphabet[:key]
    second_half = alphabet[key:]
    shifted_alphabet = second_half + first_half

    for i, letter in enumerate(text.lower()):

        if letter in alphabet:
            original_index = alphabet.index(letter)
            new_letter = shifted_alphabet[original_index]
            encrypted_text[i] = new_letter
        else:
            encrypted_text[i] = letter

    return "".join(encrypted_text)

# performs decryption by shifting the text provided
def decrypt(text, key):
    """ when the key is known """
    decrypted_text = list(range(len(text)))
    alphabet = string.ascii_lowercase
    first_half = alphabet[:key]
    second_half = alphabet[key:]
    shifted_alphabet = second_half + first_half

    for i, letter in enumerate(text.lower()):

        if letter in alphabet:
            index = shifted_alphabet.index(letter)
            original_letter = alphabet[index]
            decrypted_text[i] = original_letter
        else:
            decrypted_text[i] = letter
    return "".join(decrypted_text)

# Loop through all possible key values
def decrypt_brute_force(encrypted_message):
    """ when the shift is unknown """
    for key in range(26):
        decrypted_message = ""
        # Loop through each character in the encrypted message
        for char in encrypted_message:
            if char.isalpha():
                # Shift the character by the current key value
                shifted_char = chr((ord(char) - key - 65) % 26 + 65)
                decrypted_message += shifted_char
            else:
                # Keep non-alpha characters the same
                decrypted_message += char
        # Print the decrypted message for the current key value
        print("Key " + str(key) + ": " + decrypted_message)

# Main implementation
def main():
    key = 5
    encrypted_message = encrypt("This is a secret", key)

    print("Decrypting the message with the key known: ")
    decrypted_message_ = decrypt(encrypted_message, key)
    print("Key " + str(key) + ": " + decrypted_message_ + "\n")

    print("Decrypting the message with the key unknown - Brute force implementation: ")
    decrypt_brute_force(encrypted_message)

if __name__ == "__main__":
    main()
