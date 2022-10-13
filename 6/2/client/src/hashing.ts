import CryptoJS from "crypto-js";

const SALT_CONST = "localhost"
const KEY_SIZE = 512/8
const ITERATIONS = 2048


export function hashPass(pass: string) {
    return CryptoJS.PBKDF2(pass, SALT_CONST, {keySize: KEY_SIZE, iterations: ITERATIONS}).toString();
}
