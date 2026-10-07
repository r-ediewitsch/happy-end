import json
import urllib.request
import urllib.error

API_URL = "http://localhost:8080/dev/logs" 

def seed():
    try:
        with open("transcript.json", "r", encoding="utf-8") as f:
            logs = json.load(f)
    except Exception as e:
        print(f"Gagal membaca file JSON: {e}")
        return

    print(f"Memulai pengiriman {len(logs)} data ke {API_URL}...\n")

    # 2. Loop dan kirim request POST
    for index, item in enumerate(logs, start=1):
        payload = json.dumps(item).encode("utf-8")
        req = urllib.request.Request(
            API_URL,
            data=payload,
            headers={"Content-Type": "application/json"},
            method="POST"
        )

        try:
            with urllib.request.urlopen(req) as response:
                if response.status in (200, 201):
                    print(f"[{index}/{len(logs)}] Sukses: {item.get('char')}")
        except urllib.error.HTTPError as e:
            print(f"[{index}/{len(logs)}] Gagal ({e.code}): {e.read().decode('utf-8')}")
        except Exception as e:
            print(f"[{index}/{len(logs)}] Koneksi error: {e}")

    print("\nSelesai!")

if __name__ == "__main__":
    seed()