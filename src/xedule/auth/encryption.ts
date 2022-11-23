import * as crypto from 'crypto';

const algorithm = 'aes-256-ctr';
const secretKey = process.env.APP_SECRET;

export interface XeduleEncryptionHash {
  iv: string;
  content: string;
}

export default class XeduleEncryption {
  static encrypt(value): XeduleEncryptionHash {
    const iv = crypto.randomBytes(16);

    const cipher = crypto.createCipheriv(algorithm, secretKey, iv);

    const encrypted = Buffer.concat([cipher.update(value), cipher.final()]);

    return {
      iv: iv.toString('hex'),
      content: encrypted.toString('hex'),
    };
  }

  static decrypt(hash: XeduleEncryptionHash): string {
    const decipher = crypto.createDecipheriv(
      algorithm,
      secretKey,
      Buffer.from(hash.iv, 'hex')
    );

    const decrypted = Buffer.concat([
      decipher.update(Buffer.from(hash.content, 'hex')),
      decipher.final(),
    ]);

    return decrypted.toString();
  }
}
