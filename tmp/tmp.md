# Secrets

## Encrypt data

```bash
TO_ENCRYPT=$(echo "Some text to be encrypted" | base64)

GCP_PROJECT_ID="slate-00"
KEY_RING="slate"
KEY="slate"
ENV="dev"
curl -s -X POST "https://cloudkms.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/global/keyRings/${KEY_RING}/cryptoKeys/${KEY}:encrypt" \
    -d "{\"plaintext\":\"${TO_ENCRYPT}\"}" \
    -H "Authorization:Bearer $(gcloud auth application-default print-access-token)" \
    -H "Content-Type:application/json" \
    -o tmp_some-text_cloudkms-${ENV}.json
```

tmp_some-text_cloudkms-dev.json:

```json
{
    "name": "projects/slate-00/locations/global/keyRings/slate/cryptoKeys/slate/cryptoKeyVersions/1",
    "ciphertext": "CiQAkJV1aBJN3b0rKQ2rgh1nBv3F3t7lFKJH+n2DSXFSblOR8cwSQwB6oxFwme6y52kz535s8fr9HJZZdl2ESjy2Ofu3z+EcK6jsFVPAc0pla4tq5etYjr0qqzWE5RS932NzzqghnDlgmpk="
}
```

## Decrypt data

```bash
# ciphertext
# CiQ...gmpk=

TO_DECRYPT=$(echo "CiQ...gmpk=")

GCP_PROJECT_ID="slate-00"
KEY_RING="slate"
KEY="slate"
ENV="dev"
curl -s -X POST "https://cloudkms.googleapis.com/v1/projects/${GCP_PROJECT_ID}/locations/global/keyRings/${KEY_RING}/cryptoKeys/${KEY}:decrypt" \
    -d "{\"ciphertext\":\"${TO_DECRYPT}\"}" \
    -H "Authorization:Bearer $(gcloud auth application-default print-access-token)" \
    -H "Content-Type:application/json"
    -o tmp_some-text.json
```

tmp_some-text.json:

```json
{
    "plaintext": "WyJTb21lIHRleHQgdG8gYmUgZW5jcnlwdGVkIl0K"
}
```

```bash
DECRYPTED=$(echo "WyJTb21lIHRleHQgdG8gYmUgZW5jcnlwdGVkIl0K" | base64 -D)
# > ["Some text to be encrypted"]
```
